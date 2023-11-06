package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	g "github.com/y2hO0ol23/weiver/api"
	db "github.com/y2hO0ol23/weiver/database"
)

func init() {
	app := g.App.Group("/reviews")

	type User struct {
		ID string `json:"id"`
	}

	LoadUser := func(auth []string) *User {
		var user = &User{}

		req, err := http.NewRequest("GET", "https://discord.com/api/users/@me", nil)
		if err != nil {
			return user
		}

		req.Header.Set("Authorization", auth[0])

		r, err := http.DefaultClient.Do(req)
		if err != nil {
			return user
		}
		defer r.Body.Close()

		json.NewDecoder(r.Body).Decode(user)
		return user
	}

	app.Get("/state", func(c *fiber.Ctx) error {
		var (
			relate string
			err    error
		)

		Count, err := db.GetReviewsCount(relate)
		if err != nil {
			return fiber.NewError(500)
		}
		Average, err := db.GetReviewsScoreAvg()
		if err != nil {
			return fiber.NewError(500)
		}

		return c.JSON(fiber.Map{
			"count": Count,
			"avg":   Average,
		})
	})

	app.Get("/state/count", func(c *fiber.Ctx) error {
		var (
			relate string
			err    error
		)

		if auth, ok := c.GetReqHeaders()["Authorization"]; ok {
			relate = LoadUser(auth).ID
		}

		Count, err := db.GetReviewsCount(relate)
		if err != nil {
			return fiber.NewError(500)
		}

		return c.JSON(fiber.Map{
			"count": Count,
		})
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
		}

		var (
			queries = c.Queries()
			from    int
			count   int
			orderby string = "Like_Total desc, Time_Stamp asc"
			relate  string = ""
			err     error
		)
		// check params
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

		// check authorization -> get user id
		if auth, ok := c.GetReqHeaders()["Authorization"]; ok {
			relate = LoadUser(auth).ID
		}

		// get reviews
		res, err := db.GetReviews(from, count, orderby, relate)
		if err != nil {
			return fiber.NewError(500)
		} else if res == nil {
			return fiber.NewError(404)
		}

		List := []form{}
		for _, e := range *res {
			List = append(List, form{
				ID:        e.ID,
				Title:     e.Title,
				Score:     e.Score,
				Content:   e.Content,
				Likes:     e.LikeTotal,
				TimeStamp: e.TimeStamp,
				URL:       fmt.Sprintf("https://discord.com/channels/%v/%v/%v", e.GuildID, e.ChannelID, e.MessageID),
			})
		}

		return c.JSON(List)
	})
}
