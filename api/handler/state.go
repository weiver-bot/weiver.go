package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	handlers "github.com/y2hO0ol23/weiver/api/handler/include"
	botutil "github.com/y2hO0ol23/weiver/utils/bot"
)

func init() {
	type form struct {
		Name  string `json:"name"`
		ID    string `json:"id"`
		State string `json:"state"`
	}

	handlers.List = append(handlers.List, handlers.Form{
		Path: "/state",
		Execute: func(session *discordgo.Session) http.HandlerFunc {
			return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				s := session

				switch r.Method {
				case http.MethodGet:
					state, err := botutil.StateText()
					if err != nil {
						log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
						state = fmt.Sprintf("Hello, I am %v", s.State.User.Username)
					}
					json.NewEncoder(rw).Encode(form{
						Name:  s.State.User.Username,
						ID:    fmt.Sprintf("%v#%v", s.State.User.Username, s.State.User.Discriminator),
						State: state,
					})
				}
			})
		},
	})
}
