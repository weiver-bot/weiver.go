package slashcommands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/localization"

	g "github.com/y2hO0ol23/weiver/handler"
	this "github.com/y2hO0ol23/weiver/handler/slash-commands/review"
)

func init() {
	var DMPermission bool = false

	g.CMDList = append(g.CMDList, g.CMDForm{
		Data: &discordgo.ApplicationCommand{
			Name:                     "review",
			Description:              "review_Description",
			NameLocalizations:        localization.LoadList("#review"),
			DescriptionLocalizations: localization.LoadList("#review.Description"),
			DMPermission:             &DMPermission,
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
		Slash: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			locale := i.Locale

			options := i.ApplicationCommandData().Options
			subjectID := options[0].Value.(string)

			this.Execute(s, i, locale, subjectID)
		},
		Message: func(s *discordgo.Session, i *discordgo.InteractionCreate, locale discordgo.Locale, queries []string) string {
			if len(queries) < 1 {
				return fmt.Sprintf("`/%v @%v`", localization.Load(locale, "#review"), localization.Load(locale, "#.subject"))
			}

			if id := g.ParseOptionUser(s, i.GuildID, queries[0]); id != nil {
				this.Execute(s, i, locale, *id)
				return ""
			}
			return fmt.Sprintf("`/%v @%v`", localization.Load(locale, "#review"), localization.Load(locale, "#.subject"))
		},
	})
}
