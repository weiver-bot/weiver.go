package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	g "github.com/y2hO0ol23/weiver/api"
)

func init() {
	app := g.App.Group("/auth")

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Redirect(os.Getenv("DISCORD_OAUTH2_URL"))
	})

	app.Get("/logout", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:  "discord.refresh_token",
			Value: "",
		})
		return c.Redirect(os.Getenv("WEB_URL"))
	})

	app.Get("/callback", func(c *fiber.Ctx) error {
		code, ok := c.Queries()["code"]
		if !ok {
			return c.Redirect(os.Getenv("WEB_URL"))
		}

		// request to get token
		params := fmt.Sprintf(
			"client_id=%v&client_secret=%v&grant_type=%v&code=%v&redirect_uri=%v",
			os.Getenv("DISCORD_CLIENT_ID"),
			os.Getenv("DISCORD_CLIENT_SECRET"),
			"authorization_code",
			code,
			os.Getenv("DISCORD_REDIRECT_URI"),
		)

		req, err := http.NewRequest("POST", "https://discord.com/api/oauth2/token", bytes.NewBufferString(params))
		if err != nil {
			return fiber.NewError(404, err.Error())
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		r, err := http.DefaultClient.Do(req)
		if err != nil {
			return fiber.NewError(404, err.Error())
		}
		defer r.Body.Close()

		// parse token
		var tokens struct {
			Refresh string `json:"refresh_token"`
		}
		if err = json.NewDecoder(r.Body).Decode(&tokens); err != nil {
			return c.JSON(500)
		}

		// set token
		c.Cookie(&fiber.Cookie{
			Name:     "discord.refresh_token",
			Value:    tokens.Refresh,
			HTTPOnly: true,
			Secure:   true,
		})

		return c.Redirect(os.Getenv("WEB_URL"))
	})

	app.Post("/refresh", func(c *fiber.Ctx) error {
		refresh_token := c.Cookies("discord.refresh_token")

		req, err := http.NewRequest(
			"POST", "https://discord.com/api/oauth2/token",
			bytes.NewBufferString(
				fmt.Sprintf(
					"client_id=%v&client_secret=%v&grant_type=%v&refresh_token=%v",
					os.Getenv("DISCORD_CLIENT_ID"),
					os.Getenv("DISCORD_CLIENT_SECRET"),
					"refresh_token",
					refresh_token,
				),
			),
		)
		if err != nil {
			return fiber.NewError(404)
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		r, err := http.DefaultClient.Do(req)
		if err != nil {
			return fiber.NewError(404)
		}
		defer r.Body.Close()

		// parse token
		var tokens struct {
			Refresh string `json:"refresh_token"`
			Access  string `json:"access_token"`
		}
		if err = json.NewDecoder(r.Body).Decode(&tokens); err != nil {
			return c.JSON(500)
		}

		c.Cookie(&fiber.Cookie{
			Name:     "discord.refresh_token",
			Value:    tokens.Refresh,
			HTTPOnly: true,
			Secure:   true,
		})

		return c.JSON(fiber.Map{
			"access_token": tokens.Access,
		})
	})
}
