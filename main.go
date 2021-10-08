package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tusharyaar/task/connection"
	"github.com/tusharyaar/task/handlers"
	"github.com/tusharyaar/task/models"

	// MongoDb
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	var user models.User
	// Initialize Mongo DB
	client,ctx := connection.Connect()
	err := connection.UserCollection.FindOne(context.TODO(), bson.D{{"email","email@gmail.com"}}).Decode(&user)
	
	if err != nil {
		fmt.Println("Error while fetching user")
	}

	fmt.Println(user)
	// responds to GET /user
	http.HandleFunc("/",handlers.GetUser)
	// responds to POST /user
	http.HandleFunc("/user",handlers.CreateUser)




// These Run after the server closes
	defer client.Disconnect(ctx)


	// Listener
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}



}


