package src

import (
	"github.com/bwmarrin/discordgo"
	discordbot "github.com/xpo-nential/notify-service/servers/discord-bot"
	slashcmds "github.com/xpo-nential/notify-service/src/slash-cmds"
)

func HandlerDiscord() {
	slashcmds.HandlerCommand()
	slashcmds.RegisterCommand()

	discordbot.NewConnection().Start(func(s *discordgo.Session) {

	})

}
