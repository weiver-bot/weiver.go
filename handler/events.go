package handler

import (
	"github.com/bwmarrin/discordgo"
)

var EventList = make([]interface{}, 0)

func SetupEvents(s *discordgo.Session) {
	for _, e := range EventList {
		s.AddHandler(e)
	}
}
