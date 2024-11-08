package refinn

import "github.com/bwmarrin/discordgo"

// IRefinn เป็น interface ที่กำหนด method เพื่อดึงข้อมูลคำสั่งและ handler สำหรับคำสั่ง refinn
type IRefinn interface {
	GetCommand() []*discordgo.ApplicationCommand
	GetHandler() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

// refinn struct เก็บข้อมูล handler ของคำสั่งและข้อมูลคำสั่ง
type refinn struct {
	handler map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
	command []*discordgo.ApplicationCommand
}

// NewBotVannila สร้าง instance ใหม่ของ refinn และกำหนดค่าเริ่มต้นให้กับคำสั่งและ handler
func NewBotRefinn() IRefinn {
	return &refinn{
		command: refinnCmd, // กำหนดคำสั่งเริ่มต้น
		handler: refinnHdl, // กำหนด handler เริ่มต้น
	}
}
