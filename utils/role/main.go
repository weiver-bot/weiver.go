package role

import (
	"fmt"
	"os"

	db "github.com/y2hO0ol23/weiver/utils/database"
)

var roleFormat string = os.Getenv("ROLE_FORMAT")

func GetDisplay(memberID string) (string, error) {
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
	return fmt.Sprintf(roleFormat, avg), nil
}
