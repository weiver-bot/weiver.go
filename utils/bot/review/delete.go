package ReviewUtils

import (
	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/database"
)

func DeleteLastMessage(s *discordgo.Session, authorID string, subjectID string) error {
	// remove old reivew
	review, err := db.LoadReivewByInfo(authorID, subjectID)
	if err != nil {
		return err
	}

	if review != nil {
		s.ChannelMessageDelete(review.ChannelID, review.MessageID)
	}
	return nil
}
