package test

import (
	"mindsdb_go_sdk/connectors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRestAPI_Login(t *testing.T) {
	email := "test@example.com"
	password := "password123"

	// Create a test server to mock the login endpoint
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Unexpected request method %s", r.Method)
		}
		if r.URL.Path != "/cloud/login" {
			t.Errorf("Unexpected request path %s", r.URL.Path)
		}
		// Verify the request body contains the email and password
		err := r.ParseForm()
		if err != nil {
			t.Errorf("Failed to parse form: %v", err)
		}
		if r.PostForm.Get("email") != email {
			t.Errorf("Unexpected email %s", r.PostForm.Get("email"))
		}
		if r.PostForm.Get("password") != password {
			t.Errorf("Unexpected password %s", r.PostForm.Get("password"))
		}
		// Send a successful response
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	api := connectors.RestAPI{
		url:      Url,
		email:    Email,
		password: password,
		session:  &http.Client{},
	}
	err := api.login()
	if err != nil {
		t.Errorf("Failed to log in: %v", err)
	}
}
