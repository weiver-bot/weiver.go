package slashcommands

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/database"
	"github.com/y2hO0ol23/weiver/localization"
	ReviewUtils "github.com/y2hO0ol23/weiver/utils/bot/review"
	"github.com/y2hO0ol23/weiver/utils/builder"

	g "github.com/y2hO0ol23/weiver/handler"
)

func init() {
	var (
		DMPermission bool = false
	)

	g.CMDList = append(g.CMDList, g.CMDForm{
		Data: &discordgo.ApplicationCommand{
			Name:                     "move",
			Description:              "move_Description",
			NameLocalizations:        localization.LoadList("#move"),
			DescriptionLocalizations: localization.LoadList("#move.Description"),
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

			review, err := db.LoadReivewByInfo(authorID, subjectID)
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			}
			if review == nil {
				err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
					Content:         fmt.Sprintf("`%v`", localization.Load(locale, "#move.IsNone")),
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
				return
			}

			_, err = s.ChannelMessage(review.ChannelID, review.MessageID)
			if err != nil {
				review, err := ReviewUtils.SendReview(s, i, review)
				if err != nil {
					log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
					return
				}
				if review != nil {
					ReviewUtils.ModifyReviewDM(s, review, locale)
				}
				return
			}

			err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					ReviewUtils.BaseEmbed(review, subject.AvatarURL("")).
						SetDescription(fmt.Sprintf("https://discord.com/channels/%v/%v/%v", review.GuildID, review.ChannelID, review.MessageID)).
						MessageEmbed,
				},
				Components: []discordgo.MessageComponent{
					builder.ActionRow().AddComponents(
						builder.Button().
							SetCustomID("move-review").
							SetLable(localization.Load(locale, "#move.Move")).
							SetStyle(discordgo.SuccessButton),
					).ActionsRow,
				},
				Flags:           discordgo.MessageFlagsEphemeral,
				AllowedMentions: &discordgo.MessageAllowedMentions{},
			}))
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				return
			}

			msg, err := s.InteractionResponse(i.Interaction)
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				return
			}

			var handler func(*discordgo.Session, *discordgo.InteractionCreate)
			handler = func(s *discordgo.Session, iter *discordgo.InteractionCreate) {
				if iter.Type != discordgo.InteractionMessageComponent || i.Interaction.Member.User.ID != iter.Interaction.Member.User.ID {
					s.AddHandlerOnce(handler)
					return
				}
				if iter.Interaction.Message.ID != msg.ID {
					return
				}

				data := iter.MessageComponentData()
				if data.ComponentType != discordgo.ButtonComponent || data.CustomID != "move-review" {
					return
				}

				s.InteractionResponseDelete(i.Interaction)
				reviewNow, err := db.LoadReivewByInfo(authorID, subjectID)
				if err != nil {
					log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
					return
				}
				if review.TimeStamp != reviewNow.TimeStamp {
					err := s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{
							builder.Embed().
								SetDescription(fmt.Sprintf("❌ %v", localization.Load(locale, "$review.IsEdited"))).
								MessageEmbed,
						},
					}))
					if err != nil {
						log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
					}
					return
				}
				err = ReviewUtils.DeleteLastMessage(s, authorID, subjectID)
				if err != nil {
					log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				}
				review, err = ReviewUtils.SendReview(s, iter, review)
				if err != nil {
					log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				}
				if review != nil {
					ReviewUtils.ModifyReviewDM(s, review, locale)
				}
			}
			s.AddHandlerOnce(handler)
		},
	})
}
