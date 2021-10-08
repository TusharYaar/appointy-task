package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/tusharyaar/task/connection"
	"github.com/tusharyaar/task/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePost(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	if request.Method == "POST" {		
		ct := request.Header.Get("content-type")
	if ct != "application/json" {
		response.WriteHeader(http.StatusUnsupportedMediaType)
		response.Write([]byte(fmt.Sprintf("expected content-type 'application/json', but got '%s'", ct)))
		return
	}
	var post models.Post
	var user models.User
	json.NewDecoder(request.Body).Decode(&post)

	if(post.Caption == "" || post.Image_URL == ""|| post.User_id == "" || post.Id !=primitive.NilObjectID){
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Please provide all the required fields"))
		return
	}

	// Checking for  user
	userId, err := primitive.ObjectIDFromHex(post.User_id)
	err= connection.UserCollection.FindOne(context.TODO(), bson.D{{"_id",userId}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// User does not exists
				response.WriteHeader(http.StatusNotFound)
				response.Write([]byte("User does not exists"))
				return
		}  else {
			panic(err)
		}
	} else {
		// User exists, create post
		post.Posted_timestamp = time.Now();
		result, _ := connection.PostCollection.InsertOne(context.TODO(), post)
			json.NewEncoder(response).Encode(result)
		return
	}

} else {
	response.WriteHeader(http.StatusMethodNotAllowed)
	response.Write([]byte("method not allowed"))
}
	
}