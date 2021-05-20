package integration_tests

import (
	// "io/ioutil"
	"net/http"
	"net/http/httptest"

	// "net/url"
	// "strconv"
	// "strings"
	"log"
	"testing"
)

// Test that a GET request to / succeeds
func TestIndex(t *testing.T) {
	log.Print("PLEASE RUN THIS INDEX_TEST")
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := getRouter()

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/api/", nil)

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}
}
