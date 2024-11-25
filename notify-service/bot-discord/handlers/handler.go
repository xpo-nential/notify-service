package handlers

import "github.com/bwmarrin/discordgo"

// handler struct เก็บแผนที่ของ handlers เพื่อจัดการคำสั่งต่าง ๆ
type Handler struct {
	handlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

// instance เป็นตัวแปร singleton เพื่อให้แน่ใจว่ามีเพียงหนึ่ง instance ของ handler เท่านั้น
var instance *Handler

// NewHandlers สร้างและคืนค่า instance ของ handler แบบ singleton
// เพื่อให้แน่ใจว่าใช้ instance เดียวกันในการจัดการคำสั่งตลอดแอปพลิเคชัน
func NewHandlers() *Handler {
	if instance == nil {
		instance = &Handler{
			handlers: make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)),
		}
	}
	return instance
}

// GetHandler คืนค่าแผนที่ของ handlers ที่เก็บอยู่ใน handler instance
func (h *Handler) GetHandler() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return h.handlers
}

// SetHandler เพิ่มหรืออัพเดต handlers ในแผนที่ของ handler ที่มีอยู่
func (h *Handler) SetHandler(hdls map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	for k, v := range hdls {
		h.handlers[k] = v
	}
}
