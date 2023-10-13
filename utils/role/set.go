package role

import (
	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/utils/database"
)

func Set(s *discordgo.Session, guildID string, memberID string, display string) error {
	roleDB, err := db.GetRoleByInfo(guildID, display)
	if err != nil {
		return err
	}
	if roleDB == nil {
		role, err := s.GuildRoleCreate(guildID, &discordgo.RoleParams{Name: display})
		if err != nil {
			return err
		}

		roleDB, err = db.CreateRole(role.ID, guildID, display)
		if err != nil {
			return err
		}
	}

	s.GuildMemberRoleAdd(guildID, memberID, roleDB.RoleID)
	return db.AddRoleOnUser(roleDB.ID, memberID)
}
