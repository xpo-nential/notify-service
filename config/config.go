package config

import (
	"os"
)

type BotConfig string

var (
	DiscordBot BotConfig = `discord`
	LineBot    BotConfig = `line`
)

func NewConfig(b BotConfig) *config {
	switch b {
	case DiscordBot:
		return &config{
			app: &app{
				token: os.Getenv(`DISCORD_TOKEN`),
			},
		}
	case LineBot:
		return &config{
			app: &app{
				token:  os.Getenv(`LINE_CHANNEL_TOKEN`),
				secret: os.Getenv(`LINE_CHANNEL_SECRET`),
			},
		}
	default:
		return nil
	}
}
