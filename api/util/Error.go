package util

import (
	"fmt"
)

type ErrorResponse struct {
	Message string
	Id      string
}

var ErrInternal = fmt.Errorf("internal error")

var ErrRecordNotFound = fmt.Errorf("record not found")
