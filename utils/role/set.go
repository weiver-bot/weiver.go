package role

import (
	"log"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/utils/database"
)

func Set(s *discordgo.Session, guildID string, memberID string, display string) {
	roleDB := db.GetRoleByInfo(guildID, display)
	if roleDB == nil {
		role, err := s.GuildRoleCreate(guildID, &discordgo.RoleParams{Name: display})
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}

		roleDB = db.CreateRole(role.ID, guildID, display)
	}

	s.GuildMemberRoleAdd(guildID, memberID, roleDB.RoleID)
	db.AddRoleOnUser(roleDB.ID, memberID)
}
