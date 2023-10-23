package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bwmarrin/discordgo"
	handlers "github.com/y2hO0ol23/weiver/api/handler/include"
	db "github.com/y2hO0ol23/weiver/utils/database"
)

func init() {
	type form struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Score     string `json:"score"`
		Content   string `json:"content"`
		Like      int64  `json:"like"`
		Timestamp string `json:"timestamp"`
		URL       string `json:"url"`
	}

	handlers.List = append(handlers.List, handlers.Form{
		Path: "/reviews",
		Execute: func(session *discordgo.Session) http.HandlerFunc {
			return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				//s := session

				switch r.Method {
				case http.MethodGet:
					var (
						reviews *[]db.ReviewModel
						count   = 5
						list    = []form{}
					)

					from, err := strconv.Atoi(r.URL.Query().Get("from"))
					if err != nil {
						from = 0
						count = 0
					}

					reviews, err = db.GetReviews(from, count)
					if err == nil {
						for _, e := range *reviews {
							list = append(list, form{
								ID:        e.ID,
								Title:     e.Title,
								Score:     fmt.Sprintf("%v%v", "★★★★★"[:e.Score*3], "☆☆☆☆☆"[e.Score*3:]),
								Content:   e.Content,
								Like:      e.LikeTotal,
								Timestamp: e.TimeStamp.Format("2006.01.02 15:04 Mon"),
								URL:       fmt.Sprintf("https://discord.com/channels/%v/%v/%v", e.GuildID, e.ChannelID, e.MessageID),
							})
						}
					}
					json.NewEncoder(rw).Encode(list)
				}
			})
		},
	})
}
