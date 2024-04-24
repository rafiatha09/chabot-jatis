package entity

import "go.mongodb.org/mongo-driver/bson/primitive"



type ChatHistory struct {
	ID        primitive.ObjectID    `bson:"_id" json:"_id"`
	ClientID  string             `bson:"client_id" json:"client_id"`
	UserID    string             `bson:"user_id" json:"user_id"`
	TextReq   string             `bson:"text_req" json:"text_req"`
	TextResp  string             `bson:"text_resp" json:"text_resp"`
	Timestamp int64          `bson:"timestamp" json:"timestamp"`
}