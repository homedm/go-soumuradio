package soumu

import (
	"fmt"
)

func HTTPStatusError(sc int) error {
	return fmt.Errorf("bad response status code %d", sc)
}

// ErrorBody is json response, if request options is failed.
type ErrorBody struct {
	Errs struct {
		ErrPost  string `json:"errPost"`
		ErrCount string `json:"errCount"`
	} `json:"errs"`
	Err []struct {
		ErrMsg string `json:"errMsg"`
		ErrCd  string `json:"errCd"`
	} `json:"err"`
}
