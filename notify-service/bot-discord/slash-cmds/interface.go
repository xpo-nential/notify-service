package slashcmds

import (
	"github.com/xpo-nential/notify-service/notify-service/bot-discord/slash-cmds/refinn"
	"github.com/xpo-nential/notify-service/notify-service/bot-discord/slash-cmds/vannila"
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
	vannila vannila.IVannila
	refinn  refinn.IRefinn
}
