package events

import (
	"log"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	events "github.com/y2hO0ol23/weiver/handler/events/include"
	botutil "github.com/y2hO0ol23/weiver/utils/bot"
	db "github.com/y2hO0ol23/weiver/utils/database"
)

func init() {
	events.List = append(events.List, func(s *discordgo.Session, r *discordgo.Ready) {
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

		err = botutil.UpdateStatus(s)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		}

		log.Printf("Logged in as %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
}
