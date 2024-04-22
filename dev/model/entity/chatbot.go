package entity

type Chatbot struct {
	ID              string 			   `bson:"_id" json:"_id"`
	ClientID        string             `bson:"client_id" json:"client_id"`
	ParentChatbotID string             `bson:"parent_chatbot_id" json:"parent_chabot_id,omitempty"`
	Title           string             `bson:"title" json:"title"`
	Options         []Option           `bson:"options" json:"options"`
}

type Option struct {
	Select             string `bson:"select" json:"select"`
	Response           string `bson:"response" json:"response"`
	RelatedChatbotID   string `bson:"related_chatbot_id" json:"related_chatbot_id"`
}