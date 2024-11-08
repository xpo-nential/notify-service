package functions

import (
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	discordbot "github.com/xpo-nential/notify-service/servers/discord-bot"
)

func XGetMessageChannel(c *gin.Context) {
	query := c.DefaultQuery(`query`, `GuildText`)
	idType := discordbot.TypeChannal[query]

	discordService := discordbot.DiscordService
	resp, err := discordService.GetAllChannels(discordgo.ChannelType(idType))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			`message`: err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		`resp`: resp,
	})
}
