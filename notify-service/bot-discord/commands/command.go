package commands

import "github.com/bwmarrin/discordgo"

// command struct เก็บ slice ของคำสั่ง (commands) ที่จะถูกใช้งานโดยบอท
type command struct {
	commands []*discordgo.ApplicationCommand
}

// instance เป็นตัวแปร singleton ของ command เพื่อให้แน่ใจว่าใช้ instance เดียวกันตลอดการทำงาน
var instance *command

// NewCommands สร้างและคืนค่า instance ของ command แบบ singleton
// เพื่อให้แน่ใจว่ามีเพียงหนึ่ง instance ของ command ที่ถูกสร้างและใช้งาน
func NewCommands() *command {
	if instance == nil {
		instance = &command{
			commands: make([]*discordgo.ApplicationCommand, 0),
		}
	}
	return instance
}

// SetCommands ใช้ในการตั้งค่าคำสั่งใหม่ให้กับ command instance
func (c *command) SetCommands(cmds []*discordgo.ApplicationCommand) {
	c.commands = cmds
}

// GetCommands คืนค่า slice ของคำสั่งที่เก็บอยู่ใน command instance
func (c *command) GetCommands() []*discordgo.ApplicationCommand {
	return c.commands
}
