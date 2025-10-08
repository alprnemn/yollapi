package validator

import (
	"errors"
	"fmt"
	"strings"

	validator "github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func ValidatePayload(payload any) error {

	fmt.Println("hello from validator")

	if err := Validate.Struct(payload); err != nil {

		var verr validator.ValidationErrors

		fields_tags := []string{}

		if errors.As(err, &verr) {
			fmt.Println("len errors: ", len(verr))
			for _, v := range verr {
				msg := fmt.Sprintf(
					"'%s' failed on the '%s' tag",
					v.Field(),
					v.Tag(),
				)
				fields_tags = append(fields_tags, msg)
			}

			// v := verr[0]
			return errors.New(strings.Join(fields_tags, ","))
		}

		return fmt.Errorf("validation failed: %w", err)
	}

	return nil
}
