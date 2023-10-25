package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
		Path: "/reviews/list",
		Handler: func(rw http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				var (
					reviews *[]db.ReviewModel
					list    = []form{}
					count   = 0
					order   = "Like_Total desc, Time_Stamp asc"
				)
				from, err := strconv.Atoi(r.URL.Query().Get("from"))
				if err == nil {
					count, err = strconv.Atoi(r.URL.Query().Get("count"))
					if err != nil || count < 0 || 100 < count {
						count = 0
					}
				}
				if v := r.URL.Query().Get("order"); v != "" {
					order = v
				}

				reviews, err = db.GetReviews(from, count, order)
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
		},
	})
}

func init() {
	type form struct {
		Count int64 `json:"count"`
	}

	handlers.List = append(handlers.List, handlers.Form{
		Path: "/reviews/count",
		Handler: func(rw http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				count, err := db.GetReviewsCount()
				if err != nil {
					count = 0
				}
				json.NewEncoder(rw).Encode(form{Count: count})
			}
		},
	})
}
