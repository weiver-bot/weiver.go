package TagUtils

import (
	"fmt"
	"os"

	db "github.com/y2hO0ol23/weiver/database"
)

func GetScoreUIShort(memberID string) (string, error) {
	count, err := db.GetUserReviewCount(memberID)
	if err != nil {
		return "", err
	}

	var avg float64 = 0.0
	if count > 0 {
		avg, err = db.GetUserScore(memberID)
		if err != nil {
			return "", err
		}
	}
	return fmt.Sprintf(os.Getenv("SCORE_UI_SHORT"), avg), nil
}
