package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"

	"github.com/y2hO0ol23/weiver/handler"
	"github.com/y2hO0ol23/weiver/handler/slash_commands"
	_ "github.com/y2hO0ol23/weiver/utils/database"
	_ "github.com/y2hO0ol23/weiver/utils/env"
)

func main() {
	s, err := discordgo.New("Bot " + os.Getenv("Token"))
	if err != nil {
		log.Fatal("Error creating Discord session, ", err)
	}
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	s.Identify.Intents = 0 |
		discordgo.IntentsGuilds |
		discordgo.IntentsGuildMessages |
		discordgo.IntentsGuildMembers

	err = s.Open()
	if err != nil {
		log.Fatalf("Error opening connection\n%v", err)
	}
	defer s.Close()

	handler.Setup(s)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	slash_commands.Clean(s)
}
