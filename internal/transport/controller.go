package transport

import (
	"github.com/gin-gonic/gin"
	"messagio/internal/service"
)

type Controller struct {
	MessageController *MessageController
}

func NewController(messageService *service.MessageService) *Controller {
	return &Controller{
		MessageController: NewMessageController(messageService),
	}
}

func InitRouter(service *service.Service) error {
	router := gin.Default()

	router.Use(CORSMiddleware())

	controller := NewController(service.MessageService)
	controller.MessageController.registryProvider(router)

	err := router.Run(":8080")
	if err != nil {
		return err
	}
	return nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, UPDATE, DELETE")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	}
}
