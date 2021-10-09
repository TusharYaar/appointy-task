package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tusharyaar/task/handlers"
)

func TestCreateUserWrongMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}

}

func TestCreateUserWrongHeader(t *testing.T) {
	var jsonStr = []byte(`{"name":"xyz","email":"test@test.com"},"password":"helloworld"`)

	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusUnsupportedMediaType {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnsupportedMediaType)
	}

}
func TestCreateUserIncompleteDetails(t *testing.T) {

	var jsonStr = []byte(`{"name":"xyz","email":"test@test.com"}`)

	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}