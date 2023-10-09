package events

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func init() {
	events = append(events, func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionModalSubmit {
			return
		}

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Thank you for taking your time to fill this survey",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		if err != nil {
			log.Fatalf("%v", err)
		}
		data := i.ModalSubmitData()

		if !strings.HasPrefix(data.CustomID, "modals_survey") {
			return
		}

		userid := strings.Split(data.CustomID, "_")[2]
		_, err = s.ChannelMessageSend(i.ChannelID, fmt.Sprintf(
			"Feedback received. From <@%s>\n\n**Opinion**:\n%s\n\n**Suggestions**:\n%s",
			userid,
			data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value,
			data.Components[1].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value,
		))
		if err != nil {
			log.Fatalf("%v", err)
		}
	})
}
