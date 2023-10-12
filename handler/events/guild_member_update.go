package events

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/utils/database"
	"github.com/y2hO0ol23/weiver/utils/role"
)

func init() {
	events = append(events, func(s *discordgo.Session, g *discordgo.GuildMemberUpdate) {
		guildDB := db.LoadGuildByID(g.GuildID)
		if guildDB.AllowRole != true {
			return
		}

		var (
			needCurrentRole bool   = true
			display         string = role.GetDisplay(g.User.ID)
		)

		for _, roleID := range g.Member.Roles {
			roleDB := db.GetRoleByID(fmt.Sprintf("%s#%s", g.GuildID, roleID))
			if roleDB != nil {
				if roleDB.Display == display {
					needCurrentRole = true
				} else {
					err := s.GuildMemberRoleRemove(g.GuildID, g.User.ID, roleID)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
		if needCurrentRole {
			role.Set(s, g.GuildID, g.User.ID, display)
		}
	})
}
