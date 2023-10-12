package role

import (
	"fmt"
	"os"

	db "github.com/y2hO0ol23/weiver/utils/database"
)

var roleFormat string = os.Getenv("ROLE_FORMAT")

func GetDisplay(memberID string) string {
	average, _ := db.GetUserScoreAverage(memberID)
	return fmt.Sprintf(roleFormat, average)
}
