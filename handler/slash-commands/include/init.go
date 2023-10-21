package include

import (
	"github.com/bwmarrin/discordgo"
)

type Form struct {
	Data    *discordgo.ApplicationCommand
	Execute func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var List = []Form{}
