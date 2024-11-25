package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMessageChannel(c *gin.Context) {
	query := c.DefaultQuery(`query`, `GuildText`)
	// idType := botdiscord.TypeChannal[query]

	// discordService := botdiscord.DiscordService
	// resp, err := discordService.GetAllChannels(discordgo.ChannelType(idType))
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		`message`: err.Error(),
	// 	})
	// }

	c.JSON(http.StatusOK, gin.H{
		`resp`: query,
	})
}
