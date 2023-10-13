package role

import (
	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/utils/database"
)

func Remove(s *discordgo.Session, guildID string, memberID string, display string) error {
	roleDB, err := db.GetRoleByInfo(guildID, display)
	if err != nil {
		return err
	}
	if roleDB == nil {
		return nil
	}

	ok, err := db.RemoveRoleOnUser(roleDB.ID, memberID)
	if err != nil {
		return err
	}

	if !ok {
		s.GuildRoleDelete(guildID, roleDB.RoleID)
	}
	s.GuildMemberRoleRemove(guildID, memberID, roleDB.RoleID)

	return nil
}
