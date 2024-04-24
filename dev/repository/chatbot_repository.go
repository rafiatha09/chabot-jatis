package repository

import (
	err_helper "chatbot/dev/error_helper"
	"chatbot/dev/model/entity"
	"chatbot/dev/provider"
	"chatbot/dev/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatbotRepository interface {
	GetChatbotRepository(ctx *gin.Context, chatbodId string, title string, clientId string) (chatbt entity.Chatbot, err error)
}

type ChatbotRepositoryImpl struct {
	db *mongo.Database
	log provider.ILogger
}

func NewChatbotRepositoryImpl(client *mongo.Client, log provider.ILogger) *ChatbotRepositoryImpl {
	db := client.Database(util.Configuration.MongoDB.Database)
	return &ChatbotRepositoryImpl{db: db, log: log}
}

func (t *ChatbotRepositoryImpl)GetChatbotRepository(ctx *gin.Context, chatbodId string, title string, clientId string) (chatbot entity.Chatbot, err error) {
	logger := t.log.WithFields(provider.MongoLog,logrus.Fields{"DATABASE_NAME": util.Configuration.MongoDB.Database,"COLLECTION_NAME": util.Configuration.MongoDB.Collection.User, "CHATBOT_ID":chatbodId, "TITLE": title, "CLIENT_ID" : clientId},)
    logger.Infof("finding the chatbot into monggo db")

	var filter bson.M
	if (title != ""){
		filter = bson.M{"title": title}
	}else {
		oid, _ := primitive.ObjectIDFromHex(chatbodId)
		filter = bson.M{"_id": oid, "client_id": clientId}
	}


    err = t.db.Collection(util.Configuration.MongoDB.Collection.Chatbot).FindOne(ctx, filter).Decode(&chatbot)
    if err != nil {
		logger.Errorf("finding into mongo db chatbot failed %s", err)
        if err == mongo.ErrNoDocuments {
			err = err_helper.ErrRepositoryNotFound
            return
        }
		err = err_helper.ErrRepositoryMongo
        return 
    }
	logger.Infof("successfully finding the chatbot data")
	return
}