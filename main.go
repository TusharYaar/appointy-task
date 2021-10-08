package main

import (
	"fmt"
	"net/http"

	"github.com/tusharyaar/task/connection"
	"github.com/tusharyaar/task/handlers"
	"github.com/tusharyaar/task/models"
	// MongoDb
)

func main() {
	var user models.User
	// Initialize Mongo DB
	client,ctx,cancel := connection.Connect()
	// err := connection.UserCollection.FindOne(context.TODO(), bson.D{{"email","email@gmail.com"}}).Decode(&user)
	

	fmt.Println(user)
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


