package events

import (
	"log"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/database"
	TagUtils "github.com/y2hO0ol23/weiver/utils/bot/tag"

	g "github.com/y2hO0ol23/weiver/handler"
)

func init() {
	g.EventList = append(g.EventList, func(s *discordgo.Session, g *discordgo.GuildMemberAdd) {
		guildDB, err := db.LoadGuildByID(g.GuildID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
		if !guildDB.AllowRole {
			return
		}

		display, err := TagUtils.GetScoreUIShort(g.Member.User.ID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
		err = TagUtils.AddTag(s, g.GuildID, g.Member.User.ID, display)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
	})
}
