package vannila

import "github.com/bwmarrin/discordgo"

func (v *Vannila) GetCommand() []*discordgo.ApplicationCommand {
	return v.command
}

var vannilaCmd = []*discordgo.ApplicationCommand{
	{
		Name: "basic-vannila",
		// All commands and options must have a description
		// Commands/options without description will fail the registration
		// of the command.
		Description: "Basic vannila",
	},
	{
		Name:        `today`,
		Description: "No description",
	},
}
