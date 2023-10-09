package slash_commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/utils"
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
			subjectId := options[0].Value.(string)

			subject, err := s.State.Member(i.GuildID, subjectId)
			if err != nil {
				return // can not find subject
			}

			modal := utils.Modal(&discordgo.InteractionResponseData{
				CustomID: "review#" + i.Interaction.Member.User.ID + "#" + subjectId,
				Title:    "Review " + subject.User.Username,
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID:    "opinion",
								Label:       "What is your opinion on them?",
								Style:       discordgo.TextInputShort,
								Placeholder: "Don't be shy, share your opinion with us",
								Required:    true,
								MaxLength:   300,
								MinLength:   10,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID:  "suggestions",
								Label:     "What would you suggest to improve them?",
								Style:     discordgo.TextInputParagraph,
								Required:  false,
								MaxLength: 2000,
							},
						},
					},
				},
			})

			err = s.InteractionRespond(i.Interaction, modal)
			if err != nil {
				log.Fatalf("%v", err)
			}
		},
	})
}
