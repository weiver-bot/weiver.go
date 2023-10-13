package reviewutil

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/utils/database"
)

func UpdateStatus(s *discordgo.Session) {
	var IdleSince = 0

	err := s.UpdateStatusComplex(discordgo.UpdateStatusData{
		IdleSince: &IdleSince,
		Activities: []*discordgo.Activity{
			{
				Name:  "show reviews count",
				Type:  discordgo.ActivityTypeCustom,
				State: fmt.Sprintf("✏️ %d", db.GetReviewsCount()),
			},
		},
	})
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
	}
}
