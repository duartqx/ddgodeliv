package errors

import (
	"encoding/json"
	"errors"
)

var (
	BadRequestError = errors.New("Bad Request")
	ForbiddenError  = errors.New("Forbidden")
	InternalError   = errors.New("Internal")
	NotFoundError   = errors.New("Not Found")
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
