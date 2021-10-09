package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/tusharyaar/task/connection"
	"github.com/tusharyaar/task/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	options := options.Find()
	page, limit := Pagination(request, options)

	// add pagination value to header
	response.Header().Add("pagination-page", strconv.Itoa(int(page)))
	response.Header().Add("pagination-limit", strconv.Itoa(int(limit)))
	// fmt.Printf("Page: %d, Limit: %d", page, limit)

	cur, err := connection.PostCollection.Find(context.TODO(), bson.D{primitive.E{Key:"user_id", Value:userId}},options)

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


// Function for pagination
// Expects page and limit as query parameters
// Returns page and limit as int64
func Pagination(request *http.Request, FindOptions *options.FindOptions) (int64, int64) {
    var page, limit int64 
	if request.URL.Query().Get("limit") != "" && request.URL.Query().Get("page") != "" {
		limit, _ = strconv.ParseInt(request.URL.Query().Get("limit"), 10, 32)
		page, _ = strconv.ParseInt(request.URL.Query().Get("page"), 10, 32)
        if page == 1 {
            FindOptions.SetSkip(0)
            FindOptions.SetLimit(limit)
            return page, limit
        }

        FindOptions.SetSkip((page - 1) * limit)
        FindOptions.SetLimit(limit)
        return page, limit

    }
    FindOptions.SetSkip(0)
    FindOptions.SetLimit(0)
    return 0, 0
}