package handle

import (
	"log"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/utils/builder"
)

func MovePage(
	s *discordgo.Session,
	iter *discordgo.InteractionCreate,
	i *discordgo.InteractionCreate,
	selectMenu *builder.SelectMenuStructure,
) {
	// set select menu by page
	// send select menu
	s.InteractionRespond(iter.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})
	_, err := s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Components: &[]discordgo.MessageComponent{
			builder.ActionRow().AddComponents(selectMenu).ActionsRow,
		},
	})
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return
	}
}
