package handler

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type CMDForm struct {
	Data    *discordgo.ApplicationCommand
	Execute func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var CMDList = []CMDForm{}
var EventList = make([]interface{}, 0)

func SetupEvents(s *discordgo.Session) {
	for _, e := range EventList {
		s.AddHandler(e)
	}
}

var registeredCommands = []*discordgo.ApplicationCommand{}

func SetupSlashCommands(s *discordgo.Session) {
	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		if exec, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			exec(s, i)
		}
	})

	for _, e := range CMDList {
		commandHandlers[e.Data.Name] = e.Execute
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", e.Data)
		if err != nil {
			log.Panicf("Error creating command: %q\n%v", e.Data.Name, err)
			RemoveCommands(s)
		}
		registeredCommands = append(registeredCommands, cmd)
	}
}

func RemoveCommands(s *discordgo.Session) {
	for _, e := range registeredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", e.ID)
		if err != nil {
			log.Println("Cannot delete slash command:", e.Name, e.ID)
		}
	}
}
