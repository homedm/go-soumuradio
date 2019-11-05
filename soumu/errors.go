package soumu

import (
	"fmt"
)

func HTTPStatusError(sc int) error {
	return fmt.Errorf("bad response status code %d", sc)
}

// ErrorBody is json response, if request options is failed.
type ErrorBody struct {
	Errs Errs `json:"errs"`
	Err  Err  `json:"err"`
}

type Errs struct {
	ErrPost  string `json:"errPost"`
	ErrCount string `json:"errCount"`
}

type Err []struct {
	ErrMsg string `json:"errMsg"`
	ErrCd  string `json:"errCd"`
}
