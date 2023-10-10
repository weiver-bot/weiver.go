package slash_commands

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/utils/builder"
	"github.com/y2hO0ol23/weiver/utils/prisma"
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
				message := builder.Message(&discordgo.InteractionResponseData{
					Content: "`Can't review yourself`",
					Flags:   discordgo.MessageFlagsEphemeral,
				})
				err := s.InteractionRespond(i.Interaction, message)
				if err != nil {
					log.Printf("Error on sending message\n")
				}
				return
			}

			to, err := s.GuildMember(i.GuildID, toId)
			if err != nil {
				log.Println("Error on loadding user")
				return // can not find subject
			}

			review_db := prisma.LoadReivewByIds(fromId, toId)

			modal := builder.Modal().
				SetCustomId("review#" + fromId + "#" + toId).
				SetTitle("Review " + to.User.Username)

			score := builder.TextInput().
				SetCustomId("score").
				SetLable("score").
				SetValue(func() string {
					if review_db == nil {
						return "★★★★★"
					}
					return strings.Repeat("★", review_db.Score)
				}()).
				SetStyle(discordgo.TextInputShort).
				SetMinLength(1).SetMaxLength(5).SetRequired(true)

			title := builder.TextInput().
				SetCustomId("title").
				SetLable("title").
				SetValue(func() string {
					if review_db == nil {
						return ""
					}
					return review_db.Title
				}()).
				SetStyle(discordgo.TextInputShort).
				SetMinLength(1).SetMaxLength(20).SetRequired(true)

			content := builder.TextInput().
				SetCustomId("content").
				SetLable("content").
				SetValue(func() string {
					if review_db == nil {
						return ""
					}
					return review_db.Title
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
				log.Printf("Error on sending modal\n")
			}
		},
	})
}
