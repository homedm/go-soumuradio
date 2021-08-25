package soumuradio

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"unicode"

	"github.com/tomato3713/soumuradio/internal"
)

func checkStatusCode(res *http.Response) error {
	sc := res.StatusCode
	if sc >= 200 && sc < 300 {
		return nil
	}
	if sc == 400 {
		var errJSON internal.ErrorBody
		if err := decodeBody(res, &errJSON, nil); err != nil {
			logf("invalid Request, but cannot decode JSON Response: %v", err)
		}
		for _, errArr := range errJSON.Err {
			logf("invalid Request: %v", errArr.ErrMsg)
		}
		return errJSON.InvalidRequestError()
	}

	if sc == 500 {
		logf("system error")
		return SystemError()
	}

	return HTTPStatusError(sc)
}

func decodeBody(resp *http.Response, out interface{}, f *os.File) error {
	defer resp.Body.Close()

	if f != nil {
		resp.Body = ioutil.NopCloser(io.TeeReader(resp.Body, f))
		defer f.Close()
	}

	decorder := json.NewDecoder(resp.Body)
	return decorder.Decode(out)
}

func parseRadioSpec(str string) ([]RadioSpec, error) {
	var rs []RadioSpec
	strRune := []rune(str)

	for j := 0; j < len(strRune); j++ {
		var v RadioSpec
		for i := j; i < len(strRune); i++ {
			if unicode.IsSpace(strRune[i]) {
				continue
			}
			if string(strRune[i]) == string("\\") {
				if i+1 < len(strRune) && string(strRune[i+1]) == "t" {
					if v.RadioFormat == "" {
						v.RadioFormat = strings.ReplaceAll(string(strRune[j:i]), " ", "")
					} else {
						v.Freq = strings.ReplaceAll(string(strRune[j:i]), " ", "")
					}
					i += 2
					j = i
				} else if i+1 < len(strRune) && string(strRune[i+1]) == "n" {
					v.Power = strings.ReplaceAll(string(strRune[j:i]), " ", "")
					i += 2
					j = i - 1
					rs = append(rs, v)
					break
				}
			}
		}
	}
	return rs, nil
}
