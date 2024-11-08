package src

import (
	"github.com/gin-gonic/gin"
	"github.com/xpo-nential/notify-service/src/functions"
)

func Routes() {
	r := gin.Default()

	r.GET(`/get-massage-channel`, functions.XGetMessageChannel)
	r.POST(`/send-massage-channel`, functions.XSendMessageChannel)
	r.Run()
}
