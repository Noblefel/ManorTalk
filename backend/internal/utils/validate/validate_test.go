package validate

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

var validateStructTests = []struct {
	name           string
	input          interface{}
	expectingError bool
}{
	{
		name: "validateStruct-success",
		input: struct {
			Email    string `validate:"required,email"`
			Password string `validate:"required,min=8"`
		}{
			Email:    "test@example.com",
			Password: "password",
		},
		expectingError: false,
	},
	{
		name: "validateStruct-fail-required",
		input: struct {
			Name string `validate:"required"`
		}{
			Name: "",
		},
		expectingError: true,
	},
	{
		name: "validateStruct-fail-email",
		input: struct {
			Email string `validate:"email"`
		}{
			Email: "not-an-email",
		},
		expectingError: true,
	},
	{
		name: "validateStruct-fail-min",
		input: struct {
			Password string `validate:"min=8"`
		}{
			Password: "x",
		},
		expectingError: true,
	},
	{
		name: "validateStruct-fail-max",
		input: struct {
			Password string `validate:"max=2"`
		}{
			Password: "longpassword",
		},
		expectingError: true,
	},
}

func TestValidateStruct(t *testing.T) {
	for _, tt := range validateStructTests {
		err := Struct(tt.input)

		if tt.expectingError == false && err != nil {
			t.Errorf("%s expected no error, but got %v", tt.name, err)
		}

		if tt.expectingError == true && err == nil {
			t.Errorf("%s expecting some error, but got nil", tt.name)
		}
	}
}

var getMessageTests = []struct {
	name        string
	input       interface{}
	expectedMsg string
}{
	{
		name: "validateStruct-correct-email",
		input: struct {
			Field string `validate:"email"`
		}{
			Field: "test@example.com",
		},
	},
	{
		name: "validateStruct-correct-required",
		input: struct {
			Field2 string `validate:"required"`
		}{
			Field2: "",
		},
		expectedMsg: "Field2 cannot be empty",
	},
	{
		name: "validateStruct-correct-min",
		input: struct {
			Field3 string `validate:"min=3"`
		}{
			Field3: "a",
		},
		expectedMsg: "Field3 must be atleast 3 characters",
	},
	{
		name: "validateStruct-correct-max",
		input: struct {
			Field4 string `validate:"max=4"`
		}{
			Field4: "abcdefghij",
		},
		expectedMsg: "Field4 must not exceed 4 characters",
	},
}

func TestGetMessage(t *testing.T) {
	for _, tt := range getMessageTests {
		validate := validator.New(validator.WithRequiredStructEnabled())

		var msg string

		if err := validate.Struct(tt.input); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				msg = getMessage(err)
			}
		}

		if tt.expectedMsg != msg {
			t.Errorf("%s did not get the correct message, wanted %s but got %s", tt.name, tt.expectedMsg, msg)
		}
	}
}
