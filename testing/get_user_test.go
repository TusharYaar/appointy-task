package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tusharyaar/task/connection"
	"github.com/tusharyaar/task/handlers"
)



func TestGetUser(t *testing.T) {
	client,ctx,cancel := connection.Connect()
	defer client.Disconnect(ctx)
	defer cancel()
	req, err := http.NewRequest("GET", "/user/6161447d56188e12db944c80", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusOK)
	}

	expected := `{"_id":"6161447d56188e12db944c80","name":"Tushar S Agrawal","email":"tusharsagrawal16@gmail.com","password":"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"}`
	if expected != rr.Body.String() {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestGetUserWrongMethod(t *testing.T) {
	req, err := http.NewRequest("POST", "/user/6161447d56188e12db944c80", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}

func TestGetUserWrongId(t *testing.T) {
	req, err := http.NewRequest("GET", "/user/6161447b944c80", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}


