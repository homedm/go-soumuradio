package soumu

import (
	"fmt"
)

// HTTPStatusError is error for Http bad response
func HTTPStatusError(sc int) error {
	return fmt.Errorf("bad response status code %d", sc)
}

// ErrorBody is json response, if request options is failed.
type ErrorBody struct {
	Errs Errs `json:"errs"`
	Err  Err  `json:"err"`
}

// Errs is errors information
type Errs struct {
	ErrPost  string `json:"errPost"`
	ErrCount string `json:"errCount"`
}

// Err is error information
type Err []struct {
	ErrMsg string `json:"errMsg"`
	ErrCd  string `json:"errCd"`
}
