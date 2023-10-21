package events

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	events "github.com/y2hO0ol23/weiver/handler/events/include"
	db "github.com/y2hO0ol23/weiver/utils/database"
	"github.com/y2hO0ol23/weiver/utils/role"
)

func init() {
	events.List = append(events.List, func(s *discordgo.Session, g *discordgo.GuildMemberUpdate) {
		guildDB, err := db.LoadGuildByID(g.GuildID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
		if guildDB.AllowRole != true {
			return
		}

		var (
			needCurrentRole bool = true
			display         string
		)
		display, err = role.GetDisplay(g.User.ID)
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
			err = role.Set(s, g.GuildID, g.User.ID, display)
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				return
			}
		}
	})
}
