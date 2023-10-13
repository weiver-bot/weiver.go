package reviewutil

import (
	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/utils/database"
)

func SendDM(s *discordgo.Session, review *db.ReviewModel) error {
	s.ChannelMessageDelete(review.DMChannelID, review.DMMessageID)

	channel, _ := s.UserChannelCreate(review.ToID)
	if channel == nil {
		return nil
	}

	msg, err := s.ChannelMessageSendEmbed(channel.ID, embedDM(review))
	if err == nil {
		_, err = db.UpdateDMMessageInfoByID(review.ID, channel.ID, msg.ID)
	}
	return err
}

func ModifyDM(s *discordgo.Session, review *db.ReviewModel) {
	s.ChannelMessageEditEmbed(review.DMChannelID, review.DMMessageID, embedDM(review))
}
