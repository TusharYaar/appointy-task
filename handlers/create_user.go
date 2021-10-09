package handlers

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tusharyaar/task/connection"
	"github.com/tusharyaar/task/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	// Returns if request is not post
	if request.Method != "POST" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{"message":"method not allowed"}`))
		return 
	}
	// Returns if application/json is not in header
	ct := request.Header.Get("content-type")
	if ct != "application/json" {
		response.WriteHeader(http.StatusUnsupportedMediaType)
		response.Write([]byte(fmt.Sprintf(`{"message":"expected content-type 'application/json', but got '%s'"}`, ct)))
		return
	}
	var user models.User
	var existing_user models.User
	json.NewDecoder(request.Body).Decode(&user)
	
	// check for empty fields
	if(user.Email == "" || user.Password == "" || user.Name == "" || user.Id !=primitive.NilObjectID){
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message":"Please provide all the required fields"}`))
		return
	}

	// Checking for exising user
	err:= connection.UserCollection.FindOne(context.TODO(), bson.D{primitive.E{ Key: "email",Value: user.Email}}).Decode(&existing_user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Hashing the password
			user.Password =	fmt.Sprintf("%x",sha256.Sum256([]byte(user.Password)))
			result, _ := connection.UserCollection.InsertOne(context.TODO(), user)
			response.WriteHeader(http.StatusCreated)
			json.NewEncoder(response).Encode(result)
		}  else {
			log.Fatal(err)
		}
	} else {
		// User already exists
		response.WriteHeader(http.StatusConflict)
		response.Write([]byte(`{"message":"User already exists"}`))
		return
	}

	
}