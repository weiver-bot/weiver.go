package handler

import (
	"log"

	"github.com/bwmarrin/discordgo"
	events "github.com/y2hO0ol23/weiver/handler/events/include"
	slashcommands "github.com/y2hO0ol23/weiver/handler/slash-commands/include"

	_ "github.com/y2hO0ol23/weiver/handler/events"
	_ "github.com/y2hO0ol23/weiver/handler/slash-commands"
)

func SetupEvents(s *discordgo.Session) {
	for _, v := range events.List {
		s.AddHandler(v)
	}
}

var registeredCommands = []*discordgo.ApplicationCommand{}

func SetupSlashCommands(s *discordgo.Session) {
	registeredCommands = make([]*discordgo.ApplicationCommand, len(slashcommands.List))
	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		if exec, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			exec(s, i)
		}
	})

	for i, v := range slashcommands.List {
		commandHandlers[v.Data.Name] = v.Execute
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v.Data)
		if err != nil {
			log.Panicf("Cannot create %q\n%v", v.Data.Name, err)
			RemoveCommands(s)
		}
		registeredCommands[i] = cmd
	}
}

func RemoveCommands(s *discordgo.Session) {
	for _, v := range registeredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			log.Println("Cannot delete slash command:", v.Name, v.ID)
		}
	}
}
