package main

import (
	"github.com/xpo-nential/notify-service/configs"
	discordbot "github.com/xpo-nential/notify-service/servers/discord-bot"
	linebot "github.com/xpo-nential/notify-service/servers/line-bot"
	"github.com/xpo-nential/notify-service/src"
)

func main() {
	configs.Init()
	linebot.Init()
	go src.Routes()
	discordbot.Init()
}
