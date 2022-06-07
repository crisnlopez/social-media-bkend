package gateway_test

import (
	"testing"

	"github.com/crisnlopez/social-media-bkend/internal/user/gateway"
	user "github.com/crisnlopez/social-media-bkend/internal/user/models"
	"github.com/crisnlopez/social-media-bkend/internal/util"
	"github.com/go-playground/validator/v10"
)

func TestValidateStruct(t *testing.T) {
	tests := []struct {
		name          string
		input         user.UserRequest
		expectedError bool // True if err == validator.ValidationErrors
	}{
		{
			name: "OK",
			input: user.UserRequest{
				Email: util.RandomEmail(),
				Pass:  util.RandomPass(),
				Nick:  util.RandomNick(),
				Name:  util.RandomName(),
				Age:   util.RandomAge(),
			},
			expectedError: false,
		},
		{
			name: "Invalid Email",
			input: user.UserRequest{
				Email: "fakeinvalidemail.com",
				Pass:  util.RandomPass(),
				Nick:  util.RandomNick(),
				Name:  util.RandomName(),
				Age:   util.RandomAge(),
			},
			expectedError: true,
		},
		{
			name: "Email required",
			input: user.UserRequest{
				Pass: util.RandomPass(),
				Nick: util.RandomNick(),
				Name: util.RandomName(),
				Age:  util.RandomAge(),
			},
			expectedError: true,
		},
		{
			name: "Pass required",
			input: user.UserRequest{
				Email: util.RandomEmail(),
				Nick:  util.RandomNick(),
				Name:  util.RandomName(),
				Age:   util.RandomAge(),
			},
			expectedError: true,
		},
		{
			name: "Nick required",
			input: user.UserRequest{
				Email: util.RandomEmail(),
				Pass:  util.RandomPass(),
				Name:  util.RandomName(),
				Age:   util.RandomAge(),
			},
			expectedError: true,
		},
		{
			name: "Name required",
			input: user.UserRequest{
				Email: util.RandomEmail(),
				Pass:  util.RandomPass(),
				Nick:  util.RandomNick(),
				Age:   util.RandomAge(),
			},
			expectedError: true,
		},
		{
			name: "Age required",
			input: user.UserRequest{
				Email: util.RandomEmail(),
				Pass:  util.RandomPass(),
				Nick:  util.RandomNick(),
				Name:  util.RandomName(),
			},
			expectedError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotErr := false
			err := gateway.ValidateRequest(tc.input)

			if err != nil {
				if _, ok := err.(validator.ValidationErrors); ok {
					gotErr = true
				} else {
					t.Logf("expected ValidationErrors. Got: %v", err.Error())
				}
			}

			if tc.expectedError != gotErr {
				t.Logf("expectedError: %v Got: %v\n", tc.expectedError, gotErr)
			}
		})
	}
}
