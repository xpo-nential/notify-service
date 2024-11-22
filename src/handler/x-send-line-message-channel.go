package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	notifyservice "github.com/xpo-nential/notify-service/notify-service"
	botline "github.com/xpo-nential/notify-service/notify-service/bot-line"
)

type _SEND_LINE_MESSAGE struct {
	ChannelId string                 `json:"channel_id"`
	Message   interface{}            `json:"message"`
	Alt       *string                `json:"alt"`
	Template  map[string]interface{} `json:"template"`
}

func XSendLineMessageChannel(c *gin.Context) {

	var dto _SEND_LINE_MESSAGE

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})
	}

	jsonData, err := json.MarshalIndent(dto.Template, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	template := string(jsonData)
	// สร้าง sender สำหรับ Discord
	var sender notifyservice.INotifyService = botline.LineSender{
		Target:   dto.ChannelId,
		Message:  dto.Message,
		Alt:      dto.Alt,
		Template: &(template),
	}

	if err := sender.SendMessage(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to send message",
			"details": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		`success`: true,
	})
}
