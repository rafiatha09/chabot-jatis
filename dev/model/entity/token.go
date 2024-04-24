package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	ID       primitive.ObjectID   `bson:"_id" json:"_id"`
	ClientID string             `bson:"client_id" json:"client_id"`
	Token    string             `bson:"token" json:"token"`
	Expired  time.Time          `bson:"expired" json:"expired"`
}