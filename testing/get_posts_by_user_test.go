package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tusharyaar/task/connection"
	"github.com/tusharyaar/task/handlers"
)



func TestGetPostsByUser(t *testing.T) {
	client,ctx,cancel := connection.Connect()
	defer client.Disconnect(ctx)
	defer cancel()
	req, err := http.NewRequest("GET", "/posts/user/6161447d56188e12db944c80", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetPostsByUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusOK)
	}
}
func TestGetPostsByUserWithPagination(t *testing.T) {
	client,ctx,cancel := connection.Connect()
	defer client.Disconnect(ctx)
	defer cancel()
	req, err := http.NewRequest("GET", "/posts/user/6161447d56188e12db944c80", nil)
	if err != nil {
		t.Fatal(err)
	}
	query := req.URL.Query()
	query.Add("page", "2")
	query.Add("limit", "2")
	req.URL.RawQuery = query.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetPostsByUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusOK)
	}
}
func TestGetPostsByUserWithPaginationExcededPage(t *testing.T) {
	client,ctx,cancel := connection.Connect()
	defer client.Disconnect(ctx)
	defer cancel()
	req, err := http.NewRequest("GET", "/posts/user/6161447d56188e12db944c80", nil)
	if err != nil {
		t.Fatal(err)
	}
	query := req.URL.Query()
	query.Add("page", "12")
	query.Add("limit", "2")
	req.URL.RawQuery = query.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetPostsByUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusNotFound)
	}
}

func TestGetPostsByUserDoesNotExists(t *testing.T) {
	client,ctx,cancel := connection.Connect()
	defer client.Disconnect(ctx)
	defer cancel()
	req, err := http.NewRequest("GET", "/posts/user/6161447d56188e12db944c81", nil)

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetPostsByUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusNotFound)
	}
}

func TestGetPostsByUserWrongMethod(t *testing.T) {
	req, err := http.NewRequest("POST", "/posts/user/6161447d56188e12db944c81", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetPostsByUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}


