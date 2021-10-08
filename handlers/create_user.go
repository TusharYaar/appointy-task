package handlers

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tusharyaar/task/connection"
	"github.com/tusharyaar/task/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {		
	var user models.User
	var existing_user models.User
	response.Header().Add("Content-Type", "application/json")
	json.NewDecoder(request.Body).Decode(&user)
	// Checking for exising user
	err:= connection.UserCollection.FindOne(context.TODO(), bson.D{{"email",user.Email}}).Decode(&existing_user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Hashing the password
			user.Password =	fmt.Sprintf("%x",sha256.Sum256([]byte(user.Password)))
			result, _ := connection.UserCollection.InsertOne(context.TODO(), user)
			json.NewEncoder(response).Encode(result)
		}  else {
			panic(err)
		}
	} else {
		// User already exists
		response.WriteHeader(http.StatusConflict)
		response.Write([]byte("User already exists"))
		return
	}

}
	
}