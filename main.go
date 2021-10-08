package main

import (
	"net/http"

	"github.com/tusharyaar/task/connection"
	"github.com/tusharyaar/task/handlers"
	// MongoDb
)

func main() {
	// Initialize Mongo DB
	client,ctx,cancel := connection.Connect()
	// responds to GET /user
	http.HandleFunc("/",handlers.GetUser)
	// responds to POST /user
	http.HandleFunc("/user",handlers.CreateUser)




// These Run after the server closes
	defer client.Disconnect(ctx)
	defer cancel()

	// Listener
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}



}


