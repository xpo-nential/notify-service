package main

import (
	"github.com/xpo-nential/notify-service/config"
	discordbot "github.com/xpo-nential/notify-service/notify-service/bot-discord"
	linebot "github.com/xpo-nential/notify-service/notify-service/bot-line"
	"github.com/xpo-nential/notify-service/src"
)

func main() {
	config.Init()
	linebot.Init()
	go src.Routes()
	discordbot.Init()
}
