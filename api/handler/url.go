package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	handlers "github.com/y2hO0ol23/weiver/api/handler/include"
)

func init() {
	type form struct {
		Invite    string `json:"invite"`
		Community string `json:"community"`
	}

	handlers.List = append(handlers.List, handlers.Form{
		Path: "/url",
		Handler: func(rw http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				json.NewEncoder(rw).Encode(form{
					Invite:    os.Getenv("INVITE_URL"),
					Community: os.Getenv("COMMUNITY_URL"),
				})
			}
		},
	})
}
