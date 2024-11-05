package handlers

import "github.com/bwmarrin/discordgo"

// SlashHandler struct ใช้สำหรับจัดการ slash commands
type SlashHandler struct{}

// NewSlashHandler ให้ instance ของ SlashHandler เพื่อใช้ในการเพิ่ม command handlers
var NewSlashHandler = SlashHandler{}

// AddHandlers เพิ่มชุดของ command handlers ลงใน handler instance ที่มีอยู่
// เพื่อให้ command handlers พร้อมใช้งานในการโต้ตอบของบอท
func (SlashHandler) AddHandlers(handler map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	h := NewHandlers()
	h.SetHandler(handler)
}
