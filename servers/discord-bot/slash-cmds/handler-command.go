package slashcmds

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/xpo-nential/notify-service/servers/discord-bot/handlers"
)

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
