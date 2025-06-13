package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New(validator.WithRequiredStructEnabled())

func ValidateStruct(s interface{}) error {
	err := Validate.Struct(s)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		if len(validationErrors) == 0 {
			return nil
		}

		firstErr := validationErrors[0]
		return fmt.Errorf("%s %s %s", firstErr.Field(), firstErr.Tag(), firstErr.Param())
	}

	return err
}
