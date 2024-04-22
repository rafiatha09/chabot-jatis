package entity

type Client struct {
	ID         string             `bson:"_id" json:"_id"`
	ClientName string             `bson:"client_name" json:"client_name"`
}