package reviewutil

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/utils/database"
)

func UpdateStatus(s *discordgo.Session) error {
	var (
		IdleSince int = 0
		avg       float64
		count     int64
	)

	avg, err := db.GetReviewsScoreAvg()
	if err != nil {
		return err
	}
	count, err = db.GetReviewsCount()
	if err != nil {
		return err
	}

	return s.UpdateStatusComplex(discordgo.UpdateStatusData{
		IdleSince: &IdleSince,
		Activities: []*discordgo.Activity{
			{
				Name:  "Reviews total count",
				Type:  discordgo.ActivityTypeCustom,
				State: fmt.Sprintf("🧾%d → ⭐%.1f", count, avg),
			},
		},
	})
}
