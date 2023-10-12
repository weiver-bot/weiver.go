package slash_commands

import (
	"log"
	"sync"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/utils/database"
	"github.com/y2hO0ol23/weiver/utils/role"
)

func init() {
	var (
		DMPermission             bool  = false
		DefaultMemberPermissions int64 = discordgo.PermissionManageRoles
	)

	commands = append(commands, form{
		data: &discordgo.ApplicationCommand{
			Name:         "allow_role",
			Description:  "FOR ADMIN - DEFAULT:False",
			DMPermission: &DMPermission,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "value",
					Description: "Of role update",
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "true",
							Value: true,
						},
						{
							Name:  "false",
							Value: false,
						},
					},
					Required: true,
				},
			},
			DefaultMemberPermissions: &DefaultMemberPermissions,
		},
		execute: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			value := options[0].Value.(bool)

			guildDB := db.LoadGuildByID(i.GuildID)
			if guildDB.AllowRole != value {
				db.UpdateGuildRoleOption(i.GuildID, value)

				var (
					after string
					wait  sync.WaitGroup
				)
				for {
					res, err := s.GuildMembers(i.GuildID, after, 1000)
					if err != nil {
						log.Println(err)
						continue
					}
					if len(res) == 0 {
						break
					}

					wait.Add(1)
					go func(members []*discordgo.Member) {
						defer wait.Done()
						for _, member := range members {
							if value == true {
								role.Set(s, i.GuildID, member.User.ID, role.GetDisplay(member.User.ID))
							} else {
								role.Remove(s, i.GuildID, member.User.ID, role.GetDisplay(member.User.ID))
							}
						}
					}(res)

					after = res[len(res)-1].User.ID
				}
				wait.Wait()
			}
		},
	})
}
