package util

import (
	"fmt"
)

type ErrorResponse struct {
	Message   string
	ErrorUUID string
}

var ErrInternal = fmt.Errorf("internal error")

var ErrRecordNotFound = fmt.Errorf("record not found")
