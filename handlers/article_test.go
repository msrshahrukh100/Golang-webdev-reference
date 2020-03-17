package handlers

import (
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "strings"
  "testing"
  "github.com/msrshahrukh100/Golang-webdev-reference/tests"
)

// Test that a GET request to the home page returns the home page with
// the HTTP code 200 for an unauthenticated user

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := tests.GetRouter(true)
	r.GET("/", ShowIndexPage)
	req, _ := http.NewRequest("GET", "/", nil)

	tests.TestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)

		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}