package include

import (
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type Form struct {
	Path    string
	Handler http.HandlerFunc
}

var List = []Form{}
var Session *discordgo.Session
