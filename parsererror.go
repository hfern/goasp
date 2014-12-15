package goasp

import (
	"fmt"
)

type ParseError struct {
	message string
}

func (err ParseError) Error() string {
	return err.message
}

type UnknownObjectIdError struct {
	ParseError
	token otype
}

func (err UnknownObjectIdError) Error() string {
	return fmt.Sprintf("Unknown Object ID (0x%02X)", err.token)
}
