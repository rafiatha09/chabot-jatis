package repository

import (
	"chatbot/dev/model/entity"
	"chatbot/dev/provider"
	"chatbot/dev/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatHistoryRepository interface {
	CreateChatHistory(ctx *gin.Context, chatHistroy entity.ChatHistory) (chatHistoryId string, err error)
}

type ChatHistoryRepositoryImpl struct {
	db  *mongo.Database
	log provider.ILogger
}


func NewChatHistoryRepositoryImpl(client *mongo.Client, log provider.ILogger) *ChatHistoryRepositoryImpl {
	db := client.Database(util.Configuration.MongoDB.Database)
	return &ChatHistoryRepositoryImpl{db: db, log: log}
}

func (t *ChatHistoryRepositoryImpl) CreateChatHistory(ctx *gin.Context, chatHistory entity.ChatHistory) (chatHistoryId string, err error) {
	logger := t.log.WithFields(provider.MongoLog,logrus.Fields{"DATABASE_NAME": util.Configuration.MongoDB.Database,"COLLECTION_NAME": util.Configuration.MongoDB.Collection.ChatHistory},)
	logger.Infof("inserting into MongoDB database")


	mongoCtx := ctx.Request.Context()
	chatHistory.ID = primitive.NewObjectID()

	result, err := t.db.Collection(util.Configuration.MongoDB.Collection.ChatHistory).InsertOne(mongoCtx, chatHistory)
	if err != nil {
		logger.Errorf("creating chat history in MongoDB failed: %s", err)
		return 
	}


	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		logger.Errorf("error asserting InsertedID to ObjectID")
		return 
	}
	chatHistoryId = oid.Hex() // Convert ObjectID to string

	logger.Infof("successfully created chat history with ID: %s", chatHistoryId)
	return 
}
