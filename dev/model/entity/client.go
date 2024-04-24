package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Client struct {
	ID         primitive.ObjectID             `bson:"_id" json:"_id"`
	ClientName string             `bson:"client_name" json:"client_name"`
}