package slashcmds

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/xpo-nential/notify-service/notify-service/bot-discord/commands"
)

// RegisterCommand ลงทะเบียนคำสั่ง ApplicationCommand ตามตัวแปรสภาพแวดล้อม APP_NAME
func (sc *slashCommands) RegisterCommand() {
	var command []*discordgo.ApplicationCommand

	// เลือกคำสั่งตามตัวแปรสภาพแวดล้อม APP_NAME
	switch strings.ToLower(os.Getenv(`APP_NAME`)) {
	case strings.ToLower(`vannila`):
		command = sc.vannila.GetCommand()
	case strings.ToLower(`refinn`):
		command = sc.refinn.GetCommand()
	default:
		command = []*discordgo.ApplicationCommand{}
	}

	// เพิ่มคำสั่งที่เลือกใช้กับ NewSlashCommand ใน package commands
	commands.NewSlashCommand.AddCommand(command)
}
