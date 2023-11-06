package handle

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/database"
	"github.com/y2hO0ol23/weiver/localization"
	ReviewUtils "github.com/y2hO0ol23/weiver/utils/bot/review"
	"github.com/y2hO0ol23/weiver/utils/builder"
)

func Selected(
	s *discordgo.Session,
	iter *discordgo.InteractionCreate,
	i *discordgo.InteractionCreate,
	id int,
	timestamp string,
	locale discordgo.Locale,
) {
	// set embed by review on db
	embed := builder.Embed()
	review, err := db.LoadReivewByID(id)
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return
	}
	if review == nil || timestamp != fmt.Sprintf("%v", review.TimeStamp.Unix()) {
		// if review not exist or written time is not same with db data
		embed.SetDescription(fmt.Sprintf("❌ %v", localization.Load(locale, "$review.IsEdited")))
	} else {
		msg, err := s.ChannelMessage(review.ChannelID, review.MessageID)
		if err == nil {
			embed.SetDescription(fmt.Sprintf("https://discord.com/channels/%v/%v/%v", review.GuildID, review.ChannelID, review.MessageID)).
				SetFields(&discordgo.MessageEmbedField{
					Name:  fmt.Sprintf("📝 %v [%v%v]", review.Title, "★★★★★"[:review.Score*3], "☆☆☆☆☆"[review.Score*3:]),
					Value: fmt.Sprintf("```%v```", review.Content),
				}).
				SetThumbnail(&discordgo.MessageEmbedThumbnail{
					URL: msg.Embeds[0].Thumbnail.URL,
				})
		} else { // if message is not exist -> guild, channel or message is deleted
			_, err := s.GuildMember(review.GuildID, review.SubjectID)
			if err == nil {
				review, err = ReviewUtils.SendReview(s, iter, review)
				if err != nil {
					log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
					return
				}
				if review != nil {
					err := s.InteractionResponseDelete(i.Interaction)
					if err != nil {
						log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
					}
					ReviewUtils.ModifyReviewDM(s, review, locale)
				}
				return
			}
			// if this guild not have user that is this review's subject
			embed.SetDescription(fmt.Sprintf("https://discord.com/channels/%v/%v/%v", review.GuildID, review.ChannelID, review.MessageID)).
				SetFields(&discordgo.MessageEmbedField{
					Name:  fmt.Sprintf("🔒 %v [%v%v]", review.Title, "★★★★★"[:review.Score*3], "☆☆☆☆☆"[review.Score*3:]),
					Value: fmt.Sprintf("`%v`", localization.Load(locale, "$review.NoAuthor")),
				})
		}
	}

	// send embed
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
