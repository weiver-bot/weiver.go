package services

import (
	"os"

	"github.com/gofiber/fiber/v2"
	g "github.com/y2hO0ol23/weiver/api"
)

func init() {
	app := g.App.Group("/info")

	app.Get("/", func(c *fiber.Ctx) error {
		var form struct {
			Name          string `json:"name"`
			Discriminator string `json:"discriminator"`
			URL           struct {
				Invite    string `json:"invite"`
				Community string `json:"community"`
			}
		}
		s := g.Session

		form.Name = s.State.User.Username
		form.Discriminator = s.State.User.Discriminator
		form.URL.Invite = os.Getenv("INVITE_URL")
		form.URL.Community = os.Getenv("COMMUNITY_URL")

		return c.JSON(form)
	})
}
