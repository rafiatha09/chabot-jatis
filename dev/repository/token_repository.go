package repository

import (
	err_helper "chatbot/dev/error_helper"
	"chatbot/dev/model/entity"
	"chatbot/dev/provider"
	"chatbot/dev/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenRepository interface {
	GetTokenRepository(ctx *gin.Context, tokenAuth string) (token entity.Token, err error)
}

type TokenRepositoryImpl struct {
	db *mongo.Database
	log provider.ILogger
}

func NewTokenRepositoryImpl(client *mongo.Client, log provider.ILogger) *TokenRepositoryImpl {
	db := client.Database(util.Configuration.MongoDB.Database)
	return &TokenRepositoryImpl{db: db, log: log}
}

func (t *TokenRepositoryImpl) GetTokenRepository(ctx *gin.Context, tokenAuth string) (token entity.Token, err error) {
	filter := bson.M{"token": tokenAuth}

    logger := t.log.WithFields(provider.MongoLog,logrus.Fields{"DATABASE_NAME": util.Configuration.MongoDB.Database,"COLLECTION_NAME": util.Configuration.MongoDB.Collection.Token, "TOKEN": tokenAuth,},)
    logger.Infof("finding the token into monggo db")

    err = t.db.Collection(util.Configuration.MongoDB.Collection.Token).FindOne(ctx, filter).Decode(&token)
    if err != nil {
		logger.Errorf("finding into mongo db token failed %s", err)
        if err == mongo.ErrNoDocuments {
			err = err_helper.ErrRepositoryNotFound
            return
        }
		err = err_helper.ErrRepositoryMongo
        return 
    }
	logger.Infof("successfully finding the token data")
	return
}
