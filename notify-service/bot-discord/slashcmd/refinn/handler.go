package refinn

import "github.com/bwmarrin/discordgo"

// GetHandler คืนค่า map ของ handler สำหรับคำสั่ง refinn
func (r *Refinn) GetHandler() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return r.handler
}

// refinnHdl เป็น map ที่เก็บคำสั่งแต่ละคำสั่งให้ตรงกับฟังก์ชัน handler ของมัน
var refinnHdl = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"basic-refinn": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Handler สำหรับ "basic-command" ซึ่งจะส่งข้อความตอบกลับ
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "สวัสดี! ยินดีด้วยที่คุณได้เรียกใช้คำสั่งแรกของคุณสำเร็จแล้ว",
			},
		})
		if err != nil {
			// ส่งข้อความสำรองหากเกิดข้อผิดพลาดในการตอบกลับ
			s.ChannelMessageSend(i.ChannelID, "เกิดข้อผิดพลาดขณะตอบกลับคำสั่ง")
		}
	},
}
