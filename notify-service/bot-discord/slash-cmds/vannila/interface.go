package vannila

import "github.com/bwmarrin/discordgo"

type IVannila interface {
	GetCommand() []*discordgo.ApplicationCommand
	GetHandler() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type vannila struct {
	handler map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
	command []*discordgo.ApplicationCommand
}

func NewBotVannila() IVannila {
	return &vannila{
		command: vannilaCmd,
		handler: vannilaHdl,
	}
}
