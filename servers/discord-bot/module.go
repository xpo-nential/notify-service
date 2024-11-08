package discordbot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/xpo-nential/notify-service/servers/discord-bot/handlers"
)

// Imodule กำหนด interface ที่มีเมธอดสำหรับดึงตัวจัดการคำสั่ง (command handlers)
type Imodule interface {
	GetCommandHandler() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

// module struct แทนโมดูลที่มีการฝัง discordServer และใช้ commandHandler ในการจัดการคำสั่ง
type module struct {
	*discordServer
	commandHandler handlers.IHandler
}

// instance เป็นตัวแปร singleton เพื่อให้แน่ใจว่ามีเพียงหนึ่ง instance ของ module เท่านั้น
var instance *module

// ModuleInit ทำการสร้างและคืนค่า instance ของ module แบบ singleton
// เพื่อให้แน่ใจว่ามีการสร้างเพียงหนึ่งครั้งและสามารถใช้ซ้ำได้
func moduleInit(ds *discordServer) Imodule {
	if instance == nil {
		instance = &module{
			discordServer:  ds,
			commandHandler: handlers.NewHandlers(),
		}
	}
	return instance
}

// GetCommandHandler คืนค่าแผนที่ของ command handlers จาก commandHandler instance
func (m *module) GetCommandHandler() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return m.commandHandler.GetHandler()
}
