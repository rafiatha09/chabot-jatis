package service

import (
	"chatbot/dev/model/response"
	"chatbot/dev/provider"
	"chatbot/dev/repository"
	"context"
)



type ChabotService interface {
	Chatbot(ctx context.Context) (successResponse response.SuccessResponse, status int, err error)
}

type ChatbotServiceImpl struct {
	tokenRepository repository.TokenRepository
	userRepository repository.UserRepository
	chatbotRepository repository.ChatbotRepository
	chatHistory repository.ChatHistoryRepository
	log provider.ILogger
}

func NewChatbotServiceImpl(
	tokenRepository repository.TokenRepository,
	userRepository repository.UserRepository,
	chatbotRepository repository.ChatbotRepository,
	chatHistory repository.ChatHistoryRepository,
	log provider.ILogger) *ChatbotServiceImpl {
	return &ChatbotServiceImpl{
		tokenRepository: tokenRepository,
		userRepository: userRepository,
		chatbotRepository: chatbotRepository,
		chatHistory: chatHistory,
		log: log,
	}
}

func (p * ChatbotServiceImpl) Chatbot(ctx context.Context) (successResponse response.SuccessResponse, status int, err error) {
	return
}