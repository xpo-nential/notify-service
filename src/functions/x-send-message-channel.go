package functions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	discordbot "github.com/xpo-nential/notify-service/servers/discord-bot"
)

type _SEND_MESSAGE struct {
	ChannelId string `json:"channel_id"`
	Message   string `json:"message"`
}

func XSendMessageChannel(c *gin.Context) {

	var dto _SEND_MESSAGE

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	discordSevice := discordbot.DiscordService

	resp, err := discordSevice.ChannelMessageSend(dto.ChannelId, dto.Message)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, gin.H{
		`resp`: resp,
	})
}
