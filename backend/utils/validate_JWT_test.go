package utils

import (
	"testing"
)

func TestValidateJWT(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIxLTA1LTI4VDE2OjUyOjIyLjk4MDI4My0wNDowMCIsInVzZXJJRCI6IjEifQ.upDKJ-4Mp9Bg-OyTJ5VuntcijY3A-8POZr5Uh1mRZmI"
	userID, err := ValidateJWT(&tokenString)
	if err != nil {
		t.Error(err)
	}

	if userID == nil {
		t.Errorf("Token is expired or something wrong")
	}

	t.Log(userID)
}
