package validate

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type InputError map[string]string

var InputErrMessage = map[string]string{
	"email":    "$field is not a valid email",
	"required": "$field cannot be empty",
	"min":      "$field must be atleast $param characters",
	"max":      "$field must not exceed $param characters",
}

func Struct(i interface{}) *InputError {
	validate := validator.New(validator.WithRequiredStructEnabled())

	errors := InputError{}

	if err := validate.Struct(i); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			k := strings.ToLower(err.Field())
			errors[k] = getMessage(err)
		}

		return &errors
	}

	return nil
}

func getMessage(err validator.FieldError) string {
	msg := InputErrMessage[err.Tag()]
	msg = strings.Replace(msg, "$field", err.Field(), 1)
	msg = strings.Replace(msg, "$param", err.Param(), 1)

	return msg
}
