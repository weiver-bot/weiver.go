package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/handler/events"
	"github.com/y2hO0ol23/weiver/handler/slash_commands"
)

func Setup(s *discordgo.Session) {
	events.Setup(s)
	slash_commands.Setup(s)
}
