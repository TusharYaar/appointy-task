package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tusharyaar/task/connection"
	"github.com/tusharyaar/task/handlers"
)


func TestCreateUser(t *testing.T) {
	client,ctx,cancel := connection.Connect()
	defer client.Disconnect(ctx)
	defer cancel()
	
	var jsonStr = []byte(`{"name":"test2","email":"test2@test.com","password":"test2"}`)

	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusCreated)
	}

}

func TestCreateUserExsists(t *testing.T) {
	client,ctx,cancel := connection.Connect()
	defer client.Disconnect(ctx)
	defer cancel()
	
	var jsonStr = []byte(`{"name":"test","email":"test@test.com","password":"test"}`)

	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusConflict {
		t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusConflict)
	}

}

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
	var jsonStr = []byte(`{"name":"xyz","email":"test@test.com","password":"helloworld"}`)

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