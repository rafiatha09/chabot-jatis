package api

import (
	// err_helper "chatbot/dev/error_helper"
	err_helper "chatbot/dev/error_helper"
	"chatbot/dev/provider"
	"chatbot/dev/service"
	"chatbot/dev/util"
	"fmt"
	"net/http"

	// "strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	// "github.com/sirupsen/logrus"
)

type App struct {
	// productService service.ProductService
	chatbotService service.ChabotService
	log provider.ILogger
}

func NewApp(log provider.ILogger, chatbotService service.ChabotService) *App{
	return &App{log: log, chatbotService: chatbotService}
}

func (a *App) CreateServer(address string) (*http.Server, error) {
	gin.SetMode(util.Configuration.Server.Mode)

	r := gin.Default()
	r.Use(gin.Recovery())

	r.Use(loggingMiddleware(a.log))
	r.GET("/ping", a.checkConnectivity)
	r.Use(authorization(a.log))
	r.POST("/chatbot", a.chatbotApi)

	server := &http.Server{
		Addr:    address,
		Handler: r,
	}

	return server, nil
}


func (a *App) checkConnectivity(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (a *App) chatbotApi(ctx *gin.Context){
	var reqID string
	val, ok := ctx.Get("req-id")
	if ok {
		reqID = val.(string)
	}

	fmt.Println("chatbot nih bos")
	response, _, err := a.chatbotService.Chatbot(ctx)
	fmt.Println("senggol dong")

	if (err != nil){
		title, code, res :=  err_helper.MapError(err)
		log := a.log.WithFields(provider.AppLog, logrus.Fields{"REQUEST_ID": reqID})
		log.Errorf(title)
		ctx.JSON(code, res)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusAccepted, response)
}

// func (a *App) getProduct(ctx *gin.Context){
// 	if ctx.Param("productId") == "" {
// 		code := http.StatusBadRequest
// 		errorResp := response.ErrorResponse{Error: response.Error{Code: code, Title: "InvalidPath", Details: "Invalid path"},}
// 		a.log.WithFields(provider.AppLog, logrus.Fields{"REQUEST_ID": "id-belajar"}).Errorf("parse path error")
// 		ctx.JSON(code, errorResp)
// 		ctx.Abort()
// 		return
// 	}

// 	product, status, err := a.productService.GetProductService(ctx,  ctx.Param("productId"))
// 	if err != nil {
// 		errorResp := response.ErrorResponse{Error: response.Error{Code: status, Title: "FailedGetProduct", Details: err.Error()},}
// 		a.log.WithFields(provider.AppLog, logrus.Fields{"REQUEST_ID": "id-belajar"}).Errorf("failed get desired product: %s", err.Error())
// 		ctx.JSON(status, errorResp)
// 		return
// 	}

// 	resp := response.SuccessResponse{Message: "successfully get desired product",Data: product,}
// 	a.log.WithFields(provider.AppLog, logrus.Fields{"REQUEST_ID": "id-belajar"}).Infof("send message success, response: %#v", resp.Message)
// 	ctx.JSON(http.StatusOK, resp)
// }

// func (a *App) getAllProducts(ctx *gin.Context){
// 	products, status, err := a.productService.GetAllProductsService(ctx)
// 	if err != nil {
// 		errorResp := response.ErrorResponse{
// 			Error: response.Error{Code: status, Title: "FailedGetProducts", Details: err.Error()},
// 		}
// 		ctx.JSON(status, errorResp)
// 		return
// 	}
// 	resp := response.SuccessResponse{Message: "successfully get desired products",Data: products,}
// 	a.log.WithFields(provider.AppLog, logrus.Fields{"REQUEST_ID": "id-belajar"}).Infof("send message success, response: %#v", resp)
// 	ctx.JSON(http.StatusOK, resp)
// }

// func (a *App) deleteProduct(ctx *gin.Context){
// 	if ctx.Param("productId") == "" {
// 		code := http.StatusBadRequest
// 		errorResp := response.ErrorResponse{
// 			Error: response.Error{Code: code, Title: "InvalidPath", Details: "Invalid path"},
// 		}
// 		a.log.WithFields(provider.AppLog, logrus.Fields{"REQUEST_ID": "id-belajar"}).Errorf("parse path error")
// 		ctx.JSON(code, errorResp)
// 		ctx.Abort()
// 		return
// 	}

// 	status, err := a.productService.DeleteProductService(ctx, ctx.Param("productId"))
// 	if err != nil {
// 		errorResp := response.ErrorResponse{
// 			Error: response.Error{Code: status, Title: "FailedDeleteProduct", Details: err.Error() },
// 		}
// 		a.log.WithFields(provider.AppLog, logrus.Fields{"REQUEST_ID": "id-belajar"}).Errorf("send message error, reason: %s", err)

// 		ctx.JSON(status, errorResp)
// 		ctx.Abort()
// 		return
// 	}
// 	resp := response.SuccessResponse{Message: "successfully delete desired products",Data: nil,}
// 	a.log.WithFields(provider.AppLog, logrus.Fields{"REQUEST_ID": "id-belajar"}).Infof("delete product success, response: %#v", resp)
// 	ctx.JSON(http.StatusOK, resp)
// }
// func (a *App) updateProduct(ctx *gin.Context){
// 	var payload entity.Product
// 	if err := ctx.ShouldBindJSON(&payload); err != nil {
// 		code := http.StatusBadRequest
// 		errorResp := response.ErrorResponse{Error: response.Error{Code: code, Title: "InvalidRequestBody", Details: err.Error()},}
// 		a.log.WithFields(provider.AppLog, logrus.Fields{"REQUEST_ID": "id-belajar"}).Errorf("parse request body error: %s", err)
// 		ctx.JSON(code, errorResp)
// 		return
// 	}
// 	a.log.WithFields(provider.AppLog, logrus.Fields{"REQUEST_ID": "id-belajar"}).Infof("get request body success, request_body: %#v", payload)

// 	result, status,  err := a.productService.UpdateProductService(ctx, payload)
// 	if (err != nil){
// 		errorResp := response.ErrorResponse{Error: response.Error{Code: status, Title: "FailedUpdateProduct", Details: err.Error()},}
// 		a.log.WithFields(provider.AppLog, logrus.Fields{"REQUEST_ID": "id-belajar"}).Errorf("failed update product: %s", err)
// 		ctx.JSON(status, errorResp)
// 		return
// 	}

// 	resp := response.SuccessResponse{Message: "successfully update product", Data: result}
// 	a.log.WithFields(provider.AppLog, logrus.Fields{"REQUEST_ID": "id-belajar"}).Infof("send message success, response: %#v", resp)
// 	ctx.JSON(http.StatusOK, resp)
// }



// func (a *App) createProduct(ctx *gin.Context){
// 	var payload request.Product
// 	if err := ctx.ShouldBindJSON(&payload); err != nil {
// 		code := http.StatusBadRequest
// 		errorResp := response.ErrorResponse{Error: response.Error{Code: code, Title: "InvalidRequestBody", Details: err.Error()},}
// 		a.log.WithFields(provider.AppLog, logrus.Fields{"REQUEST_ID": "id-belajar"}).Errorf("parse request body error: %s", err)
// 		ctx.JSON(code, errorResp)
// 		return
// 	}

// 	a.log.WithFields(provider.AppLog, logrus.Fields{"REQUEST_ID": "id-belajar"}).Infof("get request body success, request_body: %#v", payload)

// 	result ,status, err := a.productService.CreateProductService(ctx, payload)
// 	if err != nil {
// 		errorResp := response.ErrorResponse{
// 			Error: response.Error{Code: status, Title: "FailedCreateProduct", Details: err.Error()},
// 		}
// 		a.log.WithFields(provider.AppLog, logrus.Fields{"REQUEST_ID": "id-belajar"}).Errorf("send message error, reason: %s", err)
// 		ctx.JSON(status, errorResp)
// 		return
// 	}


// 	resp := response.SuccessResponse{Message: "successfully created product", Data: result}
// 	a.log.WithFields(provider.AppLog, logrus.Fields{"REQUEST_ID": "id-belajar"}).Infof("send message success, response: %#v", resp)
// 	ctx.JSON(http.StatusOK, resp)

// }