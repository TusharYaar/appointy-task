package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/tusharyaar/task/connection"
	"github.com/tusharyaar/task/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetPostsByUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")

		
	if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{"message":"method not allowed"}`))
		return 
	}
	

	var allPosts[] models.Post

	parts := strings.Split(request.URL.Path, "/")
	if len(parts) != 4 {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "Invalid request"}`))
		return
	}
	if (parts[3] == "") {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "Invalid request"}`))
		return
	}

	userId := parts[3]

	cur, err := connection.PostCollection.Find(context.TODO(), bson.D{primitive.E{Key:"user_id", Value:userId}})

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var post models.Post
		err := cur.Decode(&post)

		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{"message":"Error while decoding post"}`))
			return 
		}

		allPosts = append(allPosts, post)
	}
		
	
	if len(allPosts) == 0 {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(`{"message":"Post not found"}`))
		return
	}
	json.NewEncoder(response).Encode(allPosts)
}