package main

import (
	"chatbot/dev/http/api"
	"chatbot/dev/provider"
	"chatbot/dev/repository"
	"chatbot/dev/service"
	"chatbot/dev/util"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	if err := util.LoadConfig("."); err != nil {
		log.Fatal(err)
	}

	provider.InitLogDir()
}

func main() {
    // router := gin.Default()

    // // Define a route that handles GET requests on "/"
    // router.GET("/", func(c *gin.Context) {
    //     c.JSON(http.StatusOK, gin.H{
    //         "message": "Hello world!",
    //     })
    // })

    // // Start the server on port 8080
    // router.Run(":8080")

	logger := provider.NewLogger()
    mongoClient, err := provider.NewMongoDBClient()

    if err != nil {
		log.Fatal(err)
	}

	logger.Infof(provider.AppLog, "Successfully connected to MongoDB.")
    go func(c *mongo.Client, logger provider.ILogger) {
		var tokenRepository repository.TokenRepository = repository.NewTokenRepositoryImpl(c, logger)
		var userRepository repository.UserRepository = repository.NewUserRepositoryImpl(c, logger)
		var chatbotRepository repository.ChatbotRepository = repository.NewChatbotRepositoryImpl(c, logger)
		var chatHistoryRepository repository.ChatHistoryRepository = repository.NewChatHistoryRepositoryImpl(c, logger)

		var chatbotService service.ChabotService = service.NewChatbotServiceImpl(tokenRepository,userRepository, chatbotRepository, chatHistoryRepository,logger)

		app := api.NewApp(logger, chatbotService)
		addr := fmt.Sprintf(":%v", util.Configuration.Server.Port)
		server, err := app.CreateServer(addr)
		if err != nil {
			log.Fatal(err)
		}

		logger.Infof(provider.AppLog, "Server running at: %s", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Errorf(provider.AppLog, "Server error: %v", err)
		}

	}(mongoClient, logger)

    shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGTERM)

	sig := <-shutdownCh
	logger.Infof(provider.AppLog, "Receiving signal: %s", sig)

	func(c *mongo.Client) {
		if err := c.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}

		logger.Infof(provider.AppLog, "Successfully disconnected from MongoDB.")

	}(mongoClient)

}

