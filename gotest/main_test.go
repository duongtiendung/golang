package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWelcome(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", welcome)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	s := writer.Body.String()

	if s != "Welcome to our website\n" {
		t.Errorf("Response body string is: %v", writer.Body.String())
	}
}
