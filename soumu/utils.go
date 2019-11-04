package soumu

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func checkStatusCode(res *http.Response) error {
	sc := res.StatusCode
	if sc >= 200 && sc < 300 {
		return nil
	}
	if sc == 400 {
		var errJSON ErrorBody
		if err := decodeBody(res, &errJSON, nil); err != nil {
			return HTTPStatusError(sc)
		}
		for _, errArr := range errJSON.Err {
			logger.Warnf("Invalid Request:", errArr.ErrMsg)
		}
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
