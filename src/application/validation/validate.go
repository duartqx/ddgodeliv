package validation

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"

	e "ddgodeliv/common/errors"
)

type validationErr struct {
	Tag   string      `json:"tag"`
	Value interface{} `json:"value"`
}

type Validator struct {
	*validator.Validate
}

func NewValidator() *Validator {
	v := validator.New()

	v.RegisterValidation("future", func(fl validator.FieldLevel) bool {
		t, ok := fl.Field().Interface().(time.Time)
		if !ok {
			return false
		}
		return t.After(time.Now())
	})

	return &Validator{Validate: v}
}

func (v Validator) Decode(errs error) *map[string]interface{} {

	validationErrors := map[string]interface{}{}

	for _, err := range errs.(validator.ValidationErrors) {
		validationErrors[strings.ToLower(err.Field())] = validationErr{
			Tag:   err.Tag(),
			Value: err.Value(),
		}
	}

	return &map[string]interface{}{"error": validationErrors}
}

func (v Validator) JSON(errs error) *[]byte {
	res, _ := json.Marshal(v.Decode(errs))
	return &res
}

func (v Validator) ValidateStructJson(s interface{}) *[]byte {
	if errs := v.Struct(s); errs != nil {
		return v.JSON(errs)
	}
	return nil
}

func (v Validator) ValidateStruct(s interface{}) error {
	if errs := v.Struct(s); errs != nil {
		return &e.ValidationError{Message: string(*v.JSON(errs))}
	}
	return nil
}

func (v Validator) ValidateVar(i interface{}, tag string) error {
	if errs := v.Var(i, tag); errs != nil {
		return fmt.Errorf(string(*v.JSON(errs)))
	}
	return nil
}

func (v Validator) ValidateVarJson(i interface{}, tag string) *[]byte {
	if errs := v.Var(i, tag); errs != nil {
		return v.JSON(errs)
	}
	return nil
}
