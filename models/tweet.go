package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tweet struct {
	UserID  string    `bson:"userID" json:"userID,omitempty"`
	Message string    `bson:"message" json:"message,omitempty"`
	Date    time.Time `bson:"date" json:"date,omitempty"`
}

type TweetRetrieved struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"userID" json:"userID,omitempty"`
	Message string             `bson:"message" json:"message,omitempty"`
	Date    time.Time          `bson:"date" json:"date,omitempty"`
}

type FollowerTweetRetrieved struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"userID" json:"userID,omitempty"`
	UserRelationID string             `bson:"userRelationID" json:"userRelationID"`
	Tweet          struct {
		Message string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
