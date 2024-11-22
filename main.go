package main

import (
	"github.com/xpo-nential/notify-service/configs"
	discordbot "github.com/xpo-nential/notify-service/notify-service/bot-discord"
	linebot "github.com/xpo-nential/notify-service/notify-service/bot-line"
	"github.com/xpo-nential/notify-service/src"
)

func main() {
	configs.Init()
	linebot.Init()
	go src.Routes()
	discordbot.Init()
}
