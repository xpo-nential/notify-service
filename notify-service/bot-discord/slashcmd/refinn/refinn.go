package refinn

import "github.com/bwmarrin/discordgo"

// refinn struct เก็บข้อมูล handler ของคำสั่งและข้อมูลคำสั่ง
type Refinn struct {
	handler map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
	command []*discordgo.ApplicationCommand
}

// NewBotVannila สร้าง instance ใหม่ของ refinn และกำหนดค่าเริ่มต้นให้กับคำสั่งและ handler
func NewBotRefinn() *Refinn {
	return &Refinn{
		command: refinnCmd, // กำหนดคำสั่งเริ่มต้น
		handler: refinnHdl, // กำหนด handler เริ่มต้น
	}
}
