package events

import (
	"github.com/bwmarrin/discordgo"
)

var (
	events = make([]interface{}, 0)
	err    error
)

func Setup(s *discordgo.Session) {
	for _, v := range events {
		s.AddHandler(v)
	}
}
