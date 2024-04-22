package api

import (
	"chatbot/dev/model/response"
	"chatbot/dev/provider"
	"chatbot/dev/util"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func authorization(logger provider.ILogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		bearerToken := strings.Split(ctx.GetHeader("Authorization"), " ")
		invalidTokenMsg := "Unauthorized"

		unauthorizedResp := response.ErrorResponse {
			Timestamp: util.GenerateCurrentTimestamp(),
			Error: invalidTokenMsg,
		}


		if len(bearerToken) != 2 {
			logger.WithFields(provider.AppLog, logrus.Fields{"REQUEST_ID": "reqID"}).Error(invalidTokenMsg)
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				unauthorizedResp,
			)
			return
		}

		if bearerToken[0] != "Bearer" {
			logger.WithFields(provider.AppLog, logrus.Fields{"REQUEST_ID": "reqID"}).Error(invalidTokenMsg)
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				unauthorizedResp,
			)
			return
		}

		token := bearerToken[1]

		ctx.Set("token", token)
		ctx.Next()
	}
}

func loggingMiddleware(logger provider.ILogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Starting time
		startTime := time.Now()

		// Processing request
		ctx.Next()

		// End Time
		endTime := time.Now()

		// execution time
		latencyTime := endTime.Sub(startTime)

		// Request method
		reqMethod := ctx.Request.Method

		// Request route
		reqUri := ctx.Request.RequestURI

		// status code
		statusCode := ctx.Writer.Status()

		// Request IP
		clientIP := ctx.ClientIP()

		logger.WithFields(
			provider.AppLog,
			logrus.Fields{
				"METHOD":     reqMethod,
				"URI":        reqUri,
				"STATUS":     statusCode,
				"LATENCY":    latencyTime,
				"CLIENT_IP":  clientIP,
			},
		).Info("HTTP REQUEST")

		ctx.Next()
	}
}
