package entity

import "time"

type Token struct {
	ID       string `bson:"_id" json:"_id"`
	ClientID string             `bson:"client_id" json:"client_id"`
	Token    string             `bson:"token" json:"token"`
	Expired  time.Time          `bson:"expired" json:"expired"`
}