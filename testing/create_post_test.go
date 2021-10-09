package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tusharyaar/task/connection"
	"github.com/tusharyaar/task/handlers"
)


func TestCreatePost(t *testing.T) {
	client,ctx,cancel := connection.Connect()
	defer client.Disconnect(ctx)
	defer cancel()
	// 61619a2d51169d1757ad2021 belongs to test user
	var jsonStr = []byte(`{"user_id": "61619a2d51169d1757ad2021","caption": "This image is created during testing img","image_url": "testimageurl"}`)

	req, err := http.NewRequest("POST", "/post", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreatePost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusCreated)
	}

}


func TestCreatePostWrongMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/post", nil)
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

func TestCreatePostWrongHeader(t *testing.T) {
	var jsonStr = []byte(`{"user_id": "61619a2d51169d1757ad2021","caption": "This image is created during testing img","image_url": "testimageurl"}`)
	
	req, err := http.NewRequest("POST", "/post", bytes.NewBuffer(jsonStr))
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
func TestCreatePostIncompleteDetails(t *testing.T) {

	var jsonStr = []byte(`{"caption": "This image is created during testing img","image_url": "testimageurl"}`)

	req, err := http.NewRequest("POST", "/post", bytes.NewBuffer(jsonStr))
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


func TestCreatePostUserDoesNotExist(t *testing.T) { 
	client,ctx,cancel := connection.Connect()
	defer client.Disconnect(ctx)
	defer cancel()
	// 61619a2d51169d1757ad2021 belongs to test user, change the user id to a non existing user (last digit)
	var jsonStr = []byte(`{"user_id": "61619a2d51169d1757ad2029","caption": "This image is created during testing img","image_url": "testimageurl"}`)

	req, err := http.NewRequest("POST", "/post", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreatePost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusNotFound)
	}

}
