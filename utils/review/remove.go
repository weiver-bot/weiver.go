package reviewutil

import (
	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/utils/database"
)

func DeleteMessage(s *discordgo.Session, fromID string, toID string) {
	// remove old reivew
	review := db.LoadReivewByInfo(fromID, toID)
	if review != nil {
		_, err := s.ChannelMessage(review.ChannelID, review.MessageID)
		if err == nil {
			s.ChannelMessageDelete(review.ChannelID, review.MessageID)
		}
	}
}
