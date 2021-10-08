package models

// Define Post Type struct
// MongoDb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id primitive.ObjectID   	`json:"_id" bson:"_id,omitempty"`
	Caption string 				`json:"caption" bson:"caption,omitempty"`
	Image_URL string 			`json:"image_url" bson:"image_url,omitempty"`
	Posted_timestamp time.Time  `json:"posted_timestamp" bson:"posted_timestamp,omitempty"`
	User_id string 				`json:"user_id" bson:"user_id,omitempty"`
}