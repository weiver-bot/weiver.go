package slashcommands

import (
	"fmt"
	"log"
	"runtime/debug"
	"strings"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/database"
	"github.com/y2hO0ol23/weiver/localization"
	"github.com/y2hO0ol23/weiver/utils/builder"

	g "github.com/y2hO0ol23/weiver/handler"
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
		Execute: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			locale := i.Locale

			options := i.ApplicationCommandData().Options
			authorID := i.Interaction.Member.User.ID
			subjectID := options[0].Value.(string)

			if authorID == subjectID {
				err := s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
					Content:         fmt.Sprintf("`%v`", localization.Load(locale, "#review.SelfReview")),
					Flags:           discordgo.MessageFlagsEphemeral,
					AllowedMentions: &discordgo.MessageAllowedMentions{},
				}))
				if err != nil {
					log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				}
				return
			}

			subject, err := s.GuildMember(i.GuildID, subjectID)
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				return // can not find subject
			}

			review, err := db.LoadReivewByInfo(authorID, subjectID)
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				return
			}

			modal := builder.Modal().
				SetCustomID("review#" + authorID + "#" + subjectID).
				SetTitle(fmt.Sprintf(localization.Load(locale, "#review.modal.Title"), subject.User.Username))

			score := builder.TextInput().
				SetCustomID("score").
				SetLable(localization.Load(locale, "#review.lable.Score")).
				SetValue(func() string {
					if review == nil {
						return "★★★★★"
					}
					return strings.Repeat("★", review.Score)
				}()).
				SetStyle(discordgo.TextInputShort).
				SetMinLength(1).SetMaxLength(5).SetRequired(true)

			title := builder.TextInput().
				SetCustomID("title").
				SetLable(localization.Load(locale, "#review.lable.Title")).
				SetValue(func() string {
					if review == nil {
						return ""
					}
					return review.Title
				}()).
				SetStyle(discordgo.TextInputShort).
				SetMinLength(1).SetMaxLength(20).SetRequired(true)

			content := builder.TextInput().
				SetCustomID("content").
				SetLable(localization.Load(locale, "#review.lable.Content")).
				SetValue(func() string {
					if review == nil {
						return ""
					}
					return review.Content
				}()).
				SetStyle(discordgo.TextInputParagraph).
				SetMinLength(1).SetMaxLength(300).SetRequired(true)

			modal.AddComponents(
				builder.ActionRow().AddComponents(score),
				builder.ActionRow().AddComponents(title),
				builder.ActionRow().AddComponents(content),
			)

			err = s.InteractionRespond(i.Interaction, modal.InteractionResponse)
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			}
		},
	})
}
