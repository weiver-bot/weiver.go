package events

import (
	"fmt"

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
			return
		}
		score, title, content, ok := parse.Review.ModalComponents(data.Components)
		if !ok {
			return
		}

		to, err := s.State.Member(i.GuildID, toId)
		if err != nil {
			return
		}

		// remove old reivew
		review := prisma.LoadReivewByIds(fromId, toId)
		if channelId, ok := review.ChannelID(); ok {
			if messageId, ok := review.MessageID(); ok {
				s.ChannelMessageDelete(channelId, messageId)
			}
		}

		// set db
		review = prisma.ModifyReviewByIds(fromId, toId, score, title, content)

		// set embed
		embed := builder.Embed().
			SetDescription("<@" + fromId + "> → <@" + toId + ">").
			SetField(&discordgo.MessageEmbedField{
				Name:  fmt.Sprintf("📝 %s [%s%s]", title, "★★★★★"[:score*3], "☆☆☆☆☆"[score*3:]),
				Value: fmt.Sprintf("```%s```", content),
			}).
			SetFooter(&discordgo.MessageEmbedFooter{
				Text: "👍 0",
			}).
			SetThumbnail(&discordgo.MessageEmbedThumbnail{
				URL: to.User.AvatarURL(""),
			})

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
				embed.MessageEmbed,
			},
			Components: []discordgo.MessageComponent{
				builder.ActionRow().SetComponents(button_good, button_bad).ActionsRow,
			},
			AllowedMentions: &discordgo.MessageAllowedMentions{},
		}))
		if err != nil {
			panic(err)
		}

		// add msg data on db
		msg, err := s.InteractionResponse(i.Interaction)
		if err != nil {
			panic(err)
		}
		prisma.UpdateIdsById(review.ID, i.GuildID, msg.ChannelID, msg.ID)
	})
}
