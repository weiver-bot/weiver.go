package slashcommands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/handler/slash-commands/admin"
	"github.com/y2hO0ol23/weiver/localization"

	g "github.com/y2hO0ol23/weiver/handler"
)

func init() {
	var (
		DMPermission             bool  = false
		DefaultMemberPermissions int64 = discordgo.PermissionManageRoles
	)

	g.CMDList = append(g.CMDList, g.CMDForm{
		Data: &discordgo.ApplicationCommand{
			Name:                     "admin",
			Description:              "admin_Description",
			NameLocalizations:        localization.LoadList("#admin"),
			DescriptionLocalizations: localization.LoadList("#admin.Description"),
			DMPermission:             &DMPermission,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:                     "allow-role",
					Description:              "allow-role_Description",
					NameLocalizations:        *localization.LoadList("#admin.allow-role"),
					DescriptionLocalizations: *localization.LoadList("#admin.allow-role.Description"),
					Type:                     discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:                     "value",
							Description:              "value_Description",
							NameLocalizations:        *localization.LoadList("#admin.allow-role.value"),
							DescriptionLocalizations: *localization.LoadList("#admin.allow-role.value.Description"),
							Type:                     discordgo.ApplicationCommandOptionBoolean,
							Choices: []*discordgo.ApplicationCommandOptionChoice{
								{
									Name:              "true",
									NameLocalizations: *localization.LoadList("#admin.allow-role.value.true"),
									Value:             true,
								},
								{
									Name:              "false",
									NameLocalizations: *localization.LoadList("#admin.allow-role.value.false"),
									Value:             false,
								},
							},
							Required: true,
						},
					},
				},
			},
			DefaultMemberPermissions: &DefaultMemberPermissions,
		},
		Execute: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			if len(options) == 0 {
				return
			}
			cmdName := options[0].Name

			switch cmdName {
			case "allow-role":
				admin.AllowRole(s, i, options[0].Options)
			}
		},
	})
}
