package events

import (
	"log"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/database"
	"github.com/y2hO0ol23/weiver/handler/events/modal"
	BotUtils "github.com/y2hO0ol23/weiver/utils/bot"
	ReviewUtils "github.com/y2hO0ol23/weiver/utils/bot/review"
	TagUtils "github.com/y2hO0ol23/weiver/utils/bot/tag"

	g "github.com/y2hO0ol23/weiver/handler"
)

func init() {
	g.EventList = append(g.EventList, func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionModalSubmit {
			return
		}

		data := i.ModalSubmitData()
		authorID, subjectID, ok := modal.Parse.CustomID(data.CustomID)
		if !ok {
			log.Println("Error on parse CustomID")
			return
		}
		score, title, content, ok := modal.Parse.ModalComponents(data.Components)
		if !ok {
			log.Println("Error on parse ModalComponents")
			return
		}

		// remove old reivew
		err := ReviewUtils.DeleteLastMessage(s, authorID, subjectID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		}

		// ready to change roles
		displayWas, err := TagUtils.GetScoreUIShort(subjectID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}

		// set db
		review, err := db.ModifyReviewByInfo(authorID, subjectID, score, title, content)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
		review, err = ReviewUtils.SendReview(s, i, review)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
		if review != nil {
			err = ReviewUtils.SendReviewDM(s, review, i.Locale)
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				return
			}
		}
		err = BotUtils.UpdateStatus(s)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}

		// set role
		displayNow, err := TagUtils.GetScoreUIShort(subjectID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
		if displayWas != displayNow {
			roles, err := db.GetRoleOnUser(subjectID)
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				return
			}
			for _, roleDB := range roles {
				TagUtils.RemoveTag(s, roleDB.GuildID, subjectID, displayWas)
			}
			// will add new role by GuildMemberUpdate
			// so just remove
		}

	})
}
