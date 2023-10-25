package webapi

import (
	"log"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	handlers "github.com/y2hO0ol23/weiver/api/handler/include"

	_ "github.com/y2hO0ol23/weiver/api/handler"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Access-Control-Allow-Credentials", "true")
		rw.Header().Add("Access-Control-Allow-Origin", os.Getenv("WEB_URL"))
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func Start(port string, s *discordgo.Session) {
	if port == "" {
		return
	}

	handlers.Session = s

	mux := http.NewServeMux()
	for _, v := range handlers.List {
		mux.Handle(v.Path, Middleware(v.Handler))
	}

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Panicf("Error opening api server\n%v", err)
	}
}
