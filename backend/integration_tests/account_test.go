package integration_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"fmt"
	// "log"
	"testing"
)

type SignUpResponse struct {
	Data struct {
		SignUp struct {
			Id        string
			Email     string
			CreatedAt string
			UpdatedAt string
		} `json:"signUp"`
	} `json:"data"`
	Errors []Error
}

type SignInResponse struct {
	Data struct {
		SignIn string
	} `json: data`
	Errors []Error
}

type GetUserResponse struct {
	Data struct {
		User struct {
			Email     string
			Username  string
			CreatedAt string
			UpdatedAt string
			Files     []File `json:"files"`
		} `json:"getUser"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type File struct {
	Name string
	Uuid string
	Tags string
}

type ErrorExtensions struct {
	Code    string
	Message string
}

type Error struct {
	Message    string
	Path       []string
	Extensions ErrorExtensions
	UpdatedAt  string
}

type AccountData struct {
	Email    string
	Username string
	Password string
}

var Account1 = AccountData{Email: "abc@abc.com", Password: "123", Username: "abc"}
var Account2 = AccountData{Email: "def@abc.com", Password: "456", Username: "123ABC"}
var InvalidPassword = "asiodj19024ui!!Q@#!U" // arbitrary incorrect password
var token string

// Test that a graceful failure occurs when trying to login to an account that does not exist
func TestSignInInvalidAccount(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := getRouter()

	jsonData := map[string]string{
		"query": fmt.Sprintf(`mutation{ 
			 signIn(email: "%s", password: "%s")
		 }`, Account1.Email, Account1.Password),
	}
	jsonValue, _ := json.Marshal(jsonData)
	req, err := http.NewRequest("POST", "/graphql", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		t.Fail()
	}

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	var signInResponse SignInResponse
	json.Unmarshal([]byte(w.Body.String()), &signInResponse)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	if signInResponse.Errors[0].Extensions.Code != "NotSignedUp" {
		t.Fail()
	}
}

// Test that signing up w/ the details of the Account1 struct succeeds
func TestSignUp(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := getRouter()

	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		   mutation{ 
				 signUp(email: "%s", password: "%s", username: "%s"){
					 id,
					 email,
					 createdAt,
					 updatedAt
				 }
		   }`, Account1.Email, Account1.Password, Account1.Username),
	}
	jsonValue, _ := json.Marshal(jsonData)
	req, err := http.NewRequest("POST", "/graphql", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		t.Fail()
	}

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	var signUpResponse SignUpResponse
	json.Unmarshal([]byte(w.Body.String()), &signUpResponse)

	if signUpResponse.Errors != nil {
		t.Fail()
	}

	if signUpResponse.Data.SignUp.Email != Account1.Email {
		t.Fail()
	}

	if w.Code != http.StatusOK {
		t.Fail()
	}
}

// Test that a graceful failure occurs when trying to login w/ an invalid password
func TestSignInInvalidPassword(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := getRouter()

	jsonData := map[string]string{
		"query": fmt.Sprintf(`mutation{ 
			 signIn(email: "%s", password: "%s")
		 }`, Account1.Email, InvalidPassword),
	}
	jsonValue, _ := json.Marshal(jsonData)
	req, err := http.NewRequest("POST", "/graphql", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		t.Fail()
	}

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	var signInResponse SignInResponse
	json.Unmarshal([]byte(w.Body.String()), &signInResponse)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	if signInResponse.Errors[0].Extensions.Code != "PassBad" {
		t.Fail()
	}
}

// test that a sign in occurs succesfully with valid credentials.
// stores the jwt token returned on success in variable 'token', to authenticate the next test.
func TestSignIn(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := getRouter()

	jsonData := map[string]string{
		"query": fmt.Sprintf(`mutation{ 
			 signIn(email: "%s", password: "%s")
		 }`, Account1.Email, Account1.Password),
	}
	jsonValue, _ := json.Marshal(jsonData)
	req, err := http.NewRequest("POST", "/graphql", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		t.Fail()
	}

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	var signInResponse SignInResponse
	json.Unmarshal([]byte(w.Body.String()), &signInResponse)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	if signInResponse.Errors != nil {
		t.Fail()
	}
	token = signInResponse.Data.SignIn
}

func TestGetUserNotAuth(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := getRouter()

	jsonData := map[string]string{
		"query": `{ getUser {
				files {
					name,
					tags,
					uuid
				}
		 	}
		 }`,
	}
	jsonValue, _ := json.Marshal(jsonData)
	req, err := http.NewRequest("POST", "/graphql", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		t.Fail()
	}

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	var getUserResponse GetUserResponse
	json.Unmarshal([]byte(w.Body.String()), &getUserResponse)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	// Errors should be a NotAuth error
	if getUserResponse.Errors == nil {
		t.Fail()
	}
	if getUserResponse.Errors[0].Extensions.Code != "NotAuth" {
		t.Fail()
	}
}

func TestGetUser(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := getRouter()

	jsonData := map[string]string{
		"query": `{ getUser {
				email
				username
				createdAt
				updatedAt
				files {
					name
					tags
					uuid
				}
		 	}
		 }`,
	}
	jsonValue, _ := json.Marshal(jsonData)
	req, err := http.NewRequest("POST", "/graphql", bytes.NewBuffer(jsonValue))

	req.Header.Set("Authorization", token)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	var getUserResponse GetUserResponse
	json.Unmarshal([]byte(w.Body.String()), &getUserResponse)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	if getUserResponse.Errors != nil {
		t.Fail()
	}

	if getUserResponse.Data.User.Username != Account1.Username {
		t.Fail()
	}

	if getUserResponse.Data.User.Email != Account1.Email {
		t.Fail()
	}

	if len(getUserResponse.Data.User.Files) != 0 {
		t.Fail()
	}
}
