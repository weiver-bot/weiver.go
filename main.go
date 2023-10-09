package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	"github.com/y2hO0ol23/weiver/handler"
	"github.com/y2hO0ol23/weiver/handler/slash_commands"
)

var (
	Token string
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Token = os.Getenv("Token")
}

func main() {
	s, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatalf("error creating Discord session, %v", err)
	}

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err = s.Open()
	if err != nil {
		log.Fatalf("error opening connection, %v", err)
	}

	s.Identify.Intents = discordgo.IntentsGuildMessages

	handler.Setup(s)

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	slash_commands.Clean(s)
}
