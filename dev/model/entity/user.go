package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID  			`bson:"_id,omitempty" json:"id,omitempty"`
	ClientID       string  `bson:"client_id,omitempty" json:"client_id,omitempty"`
	Point          int     `bson:"point,omitempty" json:"point,omitempty"`
	LatestChatbotID string `bson:"latest_chatbot_id,omitempty" json:"latest_chatbot_id,omitempty"`
}