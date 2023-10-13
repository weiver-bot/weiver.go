package events

import (
	"log"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/utils/database"
	"github.com/y2hO0ol23/weiver/utils/role"
)

func init() {
	events = append(events, func(s *discordgo.Session, g *discordgo.GuildMemberRemove) {
		guildDB, err := db.LoadGuildByID(g.GuildID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
		if guildDB.AllowRole != true {
			return
		}

		display, err := role.GetDisplay(g.Member.User.ID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
		err = role.Remove(s, g.GuildID, g.Member.User.ID, display)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
	})
}
