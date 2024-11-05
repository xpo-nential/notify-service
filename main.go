package main

import (
	"github.com/xpo-nential/notify-service/configs"
	"github.com/xpo-nential/notify-service/src"
)

func main() {
	configs.Init()
	src.HandlerDiscord()
}
