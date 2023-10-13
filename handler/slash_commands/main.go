package slash_commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var err error

type form struct {
	data    *discordgo.ApplicationCommand
	execute func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var commands = []form{}
var registeredCommands = []*discordgo.ApplicationCommand{}

func Setup(s *discordgo.Session) {
	registeredCommands = make([]*discordgo.ApplicationCommand, len(commands))
	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		if exec, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			exec(s, i)
		}
	})

	for i, v := range commands {
		commandHandlers[v.data.Name] = v.execute
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v.data)
		if err != nil {
			RemoveCommands(s)
			log.Fatalf("Cannot create %q\n%v", v.data.Name, err)
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
