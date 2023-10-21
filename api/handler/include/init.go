package include

import (
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type Form struct {
	Path    string
	Execute func(s *discordgo.Session) http.HandlerFunc
}

var List = []Form{}
