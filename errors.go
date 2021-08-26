package soumuradio

import (
	"fmt"
)

// HTTPStatusError is error for Http bad response
func HTTPStatusError(sc int) error {
	return fmt.Errorf("bad response status code %d", sc)
}

// SystemError is error for API Server error
func SystemError() error {
	return fmt.Errorf("system error")
}
