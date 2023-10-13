package reviewutil

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/utils/builder"
	db "github.com/y2hO0ol23/weiver/utils/database"
)

func Resend(s *discordgo.Session, i *discordgo.InteractionCreate, review *db.ReviewModel, comment string) bool {
	to, err := s.GuildMember(i.GuildID, review.ToID)
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return false
	}

	button_good := builder.Button().
		SetCustomID(fmt.Sprintf("like_review_%d", review.ID)).
		SetLable("👍").
		SetStyle(discordgo.SecondaryButton)

	button_bad := builder.Button().
		SetCustomID(fmt.Sprintf("hate_review_%d", review.ID)).
		SetLable("👎").
		SetStyle(discordgo.SecondaryButton)

	err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
		Embeds: []*discordgo.MessageEmbed{
			builder.Embed().
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
				SetTimeStamp(review.TimeStamp).
				MessageEmbed,
		},
		Components: []discordgo.MessageComponent{
			builder.ActionRow().AddComponents(button_good, button_bad).ActionsRow,
		},
		AllowedMentions: &discordgo.MessageAllowedMentions{},
	}))
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return false
	}

	// add msg data on db
	msg, err := s.InteractionResponse(i.Interaction)
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return false
	}
	db.UpdateMessageInfoByID(review.ID, i.GuildID, msg.ChannelID, msg.ID)

	// send dm to subject
	channel, err := s.UserChannelCreate(review.ToID)
	if channel != nil {
		s.ChannelMessageSendEmbeds(channel.ID, []*discordgo.MessageEmbed{
			builder.Embed().
				SetFields(&discordgo.MessageEmbedField{
					Name:  fmt.Sprintf("🔔 Your review has %s", comment),
					Value: fmt.Sprintf("➥ https://discord.com/channels/%s/%s/%s", i.GuildID, msg.ChannelID, msg.ID),
				}).
				MessageEmbed,
		})
	}

	return true
}
