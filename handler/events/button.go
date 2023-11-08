package events

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/database"
	this "github.com/y2hO0ol23/weiver/handler/events/button"
	ReviewUtils "github.com/y2hO0ol23/weiver/utils/bot/review"

	g "github.com/y2hO0ol23/weiver/handler"
)

func init() {
	var err error

	g.EventList = append(g.EventList, func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionMessageComponent {
			return
		}

		data := i.MessageComponentData()
		if data.ComponentType != discordgo.ButtonComponent {
			return
		}

		for name, handler := range db.ReviewButtonHandler {
			var review *db.ReviewModel

			if reviewID, ok := this.Parse.CustomID(data.CustomID, name); ok {
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

			subject, err := s.GuildMember(i.GuildID, review.SubjectID)
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				return
			}

			embed := ReviewUtils.BaseEmbedWithFooter(review, subject.AvatarURL("")).
				SetDescription(fmt.Sprintf("<@%v> → <@%v>", review.AuthorID, review.SubjectID))

			_, err = s.ChannelMessageEditEmbeds(review.ChannelID, review.MessageID, []*discordgo.MessageEmbed{
				embed.MessageEmbed,
			})
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			}
		}
	})
}
