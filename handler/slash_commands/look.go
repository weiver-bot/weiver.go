package slash_commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/utils/builder"
	"github.com/y2hO0ol23/weiver/utils/prisma"
)

func init() {
	commands = append(commands, form{
		data: &discordgo.ApplicationCommand{
			Name:        "look",
			Description: "Look about things",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "info",
					Description: "Look about user info",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "subject",
							Description: "Select subject",
							Type:        discordgo.ApplicationCommandOptionUser,
							Required:    true,
						},
					},
				},
				{
					Name:        "review-list",
					Description: "Look about reviews on user",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "subject",
							Description: "Select subject",
							Type:        discordgo.ApplicationCommandOptionUser,
							Required:    true,
						},
					},
				},
			},
		},
		execute: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			if len(options) == 0 {
				return
			}
			cmdName := options[0].Name

			switch cmdName {
			case "info":
				look_info(s, i, options[0].Options[0].Value.(string))
			case "review-list":
				look_reviewList(s, i, options[0].Options[0].Value.(string))
			}
		},
	})
}

func look_info(s *discordgo.Session, i *discordgo.InteractionCreate, subjectId string) {
	subject, err := s.GuildMember(i.GuildID, subjectId)
	if err != nil {
		log.Println("Error on loading member")
		return
	}

	scoreAverage, count := prisma.GetUserScoreAverage(subjectId)
	countOutput := func() string {
		if count >= 100 {
			return "100+"
		} else if count >= 50 {
			return "50+"
		}
		return fmt.Sprintf("%d", count)
	}()

	embed := builder.Embed().
		SetDescription(fmt.Sprintf("<@%s> **⭐%.1f (%s)**", subjectId, scoreAverage, countOutput)).
		SetThumbnail(&discordgo.MessageEmbedThumbnail{
			URL: subject.User.AvatarURL(""),
		})

	if count == 0 {
		embed.SetFields(&discordgo.MessageEmbedField{
			Name:  "📑 No reviews",
			Value: "``` ```",
		})
	} else {
		review := prisma.GetBestReview(subjectId)
		embed.SetFields(&discordgo.MessageEmbedField{
			Name:  fmt.Sprintf("📑 %s 〔%s%s〕", review.Title, "★★★★★"[:review.Score*3], "☆☆☆☆☆"[review.Score*3:]),
			Value: fmt.Sprintf("```%s```", review.Content),
		}).
			SetFooter(&discordgo.MessageEmbedFooter{
				Text: fmt.Sprintf("👍 %d", review.LikeTotal),
			})
	}
	err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
		Embeds: []*discordgo.MessageEmbed{
			embed.MessageEmbed,
		},
		AllowedMentions: &discordgo.MessageAllowedMentions{},
		Flags:           discordgo.MessageFlagsEphemeral,
	}))
	if err != nil {
		log.Println("Error on sending embed")
	}
}

func look_reviewList(s *discordgo.Session, i *discordgo.InteractionCreate, subjectId string) {

}
