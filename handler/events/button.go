package events

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	parse "github.com/y2hO0ol23/weiver/handler/events/button"
	db "github.com/y2hO0ol23/weiver/utils/database"
	reviewutil "github.com/y2hO0ol23/weiver/utils/review"
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
				review, err = handler(reviewID, i.Interaction.Member.User.ID)
				if err != nil {
					log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
					return
				}
			} else {
				continue
			}

			to, err := s.GuildMember(i.GuildID, review.ToID)
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				return
			}

			embed := reviewutil.EmbedMost(review, to.AvatarURL("")).
				SetDescription(fmt.Sprintf("<@%s> → <@%s>", review.FromID, review.ToID))

			_, err = s.ChannelMessageEditEmbeds(review.ChannelID, review.MessageID, []*discordgo.MessageEmbed{
				embed.MessageEmbed,
			})
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			}
		}
	})
}
