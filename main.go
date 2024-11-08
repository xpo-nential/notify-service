package main

import (
	"github.com/xpo-nential/notify-service/configs"
	discordbot "github.com/xpo-nential/notify-service/servers/discord-bot"
)

func main() {
	configs.Init()
	discordbot.Init()
}
