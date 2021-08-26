package soumuradio

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
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

func parseRadioSpec1(str string) ([]RadioSpec, error) {
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
					if len(v.RadioFormat) == 0 {
						rf := string(strRune[j:i])
						v.RadioFormat = strings.Fields(rf)
					} else {
						v.Freq = strings.ReplaceAll(string(strRune[j:i]), " ", "")
					}
					i += 2
					j = i
				} else if i+1 < len(strRune) && string(strRune[i+1]) == "n" {
					powerstr := strings.ReplaceAll(string(strRune[j:i]), " ", "")
					if len(powerstr) > 0 {
						power, err := rmSIPrefix(powerstr)
						if err != nil {
							return nil, err
						}
						v.Power = power
					}
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

func rmSIPrefix(str string) (float64, error) {
	u := "W"

	idx := strings.Index(str, u)
	if idx == -1 {
		return 0, fmt.Errorf("could not parsed error: %v", str)
	}

	corr := map[rune]int{
		'd': -1,
		'c': -2,
		'm': -3,
		'µ': -6,
		'n': -9,
		'p': -12,
		'f': -15,
		'a': -18,
		'z': -21,
		'y': -24,
		'h': 2,
		'k': 3,
		'M': 6,
		'G': 9,
		'T': 12,
		'P': 15,
		'E': 18,
		'Z': 24,
	}

	if unicode.IsNumber(rune(str[idx-1])) {
		return strconv.ParseFloat(string(str[:idx]), 64)
	}
	for k, v := range corr {
		if rune(str[idx-1]) == k {
			num, err := strconv.ParseFloat(string(str[:idx-1]), 64)
			if err != nil {
				return 0, err
			}

			return num * math.Pow10(v), nil
		}
	}
	return 0, fmt.Errorf("could not parsed error: %v", str)
}
