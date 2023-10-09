package slash_commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/db"
	"github.com/y2hO0ol23/weiver/utils/prisma"
	"github.com/y2hO0ol23/weiver/utils/simple"
)

var (
	dmPermission = false
	err          error
)

func init() {
	commands = append(commands, form{
		data: &discordgo.ApplicationCommand{
			Name:         "reivew",
			Description:  "review user",
			DMPermission: &dmPermission,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "subject",
					Description: "Select subject",
					Type:        discordgo.ApplicationCommandOptionUser,
					Required:    true,
				},
			},
		},
		execute: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			fromId := i.Interaction.Member.User.ID
			toId := options[0].Value.(string)

			if fromId == toId {
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "`Can't review yourself`",
						Flags:   discordgo.MessageFlagsEphemeral,
					},
				})
				if err != nil {
					panic(err)
				}
				return
			}

			to, err := s.State.Member(i.GuildID, toId)
			if err != nil || to == nil {
				return // can not find subject
			}

			review_db := prisma.LoadReivewByIds(fromId, toId)

			modal := simple.Modal(&discordgo.InteractionResponseData{
				CustomID: "review#" + fromId + "#" + toId,
				Title:    "Review " + to.User.Username,
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "score",
								Label:    "score",
								Value: func(db *db.ReviewModel) string {
									if review_db == nil {
										return "★★★★★"
									}
									return strings.Repeat("★", review_db.Score)
								}(review_db),
								Style:     discordgo.TextInputShort,
								MinLength: 1, MaxLength: 5, Required: true,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "title",
								Label:    "title",
								Style:    discordgo.TextInputShort,
								Value: func(db *db.ReviewModel) string {
									if review_db == nil {
										return ""
									}
									return review_db.Title
								}(review_db),
								MinLength: 1, MaxLength: 16, Required: true,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "content",
								Label:    "content",
								Value: func(db *db.ReviewModel) string {
									if review_db == nil {
										return ""
									}
									return review_db.Content
								}(review_db),
								Style:     discordgo.TextInputParagraph,
								MinLength: 1, MaxLength: 256, Required: true,
							},
						},
					},
				},
			})

			err = s.InteractionRespond(i.Interaction, modal)
			if err != nil {
				panic(err)
			}
		},
	})
}
