package events

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/database"
	TagUtils "github.com/y2hO0ol23/weiver/utils/bot/tag"

	g "github.com/y2hO0ol23/weiver/handler"
)

func init() {
	g.EventList = append(g.EventList, func(s *discordgo.Session, g *discordgo.GuildMemberUpdate) {
		guildDB, err := db.LoadGuildByID(g.GuildID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
		if !guildDB.AllowRole {
			return
		}

		var (
			needCurrentRole bool = true
			display         string
		)
		display, err = TagUtils.GetScoreUIShort(g.User.ID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}

		for _, roleID := range g.Member.Roles {
			roleDB, err := db.GetRoleByID(fmt.Sprintf("%v#%v", g.GuildID, roleID))
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				return
			}
			if roleDB != nil {
				if roleDB.Display == display {
					needCurrentRole = true
				} else {
					s.GuildMemberRoleRemove(g.GuildID, g.User.ID, roleID)
				}
			}
		}
		if needCurrentRole {
			err = TagUtils.AddTag(s, g.GuildID, g.User.ID, display)
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				return
			}
		}
	})
}
