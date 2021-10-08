package connection

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollection *mongo.Collection
var PostCollection *mongo.Collection

func Connect() (*mongo.Client, context.Context, context.CancelFunc) {
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
	// err = client.Ping(ctx, readpref.Primary())
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Ping Successfull to MongoDB")
	UserCollection = client.Database("taskDB").Collection("user")
	PostCollection = client.Database("taskDB").Collection("post")

	return client,ctx, cancel
}