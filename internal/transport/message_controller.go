package transport

import (
	"github.com/gin-gonic/gin"
	"messagio/internal/service"
	"net/http"
)

type MessageController struct {
	MessageService *service.MessageService
}

func NewMessageController(ms *service.MessageService) *MessageController {
	return &MessageController{
		MessageService: ms,
	}
}

func (dr *MessageController) registryProvider(router *gin.Engine) {
	router.POST("/api/v1/message/add", dr.addMessage)
	router.GET("/api/v1/message/stats", dr.getMessageStats)

}

func (mc *MessageController) addMessage(c *gin.Context) {
	var requestBody struct {
		Text string `json:"text"`
	}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	id, err := mc.MessageService.AddMessage(requestBody.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (mc *MessageController) getMessageStats(c *gin.Context) {
	stats, err := mc.MessageService.GetMessageStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get stats"})
		return
	}

	c.JSON(http.StatusOK, stats)
}
