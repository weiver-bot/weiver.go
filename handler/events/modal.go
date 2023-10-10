package events

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/utils/builder"
)

func init() {
	events = append(events, func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionModalSubmit {
			return
		}

		data := i.ModalSubmitData()
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

		embed := builder.Message(&discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Description: "<@" + fromId + "> → <@" + toId + ">",
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:  "📝 " + title + " [" + "★★★★★"[:score*3] + "☆☆☆☆☆]"[score*3:],
							Value: "```" + content + "```",
						},
					},
					Footer: &discordgo.MessageEmbedFooter{
						Text: "👍 0",
					},
					Thumbnail: &discordgo.MessageEmbedThumbnail{
						URL: to.User.AvatarURL(""),
					},
				},
			},
			AllowedMentions: &discordgo.MessageAllowedMentions{},
		})
		err = s.InteractionRespond(i.Interaction, embed)
		if err != nil {
			panic(err)
		}
	})
}

func parseCustomID(value string) (string, string, bool) {
	if strings.HasPrefix(value, "review") {
		data := strings.Split(value, "#")
		if len(data) == 3 {
			return data[1], data[2], false
		}
	}
	return "", "", true
}

func parseModalComponents(components []discordgo.MessageComponent) (int, string, string, bool) {
	if len(components) == 3 {
		score := func(value string) int {
			count := strings.Count(value, "★")
			if count == 0 {
				return 1
			}
			return count
		}(components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value)

		title := components[1].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
		content := components[2].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value

		return score, title, content, false
	}
	return 0, "", "", true
}
