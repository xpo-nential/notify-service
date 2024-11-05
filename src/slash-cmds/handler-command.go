package slashcmds

import (
	"github.com/bwmarrin/discordgo"
	"github.com/xpo-nential/notify-service/servers/discord-bot/handlers"
)

func HandlerCommand() {
	handlers.NewSlashHandler.AddHandlers(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"basic-command": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hey there! Congratulations, you just executed your first slash command",
				},
			})
		},
	})
}
