package errors

import "fmt"

var (
	BadRequestError = fmt.Errorf("Bad Request")
	InternalError   = fmt.Errorf("Internal")
)
