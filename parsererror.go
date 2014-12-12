package goasp

import (
	"fmt"
)

type UnknownObjectIdError struct {
	token otype
}

func (err UnknownObjectIdError) Error() string {
	return fmt.Sprintf("Unknown Object ID (0x%02X)", err.token)
}
