package events

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/db"
	parse "github.com/y2hO0ol23/weiver/handler/events/button"
	"github.com/y2hO0ol23/weiver/utils/builder"
	"github.com/y2hO0ol23/weiver/utils/prisma"
)

func init() {
	events = append(events, func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionMessageComponent {
			return
		}
		data := i.MessageComponentData()

		for key, _ := range prisma.ReviewActionHandler {
			func(key string) {
				var review *db.ReviewModel

				if reviewId, ok := parse.Like.CustomID(data.CustomID, key); ok {
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseDeferredMessageUpdate,
					})

					prisma.LoadUserById(i.Interaction.Member.User.ID)
					review = prisma.ReviewAction(reviewId, prisma.ReviewActionHandler[key](i.Interaction.Member.User.ID))
				} else {
					return
				}

				to, err := s.GuildMember(i.GuildID, review.ToID)
				if err != nil {
					log.Println("Error on loadding member")
					return
				}

				if channelId, messageId, ok := func() (string, string, bool) {
					if channelId, ok := review.ChannelID(); ok {
						if messageId, ok := review.MessageID(); ok {
							_, err := s.ChannelMessage(channelId, messageId)
							if err == nil {
								return channelId, messageId, true
							}
						}
					}
					return "", "", false
				}(); ok {
					embed := builder.Embed().
						SetDescription(fmt.Sprintf("<@%s> → <@%s>", review.FromID, review.ToID)).
						SetField(&discordgo.MessageEmbedField{
							Name:  fmt.Sprintf("📝 %s [%s%s]", review.Title, "★★★★★"[:review.Score*3], "☆☆☆☆☆"[review.Score*3:]),
							Value: fmt.Sprintf("```%s```", review.Content),
						}).
						SetFooter(&discordgo.MessageEmbedFooter{
							Text: "👍 " + func(like_total int) string {
								if like_total >= 100 {
									return "100+"
								}
								if like_total >= 50 {
									return "50+"
								}
								return fmt.Sprintf("%d", like_total)
							}(review.LikeTotal),
						}).
						SetThumbnail(&discordgo.MessageEmbedThumbnail{
							URL: to.User.AvatarURL(""),
						})

					s.ChannelMessageEditEmbeds(channelId, messageId, []*discordgo.MessageEmbed{
						embed.MessageEmbed,
					})
				} else {
					log.Println("Error on editing Embeds")
					return
				}
			}(key)
		}
	})
}
