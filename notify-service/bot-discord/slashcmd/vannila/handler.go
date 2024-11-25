package vannila

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func (v *Vannila) GetHandler() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return v.handler
}

var vannilaHdl = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"basic-vannila": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Hey there! Congratulations, you just executed your first slash command",
			},
		})
		if err != nil {
			s.ChannelMessageSend(i.ChannelID, "An error occurred while responding to the command.")
		}
	},
	"today": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: time.Now().Format(time.DateOnly),
			},
		})
		if err != nil {
			s.ChannelMessageSend(i.ChannelID, "An error occurred while responding to the command.")
		}
	},
}
