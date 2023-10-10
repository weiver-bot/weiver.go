package events

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	events = append(events, func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionMessageComponent {
			return
		}
		/*
			return

			data := i.MessageComponentData()
			fromId, toId, ok := parseCustomID(data.CustomID)
			if ok {
				return
			}
			score, title, content, ok := parseModalComponents(data.Components)
			if ok {
				return
			}

			to, err := s.State.Member(i.GuildID, toId)
			if err != nil {
				return
			}

			// set db
			review := prisma.ModifyReviewByIds(fromId, toId, score, title, content)

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
			}*/
	})
}
