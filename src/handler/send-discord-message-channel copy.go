package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	notifyservice "github.com/xpo-nential/notify-service/notify-service"
	botdiscord "github.com/xpo-nential/notify-service/notify-service/bot-discord"
)

type _SEND_DISCORD_MESSAGE struct {
	ChannelId string `json:"channel_id"`
	Message   string `json:"message"`
}

func SendDiscordMessageChannel(c *gin.Context) {

	var dto _SEND_DISCORD_MESSAGE

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})
	}

	// สร้าง sender สำหรับ Discord
	var sender notifyservice.INotifyService = botdiscord.DiscordSender{
		Target:  dto.ChannelId,
		Message: dto.Message,
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
