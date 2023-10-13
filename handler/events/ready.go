package events

import (
	"log"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/utils/database"
	reviewutil "github.com/y2hO0ol23/weiver/utils/review"
)

func init() {
	events = append(events, func(s *discordgo.Session, r *discordgo.Ready) {
		guildsDB := db.GetGuildInProgress()
		if guildsDB != nil {
			for _, guildDB := range *guildsDB {
				db.EndOFGuildProgress(guildDB.ID)
			}
		}
		reviewutil.UpdateStatus(s)
		log.Printf("Logged in as %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
}
