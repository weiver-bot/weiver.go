package reviewutil

import (
	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/utils/database"
)

func SendDM(s *discordgo.Session, review *db.ReviewModel, locale discordgo.Locale) error {
	channel, err := s.UserChannelCreate(review.ToID)
	if err != nil {
		return nil
	}
	s.ChannelMessageDelete(channel.ID, review.DMMessageID)

	msg, _ := s.ChannelMessageSendEmbed(channel.ID, embedDM(review, locale))
	if msg != nil {
		_, err = db.UpdateDMMessageInfoByID(review.ID, channel.ID, msg.ID)
		return err
	}
	return nil
}

func ModifyDM(s *discordgo.Session, review *db.ReviewModel, locale discordgo.Locale) {
	channel, err := s.UserChannelCreate(review.ToID)
	if err != nil {
		return
	}
	s.ChannelMessageEditEmbed(channel.ID, review.DMMessageID, embedDM(review, locale))
}
