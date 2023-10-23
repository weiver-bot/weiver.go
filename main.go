package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	webapi "github.com/y2hO0ol23/weiver/api"
	"github.com/y2hO0ol23/weiver/handler"
	_ "github.com/y2hO0ol23/weiver/localization"
	_ "github.com/y2hO0ol23/weiver/utils/database"
	_ "github.com/y2hO0ol23/weiver/utils/env"
)

func main() {
	s, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Panicf("Error creating Discord session\n%v", err)
	}

	s.Identify.Intents = 0 |
		discordgo.IntentsGuilds |
		discordgo.IntentsGuildMessages |
		discordgo.IntentsGuildMembers

	handler.SetupEvents(s)

	err = s.Open()
	if err != nil {
		log.Panicf("Error opening connection\n%v", err)
	}
	defer s.Close()

	go webapi.Start(os.Getenv("API_PORT"), s)

	// need appID, so execute after session is open
	handler.SetupSlashCommands(s)
	if os.Getenv("REMOVE_CMD") != "" {
		defer handler.RemoveCommands(s)
	}

	log.Println("[*] End of settings")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
}
