package src

import (
	"github.com/gin-gonic/gin"
	"github.com/xpo-nential/notify-service/src/handler"
)

func Routes() {
	r := gin.Default()

	r.GET(`discord/get-massage-channel`, handler.XGetMessageChannel)
	r.POST(`discord/send-massage-channel`, handler.XSendDiscordMessageChannel)
	r.POST(`line/send-massage-channel`, handler.XSendLineMessageChannel)

	r.Run()
}
