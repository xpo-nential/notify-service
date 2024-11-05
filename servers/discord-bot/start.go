package discordbot

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/xpo-nential/notify-service/configs"
	"github.com/xpo-nential/notify-service/servers/discord-bot/commands"
)

// กำหนดตัวแปร Bot
var (
	// GuildID ใช้กำหนดไอดีของเซิร์ฟเวอร์ทดสอบ ถ้าไม่ระบุจะลงทะเบียนคำสั่งแบบ global
	GuildID = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	// RemoveCommands กำหนดว่าจะลบคำสั่งทั้งหมดหลังจากปิดบอทหรือไม่
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

// IDiscordServer กำหนด interface ที่มีเมธอด Start สำหรับเริ่มการทำงานของบอท
type IDiscordServer interface {
	Start(handler func(s *discordgo.Session))
}

// discordServer struct เก็บข้อมูลเกี่ยวกับ session ของบอทและคำสั่งที่ใช้
type discordServer struct {
	session  *discordgo.Session
	commands []*discordgo.ApplicationCommand
}

// NewConnection สร้างการเชื่อมต่อใหม่กับ Discord โดยใช้ token ที่กำหนด
func NewConnection() IDiscordServer {
	// ดึง token จากการตั้งค่า
	token := configs.NewConfig(configs.DiscordBot).App().GetToken()
	// สร้าง session ของ Discord ด้วย token ที่กำหนด
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	return &discordServer{
		session:  session,
		commands: commands.NewCommands().GetCommands(),
	}
}

// Start เริ่มการทำงานของบอท
func (ds *discordServer) Start(handler func(s *discordgo.Session)) {
	session := ds.session

	// เพิ่ม handler สำหรับเหตุการณ์ Ready เพื่อแสดงข้อความเมื่อบอทเข้าสู่ระบบสำเร็จ
	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	// ลงทะเบียน handler ที่กำหนดโดยผู้ใช้
	handler(session)

	// กำหนด Intents ที่ต้องการให้บอทเข้าถึง
	session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	// เปิดการเชื่อมต่อกับ Discord และเริ่มฟังเหตุการณ์
	if err := session.Open(); err != nil {
		fmt.Println("error opening connection:", err)
	}

	// เริ่มต้นโมดูลและใช้ module สำหรับจัดการคำสั่งต่าง ๆ
	module := ModuleInit(ds)

	// เพิ่ม handler สำหรับเหตุการณ์ InteractionCreate เพื่อจัดการคำสั่งที่ผู้ใช้ส่งมา
	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := module.GetCommandHandler()[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	// ลงทะเบียนคำสั่งกับ Discord
	log.Println("Adding commands...")
	commands := ds.commands
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, *GuildID, v)
		if err != nil {
			log.Printf("Cannot create '%v' command: %v", v.Name, err)
			continue
		}
		registeredCommands[i] = cmd
	}

	// แสดงรายการคำสั่งที่ลงทะเบียนแล้ว
	cmds, _ := session.ApplicationCommands(session.State.User.ID, *GuildID)
	for _, v := range cmds {
		fmt.Println("cmd:", v.Name)
	}

	// ปิด session และลบคำสั่งเมื่อบอทถูกปิด
	defer func() {
		session.Close()
		if *RemoveCommands {
			log.Println("Removing commands...")
			for _, v := range cmds {
				err := session.ApplicationCommandDelete(session.State.User.ID, *GuildID, v.ID)
				if err != nil {
					log.Printf("Cannot delete '%v' command: %v", v.Name, err)
				}
			}
		}
		log.Println("Gracefully shutting down.")
	}()

	// รอจนกว่าจะมีสัญญาณหยุดการทำงาน (เช่น กด CTRL+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	<-stop
}
