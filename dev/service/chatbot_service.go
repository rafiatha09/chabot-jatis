package service

import (
	err_helper "chatbot/dev/error_helper"
	"chatbot/dev/model/entity"
	"chatbot/dev/model/request"
	"chatbot/dev/model/response"
	"chatbot/dev/provider"
	"chatbot/dev/repository"
	"chatbot/dev/util"
	"errors"
	// "strconv"
	"strings"
	"github.com/gin-gonic/gin"
)



type ChabotService interface {
	Chatbot(ctx *gin.Context) (successResponse response.SuccessResponse, status int, err error)
}

type ChatbotServiceImpl struct {
	tokenRepository repository.TokenRepository
	userRepository repository.UserRepository
	chatbotRepository repository.ChatbotRepository
	chatHistoryRepository repository.ChatHistoryRepository
	log provider.ILogger
}

func NewChatbotServiceImpl(
	tokenRepository repository.TokenRepository,
	userRepository repository.UserRepository,
	chatbotRepository repository.ChatbotRepository,
	chatHistoryRepository repository.ChatHistoryRepository,
	log provider.ILogger) *ChatbotServiceImpl {
	return &ChatbotServiceImpl{
		tokenRepository: tokenRepository,
		userRepository: userRepository,
		chatbotRepository: chatbotRepository,
		chatHistoryRepository: chatHistoryRepository,
		log: log,
	}
}

func (p * ChatbotServiceImpl) Chatbot(ctx *gin.Context) (successResponse response.SuccessResponse, status int, err error) {
	bearerToken := strings.Split(ctx.GetHeader("Authorization"), " ")
	token_ := bearerToken[1]

	token , err := p.tokenRepository.GetTokenRepository(ctx, token_) 
	if (err != nil){
		if (errors.Is(err, err_helper.ErrRepositoryNotFound)){
			err = err_helper.ErrServiceUnauthorized
			return
		}
		return
	}

	if (!util.IsFutureTime(token.Expired)){
		err = err_helper.ErrServiceUnauthorized
		return
	}

	if ctx.GetHeader("Content-Type") != "application/json" {
		err = err_helper.ErrServiceMandatoryParameter
		return
	}

	var payload request.Request
	if err = ctx.ShouldBindJSON(&payload); err != nil {
		err = err_helper.ErrServiceRequestBody
		return
	}

	user,  err := p.userRepository.GetUserRepository(ctx, payload.UserID)
	if (err != nil){
		if (errors.Is(err, err_helper.ErrRepositoryNotFound)){
			err = err_helper.ErrServiceUserNotFound
			return
		}
		return
	}
	
	current_timestamp := util.GenerateCurrentTimestamp()
	if (payload.Timestamp == 0){
		payload.Timestamp = current_timestamp
	} 
		
	if(payload.Timestamp > current_timestamp){
		err = err_helper.ErrServiceTimestamp
		return
	}
	
	var chatHistory entity.ChatHistory
	chatbot , chatbotErr := p.chatbotRepository.GetChatbotRepository(ctx, user.LatestChatbotID, "", user.ClientID)

	// CASE 1: when first step and no chatbot data
	if (user.LatestChatbotID == "" || chatbotErr != nil){
		welcomeChatbot, getChatbotErr := p.chatbotRepository.GetChatbotRepository(ctx, "", util.Configuration.Chatbot.WelcomeTitle, "")
		if (getChatbotErr != nil){
			err = err_helper.ErrRepositoryNotFound
			return
		}

		selected, _  := findChatbotAnswer(welcomeChatbot.Options, "default")

		user.LatestChatbotID = selected.RelatedChatbotID
		err = p.userRepository.UpdateUserRepository(ctx, user)
		if (err != nil){
			if (errors.Is(err, err_helper.ErrRepositoryNotFound)){
				err = err_helper.ErrServiceUnauthorized
				return
			}
			return
		}

		chatHistory = prepareChatHistory(user.ID.Hex(), user.ClientID, payload.Text, selected.Response, payload.Timestamp)
		chatHistoryId, errr := p.chatHistoryRepository.CreateChatHistory(ctx, chatHistory)
		if (errr != nil) {
			return
		}

		successResponse = prepareSuccessResponse(chatHistoryId, selected.Response)
		return
	}
	
	// CASE 2: when the request containts in options select
	selected, isInOption := findChatbotAnswer(chatbot.Options, payload.Text)
	if (isInOption){
		user.LatestChatbotID = selected.RelatedChatbotID
		err = p.userRepository.UpdateUserRepository(ctx, user)
		if (err != nil){
			if (errors.Is(err, err_helper.ErrRepositoryNotFound)){
				err = err_helper.ErrServiceUnauthorized
				return
			}
			return
		}

		chatHistory = prepareChatHistory(user.ID.Hex(), user.ClientID, payload.Text, selected.Response, payload.Timestamp)
		chatHistoryId, errr := p.chatHistoryRepository.CreateChatHistory(ctx, chatHistory)
		if (errr != nil) {
			err = errr
			return
		}

		successResponse = prepareSuccessResponse(chatHistoryId, selected.Response)
		return
	}

	return
}


func findChatbotAnswer(options []entity.Option, desiredSelect string) (option entity.Option, isInOption bool){
	for i := 0; i < len(options); i++ {
		if (options[i].Select == desiredSelect) {
			return options[i], true
		}
	}
	return option, false
}

func prepareSuccessResponse(chatHistory string, text string) response.SuccessResponse {
	return response.SuccessResponse{
		Mid: chatHistory,
		Text: text,
		Timestamp: util.GenerateCurrentTimestamp(),
	}
}

func prepareChatHistory(userId string, clientId string, textReq string, textResp string, timeStamp int64) entity.ChatHistory{
	return entity.ChatHistory{
		UserID: userId,
		ClientID: clientId,
		TextReq: textReq,
		TextResp: textResp,
		Timestamp: timeStamp,
	}
}