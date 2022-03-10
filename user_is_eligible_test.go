package main

import (
	"errors"
	"testing"
)

var test = []struct{
  email string
  password string
  age int
  expetecErr error
}{
  {
      email: "test@example.com",
      password: "1234",
      age: 19,
      expetecErr: nil,
    }, 
  {
      email: "",
      password: "12345",
      age: 18,
      expetecErr: errors.New("email can't be empty"),
    },
  {
      email: "test@example.com",
      password: "",
      age: 18,
      expetecErr: errors.New("password can't be empty"),
    },
  {
      email: "test@example.com",
      password: "12345",
      age: 16,
      expetecErr: errors.New("age can't be at least 18 years old"),
    },
  }

func TestUserIsEligible(t *testing.T) {
  for _, tt := range test {
    err := userIsEligible(tt.email, tt.password, tt.age)

    errString := ""
    expectedErrString := ""

    if err != nil {
      errString = err.Error()
    }

    if tt.expetecErr != nil {
      expectedErrString = tt.expetecErr.Error()
    }
      
    if errString != expectedErrString {
      t.Errorf("got %s, want %s",errString, expectedErrString)
    }
  }
}
