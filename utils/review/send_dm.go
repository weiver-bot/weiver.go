package reviewutil

import (
	"log"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/utils/database"
)

func SendDM(s *discordgo.Session, review *db.ReviewModel) {
	msg, _ := s.ChannelMessage(review.DMChannelID, review.DMMessageID)

	var channelID string
	if msg != nil {
		channelID = review.DMChannelID
		s.ChannelMessageDelete(review.DMChannelID, review.DMMessageID)
	} else {
		channel, _ := s.UserChannelCreate(review.ToID)
		if channel != nil {
			channelID = channel.ID
		}
	}

	msg, err := s.ChannelMessageSendEmbed(channelID, embedDM(review))
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return
	}
	db.UpdateDMMessageInfoByID(review.ID, channelID, msg.ID)
}

func ModifyDM(s *discordgo.Session, review *db.ReviewModel) {
	_, err := s.ChannelMessageEditEmbed(review.DMChannelID, review.DMMessageID, embedDM(review))
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return
	}
}
