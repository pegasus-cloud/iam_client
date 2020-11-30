package utility

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

const (
	invalidErrMsg = "The error message does not could be converted"
)

var errorList map[string]string

func init() {
	errorList = map[string]string{
		"UserID:min":            "The length of userId is out of range",
		"UserID:max":            "The length of userId is out of range",
		"UserID:required_with":  "The userId is required, if force is true",
		"UserID:required":       "The userId is required",
		"DisplayName:min":       "The length of displayname is out of range",
		"DisplayName:max":       "The length of displayname is out of range",
		"DisplayName:required":  "The displayname is required",
		"Password:required":     "The password is required",
		"GroupID:min":           "The length of groupId is out of range",
		"GroupID:max":           "The length of groupId is out of range",
		"GroupID:required_with": "The groupId is required, if force is true",
		"GroupID:required":      "The groupId is required",
	}
}

// ConvertError ...
func ConvertError(input error) (output error) {
	validationErrors, ok := input.(validator.ValidationErrors)
	if !ok {
		return errors.New(invalidErrMsg)
	}

	validationError := validationErrors[0]
	fmt.Println(validationError.Field(), validationError.Tag())
	return errors.New(errorList[fmt.Sprintf("%s:%s", validationError.Field(), validationError.Tag())])
}
