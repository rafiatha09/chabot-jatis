package repository

import (
	"chatbot/dev/model/entity"
	"chatbot/dev/provider"
	"chatbot/dev/util"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ChatbotRepository interface {
	GetChatbotRepository(ctx context.Context, chatbodId string) (chatbt entity.Chatbot, err error)
}

type ChatbotRepositoryImpl struct {
	db *mongo.Database
	log provider.ILogger
}

func NewChatbotRepositoryImpl(client *mongo.Client, log provider.ILogger) *ChatbotRepositoryImpl {
	db := client.Database(util.Configuration.MongoDB.Database)
	return &ChatbotRepositoryImpl{db: db, log: log}
}

func (t *ChatbotRepositoryImpl)GetChatbotRepository(ctx context.Context, chatbodId string) (chatbot entity.Chatbot, err error) {
	return
}