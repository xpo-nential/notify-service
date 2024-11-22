package botdiscord

import (
	"github.com/bwmarrin/discordgo"
	slashcmds "github.com/xpo-nential/notify-service/notify-service/bot-discord/slash-cmds"
)

func Init() {
	slCmd := slashcmds.SlashCommand

	slCmd.HandleCommands()
	slCmd.RegisterCommand()

	newConnection().Start(func(s *discordgo.Session) {

	})

}
