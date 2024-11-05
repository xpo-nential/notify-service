package commands

import "github.com/bwmarrin/discordgo"

// SlashCommand struct ใช้สำหรับจัดการคำสั่งแบบ Slash Command ใน Discord
type SlashCommand struct{}

// NewSlashCommand สร้าง instance ของ SlashCommand ที่สามารถใช้ในการเพิ่มคำสั่งได้
var NewSlashCommand = SlashCommand{}

// AddCommand ใช้สำหรับเพิ่มคำสั่งแบบ Slash Command ไปยังคำสั่งที่มีอยู่
// รับ parameter เป็น slice ของ ApplicationCommand ที่ต้องการเพิ่ม
func (SlashCommand) AddCommand(command []*discordgo.ApplicationCommand) {
	// เรียกใช้ NewCommands() เพื่อสร้าง instance ของ command และใช้ SetCommands() เพื่อตั้งค่าคำสั่งใหม่
	NewCommands().SetCommands(command)
}
