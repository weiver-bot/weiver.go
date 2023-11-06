package BotUtilss

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/database"
)

func UpdateStatus(s *discordgo.Session) error {
	var IdleSince int = 0

	text, err := StateText()
	if err != nil {
		return err
	}

	return s.UpdateStatusComplex(discordgo.UpdateStatusData{
		IdleSince: &IdleSince,
		Activities: []*discordgo.Activity{
			{
				Name:  "Reviews total count",
				Type:  discordgo.ActivityTypeCustom,
				State: text,
			},
		},
	})
}

func StateText() (string, error) {
	var (
		avg   float64 = 0
		count int64
	)

	count, err := db.GetReviewsCount()
	if err != nil {
		return "", err
	}
	if count > 0 {
		avg, err = db.GetReviewsScoreAvg()
		if err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("Total "+os.Getenv("SCORE_UI_SHORT")+" (%v)", avg, count), nil
}
