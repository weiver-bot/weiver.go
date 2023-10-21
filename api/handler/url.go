package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	handlers "github.com/y2hO0ol23/weiver/api/handler/include"
)

func init() {
	type form struct {
		Invite    string `json:"invite"`
		Community string `json:"community"`
	}

	handlers.List = append(handlers.List, handlers.Form{
		Path: "/url",
		Execute: func(session *discordgo.Session) http.HandlerFunc {
			return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case http.MethodGet:
					json.NewEncoder(rw).Encode(form{
						Invite:    os.Getenv("INVITE_URL"),
						Community: os.Getenv("COMMUNITY_URL"),
					})
				}
			})
		},
	})
}
