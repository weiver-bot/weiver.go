package events

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	parse "github.com/y2hO0ol23/weiver/handler/events/button"
	"github.com/y2hO0ol23/weiver/utils/builder"
	db "github.com/y2hO0ol23/weiver/utils/database"
)

func init() {
	events = append(events, func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionMessageComponent {
			return
		}

		data := i.MessageComponentData()
		if data.ComponentType != discordgo.ButtonComponent {
			return
		}

		for name, handler := range db.ReviewButtonHandler {
			var review *db.ReviewModel

			if reviewID, ok := parse.Like.CustomID(data.CustomID, name); ok {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseDeferredMessageUpdate,
				})

				db.LoadUserByID(i.Interaction.Member.User.ID)
				review = handler(reviewID, i.Interaction.Member.User.ID)
			} else {
				continue
			}

			to, err := s.GuildMember(i.GuildID, review.ToID)
			if err != nil {
				log.Println(err)
				return
			}

			embed := builder.Embed().
				SetDescription(fmt.Sprintf("<@%s> → <@%s>", review.FromID, review.ToID)).
				SetFields(&discordgo.MessageEmbedField{
					Name:  fmt.Sprintf("📝 %s [%s%s]", review.Title, "★★★★★"[:review.Score*3], "☆☆☆☆☆"[review.Score*3:]),
					Value: fmt.Sprintf("```%s```", review.Content),
				}).
				SetFooter(&discordgo.MessageEmbedFooter{
					Text: fmt.Sprintf("👍 %d", review.LikeTotal),
				}).
				SetThumbnail(&discordgo.MessageEmbedThumbnail{
					URL: to.User.AvatarURL(""),
				}).
				SetTimeStamp(review.TimeStamp)

			_, err = s.ChannelMessageEditEmbeds(review.ChannelID, review.MessageID, []*discordgo.MessageEmbed{
				embed.MessageEmbed,
			})
			if err != nil {
				log.Println(err)
			}
		}
	})
}
