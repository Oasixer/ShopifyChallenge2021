package integration_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/Oasixer/ShopifyChallenge2021/utils"
	"github.com/google/uuid"
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

type SaveFileResponse struct {
	Data struct {
		File struct {
			Name string
			Uuid string
			Tags string
			FileExt string
		} `json:"saveFile"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type File struct {
	Name string
	Uuid string
	Tags string
	FileExt string
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

// helper function to confirm that graphql errors match expectations.
// if no errors are expected, simply pass in an empty array for expectCodes
func confirmErrors(errors []Error, expectCodes []string) bool {
	// confirm equal lengths
	if len(errors) != len(expectCodes) {
		log.Printf("Fail: expected errors:\n\t%#v,\n\tgot: %s\n", expectCodes, errors)
		return false
	}

	for _, err := range errors {
		if len(expectCodes) == 0 {
			log.Printf("Fail: expected no errors, got: %s\n", err)
			return false
		}

		errIndex := utils.Find(expectCodes, err.Extensions.Code)
		if errIndex == -1 { // if error not found in expected errors
			log.Printf("Fail: expected errors:\n\t%#v,\n\tgot: %s\n", expectCodes, err)
			return false
		} else {
			// found error, remove from slice so it doesn't get double counted
			expectCodes[errIndex] = expectCodes[len(expectCodes)-1] // Copy last element to index i.
			expectCodes[len(expectCodes)-1] = ""                    // Erase last element (write zero value).
			expectCodes = expectCodes[:len(expectCodes)-1]          // Truncate slice.
		}
	}
	return true
}

var ( // TODO probably move some of these to a shared file and make the rest private
	Account1        = AccountData{Email: "abc@abc.com", Password: "123", Username: "abc"}
	Account2        = AccountData{Email: "def@abc.com", Password: "456", Username: "123ABC"}
	InvalidPassword = "asiodj19024ui!!Q@#!U" // arbitrary incorrect password
	Token           string
	fileUuid, _     = uuid.NewUUID()
	fileUuid2, _    = uuid.NewUUID()
	tags            = []string{"tag1", "tag2"}
	File1           = File{
		Name: "filename111",
		Uuid: fileUuid.String(),
		Tags: strings.Join(tags, ","),
		FileExt: "png",
	}
	File2 = File{
		Name: "filename222",
		Uuid: fileUuid2.String(),
		Tags: "",
		FileExt: "png",
	}
)

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
		log.Printf("The HTTP request failed with error %s\n", err)
		t.Fail()
	}

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		log.Printf("The request returned bad code: %d\n", w.Code)
		t.Fail()
	}

	var signInResponse SignInResponse
	json.Unmarshal([]byte(w.Body.String()), &signInResponse)

	// expect NotSignedUp error
	if !confirmErrors(signInResponse.Errors, []string{"NotSignedUp"}) {
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
		log.Printf("The HTTP request failed with error %s\n", err)
		t.Fail()
	}

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		log.Printf("The request returned bad code: %d\n", w.Code)
		t.Fail()
	}

	var signUpResponse SignUpResponse
	json.Unmarshal([]byte(w.Body.String()), &signUpResponse)

	// expect no error
	if !confirmErrors(signUpResponse.Errors, []string{}) {
		t.Fail()
	}

	if signUpResponse.Data.SignUp.Email != Account1.Email {
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
		log.Printf("The HTTP request failed with error %s\n", err)
		t.Fail()
	}

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		log.Printf("The request returned bad code: %d\n", w.Code)
		t.Fail()
	}

	var signInResponse SignInResponse
	json.Unmarshal([]byte(w.Body.String()), &signInResponse)

	// expect PassBad error
	if !confirmErrors(signInResponse.Errors, []string{"PassBad"}) {
		t.Fail()
	}
}

// test that a sign in occurs succesfully with valid credentials.
// stores the jwt token returned on success in variable 'Token', to authenticate the next test.
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
		log.Printf("The HTTP request failed with error %s\n", err)
		t.Fail()
	}

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		log.Printf("The request returned bad code: %d\n", w.Code)
		t.Fail()
	}

	var signInResponse SignInResponse
	json.Unmarshal([]byte(w.Body.String()), &signInResponse)

	// expect no error
	if !confirmErrors(signInResponse.Errors, []string{}) {
		t.Fail()
	}
	Token = signInResponse.Data.SignIn
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
		log.Printf("The HTTP request failed with error %s\n", err)
		t.Fail()
	}

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		log.Printf("The request returned bad code: %d\n", w.Code)
		t.Fail()
	}

	var getUserResponse GetUserResponse
	json.Unmarshal([]byte(w.Body.String()), &getUserResponse)

	// expect NotAuth error
	if !confirmErrors(getUserResponse.Errors, []string{"NotAuth"}) {
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

	req.Header.Set("Authorization", Token)

	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
	}

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		log.Printf("The request returned bad code: %d\n", w.Code)
		t.Fail()
	}

	var getUserResponse GetUserResponse
	json.Unmarshal([]byte(w.Body.String()), &getUserResponse)

	// expect no error
	if !confirmErrors(getUserResponse.Errors, []string{}) {
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

func TestSaveFile(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := getRouter()

	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		   mutation{ 
				 saveFile(uuid: "%s", name: "%s", tags: "%s", fileExt: "%s"){
					 uuid
					 name
					 tags
				 }
		   }`, File1.Uuid, File1.Name, File1.Tags, File1.FileExt),
	}
	jsonValue, _ := json.Marshal(jsonData)
	req, err := http.NewRequest("POST", "/graphql", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", Token)

	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
		t.Fail()
	}

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		log.Printf("The request returned bad code: %d\n", w.Code)
		t.Fail()
	}

	var saveFileResponse SaveFileResponse
	json.Unmarshal([]byte(w.Body.String()), &saveFileResponse)

	// expect no error
	if !confirmErrors(saveFileResponse.Errors, []string{}) {
		t.Fail()
	}

	if saveFileResponse.Data.File.Uuid != File1.Uuid {
		log.Printf("Fail: expected Uuid: %#v, got: %#v\n", File1.Uuid, saveFileResponse.Data.File.Uuid)
		log.Printf("Response unformatted: %#v\n\n", w.Body.String())
		log.Printf("Response formatted: %#v\n\n", saveFileResponse)
		t.Fail()
	}
	if saveFileResponse.Data.File.Tags != File1.Tags {
		log.Printf("Fail: expected Tags: %#v, got: %#v\n", File1.Tags, saveFileResponse.Data.File.Tags)
		log.Printf("Response unformatted: %#v\n\n", w.Body.String())
		log.Printf("Response formatted: %#v\n\n", saveFileResponse)
		t.Fail()
	}

	// now confirm that it persisted to the db
	getUserJsonData := map[string]string{
		"query": `{ getUser {
				files {
					name
					tags
					uuid
				}
		 	}
		 }`,
	}
	getUserJsonValue, _ := json.Marshal(getUserJsonData)
	req, err = http.NewRequest("POST", "/graphql", bytes.NewBuffer(getUserJsonValue))
	req.Header.Set("Authorization", Token)

	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
		t.Fail()
	}

	// Create the service and process the above request.
	w_ := httptest.NewRecorder()
	r.ServeHTTP(w_, req)

	if w_.Code != http.StatusOK {
		log.Printf("The request returned bad code: %d\n", w_.Code)
		t.Fail()
	}

	var getUserResponse GetUserResponse
	json.Unmarshal([]byte(w_.Body.String()), &getUserResponse)

	// expect no error
	if !confirmErrors(getUserResponse.Errors, []string{}) {
		t.Fail()
	}

	// expect a file
	if len(getUserResponse.Data.User.Files) != 1 {
		log.Println("Fail: file did not get associated to the user")
		log.Printf("response unformatted: %#v\n\n", w_.Body.String())
		t.Fail()
	}
}

// make sure that the relationship handles multiple files
func TestSaveSecondFile(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := getRouter()

	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		   mutation{ 
				 saveFile(uuid: "%s", name: "%s", tags: "%s", fileExt: "%s"){
					 uuid
					 name
					 tags
				 }
		   }`, File2.Uuid, File2.Name, File2.Tags, File2.FileExt),
	}
	jsonValue, _ := json.Marshal(jsonData)
	req, err := http.NewRequest("POST", "/graphql", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", Token)

	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
		t.Fail()
	}

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		log.Printf("The request returned bad code: %d\n", w.Code)
		t.Fail()
	}

	var saveFileResponse SaveFileResponse
	json.Unmarshal([]byte(w.Body.String()), &saveFileResponse)

	// expect no error
	if !confirmErrors(saveFileResponse.Errors, []string{}) {
		t.Fail()
	}

	if saveFileResponse.Data.File.Uuid != File2.Uuid {
		log.Printf("Fail: expected Uuid: %#v, got: %#v\n", File2.Uuid, saveFileResponse.Data.File.Uuid)
		log.Printf("Response unformatted: %#v\n\n", w.Body.String())
		log.Printf("Response formatted: %#v\n\n", saveFileResponse)
		t.Fail()
	}

	// now confirm that it persisted to the db
	getUserJsonData := map[string]string{
		"query": `{ getUser {
				files {
					name
					tags
					uuid
					fileExt
				}
		 	}
		 }`,
	}
	getUserJsonValue, _ := json.Marshal(getUserJsonData)
	req, err = http.NewRequest("POST", "/graphql", bytes.NewBuffer(getUserJsonValue))
	req.Header.Set("Authorization", Token)

	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
		t.Fail()
	}

	// Create the service and process the above request.
	w_ := httptest.NewRecorder()
	r.ServeHTTP(w_, req)

	if w_.Code != http.StatusOK {
		log.Printf("The request returned bad code: %d\n", w_.Code)
		t.Fail()
	}

	var getUserResponse GetUserResponse
	json.Unmarshal([]byte(w_.Body.String()), &getUserResponse)

	// expect no error
	if !confirmErrors(getUserResponse.Errors, []string{}) {
		t.Fail()
	}

	// expect a file
	if len(getUserResponse.Data.User.Files) != 2 {
		log.Println("Fail: file did not get associated to the user WTF bro")
		log.Printf("response unformatted: %#v\n\n", w_.Body.String())
		t.Fail()
	}

	// log.Printf("response unformatted: %#v\n\n", w_.Body.String())
	// if (getUserResponse.Data.User.Files[1].Uuid != )
}
