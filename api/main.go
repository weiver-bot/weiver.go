package api

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "github.com/y2hO0ol23/weiver/env"
)

var App = fiber.New()
var Session *discordgo.Session

func init() {
	App.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     os.Getenv("WEB_URL"),
	}))
}

func Start(s *discordgo.Session) {
	if port := os.Getenv("API_PORT"); port != "" {
		Session = s

		if err := App.Listen(":" + port); err != nil {
			log.Panicf("Error opening api server\n%v", err)
		}
	}
}
