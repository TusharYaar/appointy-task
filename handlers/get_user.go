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


func GetUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")

	var user models.User

	parts := strings.Split(request.URL.Path, "/")
	if len(parts) != 3 {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	userId, err := primitive.ObjectIDFromHex(parts[2])
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Invalid user id"))
		return
	}
	err = connection.UserCollection.FindOne(context.TODO(), bson.D{{"_id",userId}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("User not found"))
		return
	}
	json.NewEncoder(response).Encode(user)
}