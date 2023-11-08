package handler

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type CMDForm struct {
	Data    *discordgo.ApplicationCommand
	Slash   func(s *discordgo.Session, i *discordgo.InteractionCreate)
	Message func(s *discordgo.Session, i *discordgo.InteractionCreate, locale discordgo.Locale, queries []string) string
}

var CMDList = []CMDForm{}
var CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}

var registeredCommands = []*discordgo.ApplicationCommand{}

func SetupSlashCommands(s *discordgo.Session) {
	// register slash commands
	for _, e := range CMDList {
		CommandHandlers[e.Data.Name] = e.Slash
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

func ParseOptionUser(s *discordgo.Session, guildID string, data string) *string {
	if !strings.HasPrefix(data, "<@") || !strings.HasSuffix(data, ">") {
		return nil
	}

	memberID := data[2 : len(data)-1]

	_, err := s.GuildMember(guildID, memberID)
	if err != nil {
		return nil
	}

	return &memberID
}
