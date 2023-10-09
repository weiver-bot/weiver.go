package slash_commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type form struct {
	data    *discordgo.ApplicationCommand
	execute func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var commands = []form{}
var registeredCommands = []*discordgo.ApplicationCommand{}

func Setup(s *discordgo.Session) {
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}

	for i, v := range commands {
		commandHandlers[v.data.Name] = v.execute
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v.data)
		if err != nil {
			log.Fatalf("Cannot create '%v' command: %v", v.data.Name, err)
		}
		registeredCommands[i] = cmd
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		if execute, exist := commandHandlers[i.ApplicationCommandData().Name]; exist {
			execute(s, i)
		}
	})
}

func Clean(s *discordgo.Session) {
	for _, v := range registeredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			log.Fatalf("Cannot delete slash command %q: %v", v.Name, err)
		}
	}
}
