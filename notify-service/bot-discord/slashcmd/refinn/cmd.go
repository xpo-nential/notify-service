package refinn

import "github.com/bwmarrin/discordgo"

// GetCommand คืนค่ารายการคำสั่ง ApplicationCommand สำหรับ refinn
func (v *Refinn) GetCommand() []*discordgo.ApplicationCommand {
	return v.command
}

// refinnCmd เก็บคำสั่งพร้อมกับคำอธิบายที่จำเป็น
var refinnCmd = []*discordgo.ApplicationCommand{
	{
		Name:        "basic-refinn",
		Description: "Basic refinn", // จำเป็นต้องมีคำอธิบายสำหรับการลงทะเบียนคำสั่ง
	},
}
