package services

import (
	"fmt"
	"log"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	g "github.com/y2hO0ol23/weiver/api"
	db "github.com/y2hO0ol23/weiver/database"
)

func init() {
	app := g.App.Group("/reviews")

	app.Get("/state", func(c *fiber.Ctx) error {
		var (
			form struct {
				Count   int64   `json:"count"`
				Average float64 `json:"avg"`
			}
			err error
		)

		form.Count, err = db.GetReviewsCount()
		if err != nil {
			return fiber.NewError(500)
		}
		form.Average, err = db.GetReviewsScoreAvg()
		if err != nil {
			return fiber.NewError(500)
		}

		return c.JSON(form)
	})

	app.Get("/list", func(c *fiber.Ctx) error {
		type form struct {
			ID  int `json:"id"`
			URL string

			Score   int    `json:"score"`
			Title   string `json:"title"`
			Content string `json:"content"`

			Likes int64 `json:"likes"`

			TimeStamp time.Time `json:"timestamp"`

			Permission bool `json:"permission"`
		}

		// check params
		var (
			queries = c.Queries()
			from    int
			count   int
			user    string
			orderby string = "Like_Total desc, Time_Stamp asc"
			err     error
		)
		if v, ok := queries["from"]; ok {
			from, err = strconv.Atoi(v)
			if err != nil || from < 0 {
				return fiber.NewError(412, "'from' unvalued")
			}
		} else {
			return fiber.NewError(412, "'from' unvalued")
		}
		if v, ok := queries["count"]; ok {
			count, err = strconv.Atoi(v)
			if err != nil || count < 1 || 100 < count {
				return fiber.NewError(412, "'count' unvalued")
			}
		} else {
			return fiber.NewError(412, "'count' unvalued")
		}
		if v, ok := queries["orderby"]; ok {
			orderby = v
		}
		if v, ok := queries["user"]; ok {
			user = v
		}

		// get reviews
		res, err := db.GetReviews(from, count, orderby)
		if err != nil {
			return fiber.NewError(500)
		} else if res == nil {
			return fiber.NewError(404)
		}

		List := []form{}
		for _, e := range *res {
			var permission bool
			if user != "" {
				value, err := g.Session.UserChannelPermissions(user, e.ChannelID)
				if err != nil {
					log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
					permission = true
				} else {
					permission = (value & discordgo.PermissionViewChannel) != 0
				}
			} else {
				permission = true
			}
			List = append(List, form{
				ID:         e.ID,
				Title:      e.Title,
				Score:      e.Score,
				Content:    e.Content,
				Likes:      e.LikeTotal,
				TimeStamp:  e.TimeStamp,
				URL:        fmt.Sprintf("https://discord.com/channels/%v/%v/%v", e.GuildID, e.ChannelID, e.MessageID),
				Permission: permission,
			})
		}

		return c.JSON(List)
	})
}
