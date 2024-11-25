package src

import (
	"github.com/gin-gonic/gin"
	"github.com/xpo-nential/notify-service/src/handler"
)

func Routes() {
	r := gin.Default()

	r.GET(`discord/get-massage-channel`, handler.GetMessageChannel)
	r.POST(`discord/send-massage-channel`, handler.SendDiscordMessageChannel)
	r.POST(`line/send-massage-channel`, handler.SendLineMessageChannel)

	r.Run()
}
