package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/y2hO0ol23/weiver/handler/events"
	"github.com/y2hO0ol23/weiver/handler/slash_commands"
	_ "github.com/y2hO0ol23/weiver/utils/database"
	_ "github.com/y2hO0ol23/weiver/utils/env"
)

func main() {
	s, err := discordgo.New("Bot " + os.Getenv("Token"))
	if err != nil {
		log.Fatalf("Error creating Discord session\n%v", err)
	}

	s.Identify.Intents = 0 |
		discordgo.IntentsGuilds |
		discordgo.IntentsGuildMessages |
		discordgo.IntentsGuildMembers

	events.Setup(s)

	err = s.Open()
	if err != nil {
		log.Fatalf("Error opening connection\n%v", err)
	}
	defer s.Close()

	// need appID, so execute after session is open
	slash_commands.Setup(s)
	defer slash_commands.RemoveCommands(s)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
}
