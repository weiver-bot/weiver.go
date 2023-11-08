package slashcommands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	this "github.com/y2hO0ol23/weiver/handler/slash-commands/look"
	"github.com/y2hO0ol23/weiver/localization"

	g "github.com/y2hO0ol23/weiver/handler"
)

func init() {
	var DMPermission bool = false

	g.CMDList = append(g.CMDList, g.CMDForm{
		Data: &discordgo.ApplicationCommand{
			Name:                     "look",
			Description:              "look_Description",
			NameLocalizations:        localization.LoadList("#look"),
			DescriptionLocalizations: localization.LoadList("#look.Description"),
			DMPermission:             &DMPermission,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:                     "info",
					Description:              "info_Description",
					NameLocalizations:        *localization.LoadList("#look.info"),
					DescriptionLocalizations: *localization.LoadList("#look.info.Description"),
					Type:                     discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:                     "subject",
							Description:              "subject_Description",
							NameLocalizations:        *localization.LoadList("#.subject"),
							DescriptionLocalizations: *localization.LoadList("#.subject.Description"),
							Type:                     discordgo.ApplicationCommandOptionUser,
							Required:                 true,
						},
					},
				},
				{
					Name:                     "reviews",
					Description:              "reviews_Description",
					NameLocalizations:        *localization.LoadList("#look.reviews"),
					DescriptionLocalizations: *localization.LoadList("#look.reviews.Description"),
					Type:                     discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:                     "subject",
							Description:              "subject_Description",
							NameLocalizations:        *localization.LoadList("#.subject"),
							DescriptionLocalizations: *localization.LoadList("#.subject.Description"),
							Type:                     discordgo.ApplicationCommandOptionUser,
							Required:                 true,
						},
					},
				},
			},
		},
		Slash: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			if len(options) == 0 {
				return
			}
			cmdName := options[0].Name

			switch cmdName {
			case "info":
				this.Info(s, i, options[0].Options[0].Value.(string))
			case "reviews":
				this.Reviews(s, i, options[0].Options[0].Value.(string))
			}
		},
		Message: func(s *discordgo.Session, i *discordgo.InteractionCreate, locale discordgo.Locale, queries []string) string {
			if len(queries) < 2 {
				return fmt.Sprintf(
					"`/%v %v ... or\n/%v %v ...`",
					localization.Load(locale, "#look"), localization.Load(locale, "#look.info"),
					localization.Load(locale, "#look"), localization.Load(locale, "#look.reviews"),
				)
			}

			switch queries[0] {
			case "info":
				if id := g.ParseOptionUser(s, i.GuildID, queries[1]); id != nil {
					this.Info(s, i, *id)
				} else {
					return fmt.Sprintf(
						"`/%v %v @%v`",
						localization.Load(locale, "#look"),
						localization.Load(locale, "#look.info"),
						localization.Load(locale, "#.subject"),
					)
				}
			case "reviews":
				if id := g.ParseOptionUser(s, i.GuildID, queries[1]); id != nil {
					this.Reviews(s, i, *id)
				} else {
					return fmt.Sprintf(
						"`/%v %v @%v`",
						localization.Load(locale, "#look"),
						localization.Load(locale, "#look.reviews"),
						localization.Load(locale, "#.subject"),
					)
				}
			default:
				return fmt.Sprintf(
					"`/%v %v ... or\n/%v %v ...`",
					localization.Load(locale, "#look"), localization.Load(locale, "#look.info"),
					localization.Load(locale, "#look"), localization.Load(locale, "#look.reviews"),
				)
			}
			return ""
		},
	})
}
