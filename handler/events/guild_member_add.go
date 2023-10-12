package events

import (
	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/utils/database"
	"github.com/y2hO0ol23/weiver/utils/role"
)

func init() {
	events = append(events, func(s *discordgo.Session, g *discordgo.GuildMemberAdd) {
		guildDB := db.LoadGuildByID(g.GuildID)
		if guildDB.AllowRole != true {
			return
		}

		role.Set(s, g.GuildID, g.Member.User.ID, role.GetDisplay(g.Member.User.ID))
	})
}
