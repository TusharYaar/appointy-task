package main

import (
	"net/http"

	// Connection contains function to connect to database and defines collection
	"github.com/tusharyaar/task/connection"

	//Handlers contains all route handlers
	"github.com/tusharyaar/task/handlers"
)

func main() {
	// Initialize Mongo DB
	client,ctx,cancel := connection.Connect()
	// responds to GET /user/<user_id>
	http.HandleFunc("/user/",handlers.GetUser)

	// responds to POST /user
	http.HandleFunc("/user",handlers.CreateUser)

	// responds to Post /post
	http.HandleFunc("/post",handlers.CreatePost)

	// responds to GET /post/<post_id>
	http.HandleFunc("/post/",handlers.GetPost)
	
	// responds to GET /posts/user/<user_id>
	http.HandleFunc("/posts/user/",handlers.GetPostsByUser)

	// These Run after the server closes
	defer client.Disconnect(ctx)
	defer cancel()

	// Listener
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}



}


