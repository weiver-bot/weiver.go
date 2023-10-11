package events

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	parse "github.com/y2hO0ol23/weiver/handler/events/modal"
	"github.com/y2hO0ol23/weiver/utils/builder"
	"github.com/y2hO0ol23/weiver/utils/prisma"
)

func init() {
	events = append(events, func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionModalSubmit {
			return
		}

		data := i.ModalSubmitData()
		fromId, toId, ok := parse.Review.CustomID(data.CustomID)
		if !ok {
			log.Println("Error on parse CustomID")
			return
		}
		score, title, content, ok := parse.Review.ModalComponents(data.Components)
		if !ok {
			log.Println("Error on parse ModalComponents")
			return
		}

		to, err := s.GuildMember(i.GuildID, toId)
		if err != nil {
			log.Println("Error on loadding user")
			return
		}

		// remove old reivew
		review := prisma.LoadReivewByIds(fromId, toId)
		if review != nil {
			if channelId, ok := review.ChannelID(); ok {
				if messageId, ok := review.MessageID(); ok {
					_, err := s.ChannelMessage(channelId, messageId)
					if err == nil {
						s.ChannelMessageDelete(channelId, messageId)
					}
				}
			}
		}

		// set db
		review = prisma.ModifyReviewByIds(fromId, toId, score, title, content)

		button_good := builder.Button().
			SetCustomId(fmt.Sprintf("like_review_%d", review.ID)).
			SetLable("👍").
			SetStyle(discordgo.SecondaryButton)

		button_bad := builder.Button().
			SetCustomId(fmt.Sprintf("hate_review_%d", review.ID)).
			SetLable("👎").
			SetStyle(discordgo.SecondaryButton)

		err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				builder.Embed().
					SetDescription(fmt.Sprintf("<@%s> → <@%s>", review.FromID, review.ToID)).
					SetFields(&discordgo.MessageEmbedField{
						Name:  fmt.Sprintf("📝 %s [%s%s]", title, "★★★★★"[:score*3], "☆☆☆☆☆"[score*3:]),
						Value: fmt.Sprintf("```%s```", content),
					}).
					SetFooter(&discordgo.MessageEmbedFooter{
						Text: "👍 0",
					}).
					SetThumbnail(&discordgo.MessageEmbedThumbnail{
						URL: to.User.AvatarURL(""),
					}).
					MessageEmbed,
			},
			Components: []discordgo.MessageComponent{
				builder.ActionRow().SetComponents(button_good, button_bad).ActionsRow,
			},
			AllowedMentions: &discordgo.MessageAllowedMentions{},
		}))
		if err != nil {
			log.Printf("Error on sending embed\n")
			return
		}

		// add msg data on db
		msg, err := s.InteractionResponse(i.Interaction)
		if err != nil {
			log.Printf("Error on loading interaction\n")
			return
		}
		prisma.UpdateIdsById(review.ID, i.GuildID, msg.ChannelID, msg.ID)

		// send dm to subject
		channel, err := s.UserChannelCreate(toId)
		if err != nil {
			log.Printf("Error on sending DM, to %v\n", channel)
		}
		s.ChannelMessageSendEmbeds(channel.ID, []*discordgo.MessageEmbed{
			builder.Embed().
				SetFields(&discordgo.MessageEmbedField{
					Name:  "🔔 Your review has written",
					Value: fmt.Sprintf("➥ https://discord.com/channels/%s/%s/%s", i.GuildID, msg.ChannelID, msg.ID),
				}).
				MessageEmbed,
		})
	})
}
