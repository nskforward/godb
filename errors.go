package godb

import (
	"fmt"
)

func ValueError(message, value string) error {
	return fmt.Errorf("%s: '%s'", message, value)
}
