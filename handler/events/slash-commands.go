package events

import (
	"github.com/bwmarrin/discordgo"

	g "github.com/y2hO0ol23/weiver/handler"
)

func init() {
	g.EventList = append(g.EventList, func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		if exec, ok := g.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			exec(s, i)
		}
	})
}
