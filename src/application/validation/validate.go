package validation

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type validationErr struct {
	Tag   string
	Value interface{}
}

type Validator struct {
	*validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		Validate: validator.New(),
	}
}

func (v Validator) Decode(errs error) *map[string]interface{} {

	validationErrors := map[string]interface{}{}

	for _, err := range errs.(validator.ValidationErrors) {
		validationErrors[err.Field()] = validationErr{
			Tag:   err.Tag(),
			Value: err.Value(),
		}
	}

	return &validationErrors
}

func (v Validator) JSON(errs error) *[]byte {
	res, _ := json.Marshal(v.Decode(errs))
	return &res
}
