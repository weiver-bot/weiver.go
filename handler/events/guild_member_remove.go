package events

import (
	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/utils/database"
	"github.com/y2hO0ol23/weiver/utils/role"
)

func init() {
	events = append(events, func(s *discordgo.Session, g *discordgo.GuildMemberRemove) {
		guildDB := db.LoadGuildByID(g.GuildID)
		if guildDB.AllowRole != true {
			return
		}

		role.Remove(s, g.GuildID, g.Member.User.ID, role.GetDisplay(g.Member.User.ID))
	})
}
