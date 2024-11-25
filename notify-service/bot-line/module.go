package botline

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/xpo-nential/notify-service/config"
)

type IBotLine interface {
	BubbleMessage(target string, msg string, alt string) error
	TextMessage(target string, msg string) error
}

type botline struct {
	client *linebot.Client
}

var lineBot botline

func newLineBot() {
	cfg := config.NewConfig(config.LineBot).App()
	bot, err := linebot.New(cfg.GetSecret(), cfg.GetToken())
	if err != nil {
		log.Fatal(err.Error())
	}

	if bot != nil {
		log.Println(`line bot is connected!`)
	}
	lineBot.client = bot
}
