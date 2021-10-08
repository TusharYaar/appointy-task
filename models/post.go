package models

// Define Post Type struct
// MongoDb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id primitive.ObjectID `json:"id"`
	Caption string `json:"caption"`
	Image_URL string `json:"image_url"`
	Posted_timestamp string `json:"posted_timestamp"`
	User_id string `json:"user_id"`
}