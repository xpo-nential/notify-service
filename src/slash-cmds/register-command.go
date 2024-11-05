package slashcmds

import (
	"github.com/bwmarrin/discordgo"
	"github.com/xpo-nential/notify-service/servers/discord-bot/commands"
)

func RegisterCommand() {
	commands.NewSlashCommand.AddCommand([]*discordgo.ApplicationCommand{
		{
			Name: "basic-command",
			// All commands and options must have a description
			// Commands/options without description will fail the registration
			// of the command.
			Description: "Basic command",
		},
	})
}
