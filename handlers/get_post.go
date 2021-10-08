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
	"go.mongodb.org/mongo-driver/mongo"
)


func GetPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")

	var post models.Post

	parts := strings.Split(request.URL.Path, "/")
	if len(parts) != 3 {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	postId, err := primitive.ObjectIDFromHex(parts[2])
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Invalid post id"))
		return
	}
	err = connection.PostCollection.FindOne(context.TODO(), bson.D{{"_id",postId}}).Decode(&post)
	if err == mongo.ErrNoDocuments {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("Post not found"))
		return
	}
	json.NewEncoder(response).Encode(post)
}