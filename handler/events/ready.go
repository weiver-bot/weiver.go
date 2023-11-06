package events

import (
	"log"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/database"
	BotUtils "github.com/y2hO0ol23/weiver/utils/bot"

	g "github.com/y2hO0ol23/weiver/handler"
)

func init() {
	g.EventList = append(g.EventList, func(s *discordgo.Session, r *discordgo.Ready) {
		guildsDB, err := db.GetGuildInProgress()
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
		if guildsDB != nil {
			for _, guildDB := range *guildsDB {
				db.EndOFGuildProgress(guildDB.ID)
			}
		}

		err = BotUtils.UpdateStatus(s)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		}

		log.Printf("Logged in as %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
}
