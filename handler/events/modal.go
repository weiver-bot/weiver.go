package events

import (
	"log"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	parse "github.com/y2hO0ol23/weiver/handler/events/modal"
	botutil "github.com/y2hO0ol23/weiver/utils/bot"
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
		err = reviewutil.DeleteMessage(s, fromID, toID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		}

		// ready to change roles
		displayWas, err := role.GetDisplay(toID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
		roleList, err := db.GetRoleOnUser(toID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}

		// set db
		review, err := db.ModifyReviewByInfo(fromID, toID, score, title, content)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
		review, err = reviewutil.Send(s, i, review)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
		if review != nil {
			err = reviewutil.SendDM(s, review, i.Locale)
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				return
			}
		}
		err = botutil.UpdateStatus(s)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}

		// set role
		displayNow, err := role.GetDisplay(toID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
		if displayWas != displayNow {
			for _, roleDB := range roleList {
				role.Remove(s, roleDB.GuildID, toID, displayWas)
			}
			// will add new role by GuildMemberUpdate
			// so just remove
		}

	})
}
