package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/tusharyaar/task/handlers"

	// MongoDb

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Define User struct




func main() {

	// Initialize Mongo DB

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://task:1234567890@cluster0.a3iuv.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB")

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
    panic(err)
}
	fmt.Println("Ping Successfull to MongoDB")


// responds to GET /user
	http.HandleFunc("/",handlers.GetUser)

	// responds to POST /user
	http.HandleFunc("/user",handlers.CreateUser)




// These Run after the server closes
	defer cancel()
	defer client.Disconnect(ctx)


	// Listener
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}



}


