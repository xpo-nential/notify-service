package linebot

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/xpo-nential/notify-service/configs"
)

type IBotLine interface {
	Message(target string, msg string, alt string) error
}

type botline struct {
	client *linebot.Client
}

var LineBot botline

func newLineBot() {
	cfg := configs.NewConfig(configs.LineBot).App()
	bot, err := linebot.New(cfg.GetSecret(), cfg.GetToken())
	if err != nil {
		log.Fatal(err.Error())
	}

	if bot != nil {
		log.Println(`line bot is connected!`)
	}
	LineBot.client = bot
}
