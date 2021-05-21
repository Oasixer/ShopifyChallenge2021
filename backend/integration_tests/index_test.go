package integration_tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test that a GET request to / succeeds
func TestIndex(t *testing.T) {
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
