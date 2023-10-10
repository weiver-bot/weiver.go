package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	"github.com/y2hO0ol23/weiver/handler"
	"github.com/y2hO0ol23/weiver/handler/slash_commands"
	_ "github.com/y2hO0ol23/weiver/utils/prisma"
)

var (
	Token string
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file\n%v", err)
	}

	Token = os.Getenv("Token")
}

func main() {
	s, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatalf("Error creating Discord session\n%v", err)
	}

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err = s.Open()
	if err != nil {
		log.Fatalf("Error opening connection\n%v", err)
	}

	botintents := 0 |
		discordgo.IntentsGuilds |
		discordgo.IntentsGuildMessages |
		discordgo.IntentsGuildMembers

	s.Identify.Intents = discordgo.MakeIntent(botintents)

	handler.Setup(s)

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	slash_commands.Clean(s)
}
