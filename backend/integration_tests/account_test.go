package integration_tests

import (
	// "io/ioutil"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"fmt"
	// "strings"
	"log"
	"testing"
)
type SignUpResponse struct {
	Data struct {
		SignUp struct {
			Id string
			Email string
			CreatedAt string
			UpdatedAt string
		} `json: signUp`
	} `json: data`
	Errors []Error
}

type SignInResponse struct {
	Data struct {
		SignIn string
	} `json: data`
	Errors []Error
}

type ErrorExtensions struct {
	Code string
	Message string
}

type Error struct {
	Message string
	Path []string
	Extensions ErrorExtensions
	UpdatedAt string
}

type AccountData struct {
	Email string
	Username string
	Password string
}

var Account1 = AccountData{Email: "abc@abc.com", Password: "123", Username: "abc"}
var Account2 = AccountData{Email: "def@abc.com", Password: "456", Username: "123ABC"}
var InvalidPassword = "asiodj19024ui!!Q@#!U" // arbitrary incorrect password

// Test that a graceful failure occurs when trying to login to an account that does not exist
func TestSignInInvalidAccount(t *testing.T) {
	log.Print("INTEGRATION SIGN IN TEST INVALID ACC")
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
	}

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	fmt.Printf("response: %#v", w.Body.String())
	var signInResponse SignInResponse
	json.Unmarshal([]byte(w.Body.String()), &signInResponse)
	fmt.Printf("Data: %s", signInResponse.Data)
	fmt.Printf("%#v", signInResponse)

	if w.Code != http.StatusOK {
		t.Fail()
	}
	if signInResponse.Errors[0].Extensions.Code != "NotSignedUp"{
		t.Fail()
	}
}

// Test that a GET request to / succeeds
func TestSignUp(t *testing.T) {
	log.Print("INTEGRATION SIGN UP TEST")

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
	}

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	fmt.Printf("response: %#v\n\n", w.Body.String())
	var signUpResponse SignUpResponse	
	json.Unmarshal([]byte(w.Body.String()), &signUpResponse)
	// fmt.Printf("id: %s, createdAt: %s", signUpResponse.Data.Id, signUpResponse.Data.CreatedAt)
	fmt.Printf("Response formatted: %#v", signUpResponse)
	
	// log.Printf("\nsignUpResponse.Errors: %#v", signUpResponse.Errors)
	if signUpResponse.Errors != nil {
		t.Fail()
	}

	if signUpResponse.Data.SignUp.Email != Account1.Email{
		t.Fail()
	}

	if w.Code != http.StatusOK {
		t.Fail()
	}
}

// Test that a graceful failure occurs when trying to login to an invalid account
func TestSignInInvalidPassword(t *testing.T) {
	log.Print("\nINTEGRATION SIGN IN INVALID PW TEST")
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
	}

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	// fmt.Printf("response: %#v", w.Body.String())
	var signInResponse SignInResponse
	json.Unmarshal([]byte(w.Body.String()), &signInResponse)
	// fmt.Printf("Data: %s", signInResponse.Data)
	// fmt.Printf("%#v", signInResponse)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	if signInResponse.Errors[0].Extensions.Code != "PassBad"{
		t.Fail()
	}
}

// Test that a graceful failure occurs when trying to login to an invalid account
func TestSignIn(t *testing.T) {
	log.Print("\nINTEGRATION TEST SIGN IN")
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
	}

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	fmt.Printf("response: %#v", w.Body.String())
	var signInResponse SignInResponse
	json.Unmarshal([]byte(w.Body.String()), &signInResponse)
	// fmt.Printf("Data: %s", signInResponse.Data)
	// fmt.Printf("%#v", signInResponse)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	if signInResponse.Errors != nil{
		t.Fail()
	}
	fmt.Printf("%#v", signInResponse.Data)
}
