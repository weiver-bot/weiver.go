package events

import (
	"log"

	"github.com/bwmarrin/discordgo"
	parse "github.com/y2hO0ol23/weiver/handler/events/modal"
	db "github.com/y2hO0ol23/weiver/utils/database"
	reviewutil "github.com/y2hO0ol23/weiver/utils/review"
	"github.com/y2hO0ol23/weiver/utils/role"
)

func init() {
	events = append(events, func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionModalSubmit {
			return
		}

		data := i.ModalSubmitData()
		fromID, toID, ok := parse.Review.CustomID(data.CustomID)
		if !ok {
			log.Println("Error on parse CustomID")
			return
		}
		score, title, content, ok := parse.Review.ModalComponents(data.Components)
		if !ok {
			log.Println("Error on parse ModalComponents")
			return
		}

		// remove old reivew
		reviewutil.DeleteMessage(s, fromID, toID)

		// ready to change roles
		displayWas := role.GetDisplay(toID)
		roleList := db.GetRoleOnUser(toID)

		// set db
		review := db.ModifyReviewByInfo(fromID, toID, score, title, content)
		reviewutil.Resend(s, i, review, "written")

		// set role
		displayNow := role.GetDisplay(toID)
		if displayWas != displayNow {
			for _, roleDB := range roleList {
				role.Remove(s, roleDB.GuildID, toID, displayWas)
			}
			// will add new role by GuildMemberUpdate
			// so just remove
		}
	})
}
