package repository

import (
	"chatbot/dev/model/entity"
	"chatbot/dev/provider"
	"chatbot/dev/util"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type TokenRepository interface {
	GetTokenRepository(ctx context.Context, tokenAuth string) (token entity.Token, err error)
}

type TokenRepositoryImpl struct {
	db *mongo.Database
	log provider.ILogger
}

func NewTokenRepositoryImpl(client *mongo.Client, log provider.ILogger) *TokenRepositoryImpl {
	db := client.Database(util.Configuration.MongoDB.Database)
	return &TokenRepositoryImpl{db: db, log: log}
}

func (t *TokenRepositoryImpl) GetTokenRepository(ctx context.Context, tokenAuth string) (token entity.Token, err error) {
	return
}
