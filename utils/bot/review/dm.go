package ReviewUtils

import (
	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/database"
)

func SendReviewDM(s *discordgo.Session, review *db.ReviewModel, locale discordgo.Locale) error {
	channel, err := s.UserChannelCreate(review.SubjectID)
	if err != nil {
		return nil
	}
	s.ChannelMessageDelete(channel.ID, review.DMMessageID)

	msg, _ := s.ChannelMessageSendEmbed(channel.ID, DMEmbed(review, locale))
	if msg != nil {
		_, err = db.UpdateDMMessageInfoByID(review.ID, channel.ID, msg.ID)
		return err
	}
	return nil
}

func ModifyReviewDM(s *discordgo.Session, review *db.ReviewModel, locale discordgo.Locale) {
	channel, err := s.UserChannelCreate(review.SubjectID)
	if err != nil {
		return
	}
	s.ChannelMessageEditEmbed(channel.ID, review.DMMessageID, DMEmbed(review, locale))
}
