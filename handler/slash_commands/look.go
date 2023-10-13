package slash_commands

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/utils/builder"
	db "github.com/y2hO0ol23/weiver/utils/database"
	reviewutil "github.com/y2hO0ol23/weiver/utils/review"
)

func init() {
	var DMPermission bool = false

	commands = append(commands, form{
		data: &discordgo.ApplicationCommand{
			Name:         "look",
			Description:  "Look about things",
			DMPermission: &DMPermission,
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
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
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
		Flags:           discordgo.MessageFlagsEphemeral,
		AllowedMentions: &discordgo.MessageAllowedMentions{},
	}))
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
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
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return
	}

	reviews := db.GetReviewsByUserID(subjectID)
	if reviews == nil {
		message := builder.Message(&discordgo.InteractionResponseData{
			Content:         "`No review exists`",
			Flags:           discordgo.MessageFlagsEphemeral,
			AllowedMentions: &discordgo.MessageAllowedMentions{},
		})
		err = s.InteractionRespond(i.Interaction, message)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
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
		Flags:           discordgo.MessageFlagsEphemeral,
		AllowedMentions: &discordgo.MessageAllowedMentions{},
	}))
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
	}

	msg, err := s.InteractionResponse(i.Interaction)
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
	}

	// make handler because bot can not find message that has ephemeral flag
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
		if data.CustomID != "review-list" {
			return
		}

		value := data.Values[0]
		if strings.HasPrefix(value, "page/") { // move page
			s.AddHandlerOnce(handler)

			pageNow, err := strconv.Atoi(strings.Split(value, ":")[1])
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
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
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			}
		} else if strings.HasPrefix(value, "review") { // show page link
			data := strings.Split(value[7:], "#")
			id, err := strconv.Atoi(data[0])
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				return
			}

			embed := builder.Embed()
			review := db.LoadReivewByID(id)
			if review == nil || data[1] != fmt.Sprintf("%d", review.TimeStamp.Unix()) {
				embed.SetDescription("❌ This review has been edited")
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
						if review = reviewutil.Resend(s, iter, review); review != nil {
							err := s.InteractionResponseDelete(i.Interaction)
							if err != nil {
								log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
							}
							reviewutil.AlertByDM(s, review)
						}
						return
					} else {
						embed.SetDescription(fmt.Sprintf("https://discord.com/channels/%s/%s/%s", review.GuildID, review.ChannelID, review.MessageID)).
							SetFields(&discordgo.MessageEmbedField{
								Name:  fmt.Sprintf("🔒 %s [%s%s]", review.Title, "★★★★★"[:review.Score*3], "☆☆☆☆☆"[review.Score*3:]),
								Value: "`Review has removed but author not in this server`",
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
				SetValue(fmt.Sprintf("review:%d#%d", review.ID, review.TimeStamp.Unix())),
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
