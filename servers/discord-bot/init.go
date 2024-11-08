package discordbot

import (
	"github.com/bwmarrin/discordgo"
	slashcmds "github.com/xpo-nential/notify-service/servers/discord-bot/slash-cmds"
)

func Init() {
	slCmd := slashcmds.SlashCommand

	slCmd.HandleCommands()
	slCmd.RegisterCommand()

	newConnection().Start(func(s *discordgo.Session) {

	})

}

// type Notify interface {
// 	Message()
// }

// type DiscordBot struct {

// }

// type LineBot struct{

// }

// func (v *DiscordBot) Message(){

// }

// func (v *LineBot) Message(){

// }

// var VanilaDcBot Notify = &DiscordBot{

// }
// var VanilaLnBot Notify = &LineBot{

// }

// var Refinbot Notify = &LineBot{}
