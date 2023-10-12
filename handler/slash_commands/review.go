package slash_commands

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/utils/builder"
	db "github.com/y2hO0ol23/weiver/utils/database"
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
			fromID := i.Interaction.Member.User.ID
			toID := options[0].Value.(string)

			if fromID == toID {
				message := builder.Message(&discordgo.InteractionResponseData{
					Content: "`Can't review yourself`",
					Flags:   discordgo.MessageFlagsEphemeral,
				})
				err = s.InteractionRespond(i.Interaction, message)
				if err != nil {
					log.Println(err)
				}
				return
			}

			to, err := s.GuildMember(i.GuildID, toID)
			if err != nil {
				log.Println(err)
				return // can not find subject
			}

			review := db.LoadReivewByInfo(fromID, toID)

			modal := builder.Modal().
				SetCustomID("review#" + fromID + "#" + toID).
				SetTitle("Review " + to.User.Username)

			score := builder.TextInput().
				SetCustomID("score").
				SetLable("score").
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
				SetLable("title").
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
				SetLable("content").
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
				log.Println(err)
			}
		},
	})
}
