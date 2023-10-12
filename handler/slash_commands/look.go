package slash_commands

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/handler/events"
	"github.com/y2hO0ol23/weiver/utils/builder"
	db "github.com/y2hO0ol23/weiver/utils/database"
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

func look_info(s *discordgo.Session, i *discordgo.InteractionCreate, subjectID string) {
	subject, err := s.GuildMember(i.GuildID, subjectID)
	if err != nil {
		log.Println(err)
		return
	}

	average, count := db.GetUserScoreAverage(subjectID)
	countOutput := func() string {
		if count >= 100 {
			return "100+"
		} else if count >= 50 {
			return "50+"
		}
		return fmt.Sprintf("%d", count)
	}()

	embed := builder.Embed().
		SetDescription(fmt.Sprintf("<@%s> **⭐%.1f (%s)**", subjectID, average, countOutput)).
		SetThumbnail(&discordgo.MessageEmbedThumbnail{
			URL: subject.User.AvatarURL(""),
		})

	if count == 0 {
		embed.SetFields(&discordgo.MessageEmbedField{
			Name:  "📑 No reviews",
			Value: "``` ```",
		})
	} else {
		review := db.GetReviewBest(subjectID)
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
		log.Println(err)
	}
}

var pageRow int

func init() {
	pageRow, err = strconv.Atoi(os.Getenv("PAGE_ROW"))
	if err != nil || pageRow < 1 || 25 < pageRow {
		pageRow = 10
	}
}

func look_reviewList(s *discordgo.Session, i *discordgo.InteractionCreate, subjectID string) {
	subject, err := s.GuildMember(i.GuildID, subjectID)
	if err != nil {
		log.Println(err)
		return
	}

	reviews := db.GetReviewsByUserID(subjectID)
	if reviews == nil {
		message := builder.Message(&discordgo.InteractionResponseData{
			Content: "`No review exists`",
			Flags:   discordgo.MessageFlagsEphemeral,
		})
		err = s.InteractionRespond(i.Interaction, message)
		if err != nil {
			log.Println(err)
		}
		return
	}

	pageNow := 1
	pageCount := (len(*reviews)-1)/pageRow + 1
	selectMenu := BuildSelectMenu(*reviews, subject.User.Username, pageNow, pageCount)

	err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
		Components: []discordgo.MessageComponent{
			builder.ActionRow().AddComponents(selectMenu).ActionsRow,
		},
		AllowedMentions: &discordgo.MessageAllowedMentions{},
		Flags:           discordgo.MessageFlagsEphemeral,
	}))
	if err != nil {
		log.Println(err)
	}

	msg, err := s.InteractionResponse(i.Interaction)
	if err != nil {
		log.Println(err)
	}

	var handler func(s *discordgo.Session, iter *discordgo.InteractionCreate)
	handler = func(s *discordgo.Session, iter *discordgo.InteractionCreate) {
		if iter.Type != discordgo.InteractionMessageComponent || iter.Interaction.Message.ID != msg.ID {
			s.AddHandlerOnce(handler)
			return
		}

		data := iter.MessageComponentData()
		if data.CustomID != "review-list" {
			s.AddHandlerOnce(handler)
			return
		}

		value := data.Values[0]
		if strings.HasPrefix(value, "review") {
			id, err := strconv.Atoi(value[7:])
			if err != nil {
				log.Println(err)
				return
			}

			embed := builder.Embed()
			review := db.LoadReivewByID(id)
			if review == nil {
				embed.SetFields(&discordgo.MessageEmbedField{
					Name: "❌ This review has been edited",
				})
			} else {
				_, err := s.ChannelMessage(review.ChannelID, review.MessageID)
				if err == nil {
					embed.SetDescription(fmt.Sprintf("https://discord.com/channels/%s/%s/%s", review.GuildID, review.ChannelID, review.MessageID)).
						SetFields(&discordgo.MessageEmbedField{
							Name:  fmt.Sprintf("📝 %s [%s%s]", review.Title, "★★★★★"[:review.Score*3], "☆☆☆☆☆"[review.Score*3:]),
							Value: fmt.Sprintf("```%s```", review.Content),
						}).
						SetThumbnail(&discordgo.MessageEmbedThumbnail{
							URL: subject.User.AvatarURL(""),
						})
				} else {
					_, err := s.GuildMember(review.GuildID, review.FromID)
					if err == nil {
						if events.ResendReview(s, iter, review, "resend") {
							if s.InteractionResponseDelete(i.Interaction) != nil {
								log.Println(err)
							}
						}
						return
					} else {
						embed.SetDescription(fmt.Sprintf("https://discord.com/channels/%s/%s/%s", review.GuildID, review.ChannelID, review.MessageID)).
							SetFields(&discordgo.MessageEmbedField{
								Name:  fmt.Sprintf("🔒 %s [%s%s]", review.Title, "★★★★★"[:review.Score*3], "☆☆☆☆☆"[review.Score*3:]),
								Value: "`Review has removed and author not in this server`",
							})
					}
				}
			}

			s.InteractionRespond(iter.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseDeferredMessageUpdate,
			})
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{
					embed.MessageEmbed,
				},
				Components: &[]discordgo.MessageComponent{},
			})
		} else if strings.HasPrefix(value, "page/") {
			s.AddHandlerOnce(handler)

			pageNow, err := strconv.Atoi(strings.Split(value, ":")[1])
			if err != nil {
				log.Println(err)
				return
			}

			selectMenu = BuildSelectMenu(*reviews, subject.User.Username, pageNow, pageCount)

			s.InteractionRespond(iter.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseDeferredMessageUpdate,
			})
			_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Components: &[]discordgo.MessageComponent{
					builder.ActionRow().AddComponents(selectMenu).ActionsRow,
				},
			})
			if err != nil {
				log.Println(err)
			}
		}
	}
	s.AddHandlerOnce(handler)
}

func BuildSelectMenu(reviews []db.ReviewModel, subjectName string, pageNow int, pageCount int) *builder.SelectMenuStructure {
	pageBack := (pageNow+pageCount-2)%pageCount + 1
	pageNext := pageNow%pageCount + 1

	selectMenu := builder.SelectMenu().
		SetCustomID("review-list").
		SetPlaceholder(fmt.Sprintf("Reviews on %s (%d/%d)", subjectName, pageNow, pageCount))

	if pageCount > 1 {
		selectMenu.AddOptions(
			builder.SelectMenuOption().
				SetLabel("▲").
				SetValue(fmt.Sprintf("page/back:%d", pageBack)).
				SetDescription(fmt.Sprintf("page %d", pageBack)),
		)
	}

	for i := (pageNow - 1) * pageRow; i < pageNow*pageRow; i++ {
		if i >= len(reviews) {
			break
		}

		review := reviews[i]
		selectMenu.AddOptions(
			builder.SelectMenuOption().
				SetLabel(fmt.Sprintf("%s 〔%s%s〕", review.Title, "★★★★★"[:review.Score*3], "☆☆☆☆☆"[review.Score*3:])).
				SetDescription(fmt.Sprintf("👍 %d", review.LikeTotal)).
				SetValue(fmt.Sprintf("review:%d", review.ID)),
		)
	}

	if pageCount > 1 {
		selectMenu.AddOptions(
			builder.SelectMenuOption().
				SetLabel("▼").
				SetValue(fmt.Sprintf("page/next:%d", pageNext)).
				SetDescription(fmt.Sprintf("page %d", pageNext)),
		)
	}

	return selectMenu
}
