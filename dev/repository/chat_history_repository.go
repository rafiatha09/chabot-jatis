package repository

import (
	"chatbot/dev/model/entity"
	"chatbot/dev/provider"
	"chatbot/dev/util"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ChatHistoryRepository interface {
	CreateChatHistory(ctx context.Context, chatHistroy entity.ChatHistory) (err error)
}

type ChatHistoryRepositoryImpl struct {
	db  *mongo.Database
	log provider.ILogger
}


func NewChatHistoryRepositoryImpl(client *mongo.Client, log provider.ILogger) *ChatHistoryRepositoryImpl {
	db := client.Database(util.Configuration.MongoDB.Database)
	return &ChatHistoryRepositoryImpl{db: db, log: log}
}

func (t *ChatHistoryRepositoryImpl) CreateChatHistory(ctx context.Context, chatHistroy entity.ChatHistory) (err error) {
	return
}
