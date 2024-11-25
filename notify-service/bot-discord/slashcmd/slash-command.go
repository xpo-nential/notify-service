package slashcmd

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/xpo-nential/notify-service/notify-service/bot-discord/commands"
	"github.com/xpo-nential/notify-service/notify-service/bot-discord/handlers"
	"github.com/xpo-nential/notify-service/notify-service/bot-discord/slashcmd/refinn"
	"github.com/xpo-nential/notify-service/notify-service/bot-discord/slashcmd/vannila"
)

// ISlashCommands เป็น interface ที่กำหนด method สำหรับจัดการคำสั่งแบบ Slash Commands
type ISlashCommands interface {
	HandleCommands()
	RegisterCommand()
}

// SlashCommandHandler เป็น instance ที่ใช้จัดการคำสั่งแบบ Slash Commands สำหรับแอปพลิเคชัน "vannila" และ "refinn"
var SlashCommand ISlashCommands = &slashCommands{
	vannila.NewBotVannila(),
	refinn.NewBotRefinn(),
}

// slashCommands struct เก็บ instance สำหรับจัดการคำสั่ง "vannila" และ "refinn"
type slashCommands struct {
	vannila *vannila.Vannila
	refinn  *refinn.Refinn
}

// HandleCommands เป็นเมธอดที่กำหนด handler ของคำสั่งแบบ Slash Command ตามชื่อแอปพลิเคชัน
func (sc *slashCommands) HandleCommands() {
	var handler map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)

	// เลือก handler ตามตัวแปรสภาพแวดล้อม APP_NAME
	switch strings.ToLower(os.Getenv(`APP_NAME`)) {
	case strings.ToLower(`vannila`):
		handler = sc.vannila.GetHandler()
	case strings.ToLower(`refinn`):
		handler = sc.refinn.GetHandler()
	default:
		handler = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}
	}

	// เพิ่ม handler ที่เลือกใช้กับ NewSlashHandler ใน package handlers
	handlers.NewSlashHandler.AddHandlers(handler)
}

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
