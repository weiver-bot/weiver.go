package slashcommands

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/localization"
	"github.com/y2hO0ol23/weiver/utils/builder"
	db "github.com/y2hO0ol23/weiver/utils/database"
	reviewutil "github.com/y2hO0ol23/weiver/utils/review"
)

func init() {
	var DMPermission bool = false

	commands = append(commands, form{
		data: &discordgo.ApplicationCommand{
			Name:                     "look",
			Description:              "look_Description",
			NameLocalizations:        localization.LoadList("#look"),
			DescriptionLocalizations: localization.LoadList("#look.Description"),
			DMPermission:             &DMPermission,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:                     "info",
					Description:              "info_Description",
					NameLocalizations:        *localization.LoadList("#look.info"),
					DescriptionLocalizations: *localization.LoadList("#look.info.Description"),
					Type:                     discordgo.ApplicationCommandOptionSubCommand,
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
				{
					Name:                     "reviews",
					Description:              "reviews_Description",
					NameLocalizations:        *localization.LoadList("#look.reviews"),
					DescriptionLocalizations: *localization.LoadList("#look.reviews.Description"),
					Type:                     discordgo.ApplicationCommandOptionSubCommand,
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
			case "reviews":
				look_reviewList(s, i, options[0].Options[0].Value.(string))
			}
		},
	})
}

func look_info(s *discordgo.Session, i *discordgo.InteractionCreate, subjectID string) {
	locale := i.Locale

	subject, err := s.GuildMember(i.GuildID, subjectID)
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return
	}

	var avg float64 = 0.0
	count, err := db.GetUserReviewCount(subjectID)
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return
	}
	if count > 0 {
		avg, err = db.GetUserScore(subjectID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
	}
	countOutput := func() string {
		if count >= 100 {
			return "100+"
		} else if count >= 50 {
			return "50+"
		}
		return fmt.Sprintf("%d", count)
	}()

	embed := builder.Embed().
		SetDescription(fmt.Sprintf("<@%s> **⭐%.1f (%s)**", subjectID, avg, countOutput)).
		SetThumbnail(&discordgo.MessageEmbedThumbnail{
			URL: subject.AvatarURL(""),
		})

	if count == 0 {
		embed.SetFields(&discordgo.MessageEmbedField{
			Name:  fmt.Sprintf("📑 %s", localization.Load(locale, "#look.info.IsNone")),
			Value: "``` ```",
		})
	} else {
		review, err := db.GetReviewBest(subjectID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
		embed.
			SetFields(&discordgo.MessageEmbedField{
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
		return
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
	locale := i.Locale

	subject, err := s.GuildMember(i.GuildID, subjectID)
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return
	}

	reviews, err := db.GetReviewsByUserID(subjectID)
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return
	}
	if reviews == nil {
		message := builder.Message(&discordgo.InteractionResponseData{
			Content:         fmt.Sprintf("`%s`", localization.Load(locale, "#look.reviews.IsNone")),
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
	selectMenu := BuildSelectMenu(*reviews, locale, subject.User.Username, pageNow, pageCount)

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
		if data.CustomID != "reviews" {
			return
		}

		locale = iter.Locale

		value := data.Values[0]
		if strings.HasPrefix(value, "page/") { // move page
			s.AddHandlerOnce(handler)

			pageNow, err := strconv.Atoi(strings.Split(value, ":")[1])
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				return
			}

			selectMenu = BuildSelectMenu(*reviews, locale, subject.User.Username, pageNow, pageCount)

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
			review, err := db.LoadReivewByID(id)
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				return
			}
			if review == nil || data[1] != fmt.Sprintf("%d", review.TimeStamp.Unix()) {
				embed.SetDescription(fmt.Sprintf("❌ %s", localization.Load(locale, "$review.IsEdited")))
			} else {
				_, err := s.ChannelMessage(review.ChannelID, review.MessageID)
				if err == nil {
					embed.SetDescription(fmt.Sprintf("https://discord.com/channels/%s/%s/%s", review.GuildID, review.ChannelID, review.MessageID)).
						SetFields(&discordgo.MessageEmbedField{
							Name:  fmt.Sprintf("📝 %s [%s%s]", review.Title, "★★★★★"[:review.Score*3], "☆☆☆☆☆"[review.Score*3:]),
							Value: fmt.Sprintf("```%s```", review.Content),
						}).
						SetThumbnail(&discordgo.MessageEmbedThumbnail{
							URL: subject.AvatarURL(""),
						})
				} else {
					_, err := s.GuildMember(review.GuildID, review.FromID)
					if err == nil {
						review, err = reviewutil.Send(s, iter, review)
						if err != nil {
							log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
							return
						}
						if review != nil {
							err := s.InteractionResponseDelete(i.Interaction)
							if err != nil {
								log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
							}
							reviewutil.ModifyDM(s, review, locale)
						}
						return
					} else {
						embed.SetDescription(fmt.Sprintf("https://discord.com/channels/%s/%s/%s", review.GuildID, review.ChannelID, review.MessageID)).
							SetFields(&discordgo.MessageEmbedField{
								Name:  fmt.Sprintf("🔒 %s [%s%s]", review.Title, "★★★★★"[:review.Score*3], "☆☆☆☆☆"[review.Score*3:]),
								Value: fmt.Sprintf("`%s`", localization.Load(locale, "$review.NoAuthor")),
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

func BuildSelectMenu(reviews []db.ReviewModel, locale discordgo.Locale, subjectName string, pageNow int, pageCount int) *builder.SelectMenuStructure {
	pageBack := (pageNow+pageCount-2)%pageCount + 1
	pageNext := pageNow%pageCount + 1

	selectMenu := builder.SelectMenu().
		SetCustomID("reviews").
		SetPlaceholder(fmt.Sprintf(localization.Load(locale, "#look.reviews.menu.Title")+" (%d/%d)", subjectName, pageNow, pageCount))

	if pageCount > 1 {
		selectMenu.AddOptions(
			builder.SelectMenuOption().
				SetLabel("▲").
				SetValue(fmt.Sprintf("page/back:%d", pageBack)).
				SetDescription(fmt.Sprintf(localization.Load(locale, "#look.reviews.menu.Page"), pageBack)),
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
				SetDescription(fmt.Sprintf(localization.Load(locale, "#look.reviews.menu.Page"), pageNext)),
		)
	}

	return selectMenu
}
