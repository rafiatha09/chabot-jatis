package entity

import "time"


type ChatHistory struct {
	ID        string  `bson:"_id" json:"_id"`
	ClientID  string             `bson:"client_id" json:"client_id"`
	UserID    string             `bson:"user_id" json:"user_id"`
	TextReq   string             `bson:"text_req" json:"text_req"`
	TextResp  string             `bson:"text_resp" json:"text_resp"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
}