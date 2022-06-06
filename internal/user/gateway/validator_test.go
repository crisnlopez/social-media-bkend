package gateway_test

import (
	"testing"

	"github.com/crisnlopez/social-media-bkend/internal/user/gateway"
	user "github.com/crisnlopez/social-media-bkend/internal/user/models"
	"github.com/go-playground/validator/v10"
)

func TestValidateStruct(t *testing.T) {
  u := user.UserRequest{
    Email: "test@example.com",
    Pass: "1234",
    Nick: "nick test",
    Name: "Name test",
    Age: 20,
  }

	err := gateway.ValidateRequest(u)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			t.Fatal("Error InvalidValidationError", err)
		}

		if _, ok := err.(validator.ValidationErrors); ok {
			t.Fatal("Error ValidationErrors:",err)
		}
	} 
}
