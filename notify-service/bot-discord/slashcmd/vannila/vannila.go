package vannila

import "github.com/bwmarrin/discordgo"

type Vannila struct {
	handler map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
	command []*discordgo.ApplicationCommand
}

func NewBotVannila() *Vannila {
	return &Vannila{
		command: vannilaCmd,
		handler: vannilaHdl,
	}
}
