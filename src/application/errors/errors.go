package errors

import (
	"encoding/json"
	"errors"
)

var (
	BadRequestError = errors.New("Bad Request")
	InternalError   = errors.New("Internal")
	ForbiddenError  = errors.New("Forbidden")
)

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func (e ValidationError) Decode() *map[string]interface{} {
	var res map[string]interface{}
	if err := json.Unmarshal([]byte(e.Message), &res); err != nil {
		return nil
	}
	return &res
}
