package role

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/utils/database"
)

func Set(s *discordgo.Session, guildID string, memberID string, display string) {
	roleDB := db.GetRoleByInfo(guildID, display)
	if roleDB == nil {
		role, err := s.GuildRoleCreate(guildID, &discordgo.RoleParams{Name: display})
		if err != nil {
			fmt.Println(err)
			return
		}

		roleDB = db.CreateRole(role.ID, guildID, display)
	}

	err := s.GuildMemberRoleAdd(guildID, memberID, roleDB.RoleID)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AddRoleOnUser(roleDB.ID, memberID)
}
