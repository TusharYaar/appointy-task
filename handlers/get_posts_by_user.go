package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/tusharyaar/task/connection"
	"github.com/tusharyaar/task/models"
	"go.mongodb.org/mongo-driver/bson"
)


func GetPostsByUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")

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

	cur, err := connection.PostCollection.Find(context.TODO(), bson.D{{"user_id",userId}})

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
			response.Write([]byte("Error while decoding post"))
			return 
		}

		allPosts = append(allPosts, post)
	}
		
	
	if len(allPosts) == 0 {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("Post not found"))
		return
	}
	json.NewEncoder(response).Encode(allPosts)
}