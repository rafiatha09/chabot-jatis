package repository

import (
	"chatbot/dev/model/entity"
	"chatbot/dev/provider"
	"chatbot/dev/util"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUserRepository(ctx context.Context, userId string) (user entity.User, err error)
	UpdateUserRepository(ctx context.Context, userUpdated entity.User, userId string)(err error) 
}

type UserRepositoryImpl struct {
	db *mongo.Database
	log provider.ILogger
}

func NewUserRepositoryImpl(client *mongo.Client, log provider.ILogger) *UserRepositoryImpl {
	db := client.Database(util.Configuration.MongoDB.Database)
	return &UserRepositoryImpl{db: db, log: log}
}

func (u *UserRepositoryImpl) GetUserRepository(ctx context.Context, userId string) (user entity.User, err error) {

	return
}


func (u *UserRepositoryImpl) UpdateUserRepository(ctx context.Context, userUpdated entity.User, userId string) (err error) {

	return
}

