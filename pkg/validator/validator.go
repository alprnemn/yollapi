package validator

import (
	"errors"
	"fmt"

	validator "github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func ValidatePayload(payload any) error {

	if err := Validate.Struct(payload); err != nil {

		var verr validator.ValidationErrors

		if errors.As(err, &verr) {
			first := verr[0]
			msg := fmt.Sprintf(
				"Field validation for '%s' failed on the '%s' tag",
				first.Field(),
				first.Tag(),
			)
			return errors.New(msg)
		}

		return fmt.Errorf("validation failed: %w", err)
	}

	return nil
}
