package gateway

import (
	user "github.com/crisnlopez/social-media-bkend/internal/user/models"
	"github.com/go-playground/validator/v10"
)

type validate *validator.Validate

func ValidateRequest(u user.UserRequest) error {
	err := validator.New().Struct(u)

	if err != nil {
		if invalid, ok := err.(*validator.InvalidValidationError); ok {
			return invalid
		}

		return err.(validator.ValidationErrors)
	}

	return nil
}
