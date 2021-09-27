package godb

import (
	"fmt"
)

func stringValueError(message, value string) error {
	return fmt.Errorf("%s: '%s'", message, value)
}
