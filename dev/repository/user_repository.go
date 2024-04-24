package repository

import (
	err_helper "chatbot/dev/error_helper"
	"chatbot/dev/model/entity"
	"chatbot/dev/provider"
	"chatbot/dev/util"

	// "fmt"

	// "fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUserRepository(ctx *gin.Context, userId string) (user entity.User, err error)
	UpdateUserRepository(ctx *gin.Context, userUpdated entity.User)(err error) 
}

type UserRepositoryImpl struct {
	db *mongo.Database
	log provider.ILogger
}

func NewUserRepositoryImpl(client *mongo.Client, log provider.ILogger) *UserRepositoryImpl {
	db := client.Database(util.Configuration.MongoDB.Database)
	return &UserRepositoryImpl{db: db, log: log}
}

func (u *UserRepositoryImpl) GetUserRepository(ctx *gin.Context, userId string) (user entity.User, err error) {
    logger := u.log.WithFields(provider.MongoLog,logrus.Fields{"DATABASE_NAME": util.Configuration.MongoDB.Database,"COLLECTION_NAME": util.Configuration.MongoDB.Collection.User, "USER_ID":userId,},)
    logger.Infof("finding the product into monggo db")

	oid, _ := primitive.ObjectIDFromHex(userId)


    // Finding the document by ID
    filter := bson.M{"_id": oid}


    err = u.db.Collection(util.Configuration.MongoDB.Collection.User).FindOne(ctx, filter).Decode(&user)
	// fmt.Println(user.ID.Hex())
    if err != nil {
		logger.Errorf("finding into mongo db user failed %s", err)
        if err == mongo.ErrNoDocuments {
			err = err_helper.ErrRepositoryNotFound
            return
        }
		err = err_helper.ErrRepositoryMongo
        return 
    }
	logger.Infof("successfully finding the user data")
	return
}



func (u *UserRepositoryImpl) UpdateUserRepository(ctx *gin.Context, userUpdated entity.User) (err error) {
	logger := u.log.WithFields(provider.MongoLog,logrus.Fields{"DATABASE_NAME": util.Configuration.MongoDB.Database,"COLLECTION_NAME": util.Configuration.MongoDB.Collection.User, "USER_ID":userUpdated.ID.Hex(),},)
	logger.Infof("updating user in MongoDB")

	oid, err := primitive.ObjectIDFromHex(userUpdated.ID.Hex())
	if err != nil {
		logger.Errorf("invalid user ID format: %s", err)
		return 
	}

	mongoCtx := ctx.Request.Context()

	filter := bson.M{"_id": oid}
	
	update := bson.M{}
    if userUpdated.LatestChatbotID == "" {
		update["$set"] = bson.M{
			"point":  userUpdated.Point,
			"client_id": userUpdated.ClientID,
			"latest_chatbot_id" : "",
		}
    } else {
		update["$set"] = bson.M{
			"point":  userUpdated.Point,
			"client_id": userUpdated.ClientID,
			"latest_chatbot_id" : userUpdated.LatestChatbotID,
		}
	}
	
	result, err := u.db.Collection(util.Configuration.MongoDB.Collection.User).UpdateOne(mongoCtx, filter, update)
	if err != nil {
		logger.Errorf("error updating user in MongoDB: %s", err)
		return err_helper.ErrRepositoryMongo
	}

	if result.MatchedCount == 0 {
		logger.Infof("no document matches the provided ID")
		return err_helper.ErrRepositoryNotFound
	}

	logger.Infof("User updated successfully")
	return nil
}

