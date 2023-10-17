package reviewutil

import (
	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/utils/database"
)

func DeleteReviewMessage(s *discordgo.Session, fromID string, toID string) error {
	// remove old reivew
	review, err := db.LoadReivewByInfo(fromID, toID)
	if err != nil {
		return err
	}

	if review != nil {
		s.ChannelMessageDelete(review.ChannelID, review.MessageID)
	}
	return nil
}
