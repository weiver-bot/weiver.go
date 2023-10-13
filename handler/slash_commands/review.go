package slash_commands

import (
	"log"
	"runtime/debug"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/utils/builder"
	db "github.com/y2hO0ol23/weiver/utils/database"
)

func init() {
	var DMPermission bool = false

	commands = append(commands, form{
		data: &discordgo.ApplicationCommand{
			Name:         "reivew",
			Description:  "review user",
			DMPermission: &DMPermission,
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
				err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
					Content:         "`Can't review yourself`",
					Flags:           discordgo.MessageFlagsEphemeral,
					AllowedMentions: &discordgo.MessageAllowedMentions{},
				}))
				if err != nil {
					log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				}
				return
			}

			to, err := s.GuildMember(i.GuildID, toID)
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
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
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			}
		},
	})
}
