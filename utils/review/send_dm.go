package reviewutil

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/utils/builder"
	db "github.com/y2hO0ol23/weiver/utils/database"
)

func SendDM(s *discordgo.Session, review *db.ReviewModel) {
	msg, _ := s.ChannelMessage(review.DMChannelID, review.DMMessageID)

	embed := builder.Embed().
		SetDescription(fmt.Sprintf("https://discord.com/channels/%s/%s/%s → <@%s>", review.GuildID, review.ChannelID, review.MessageID, review.FromID)).
		AddFields(&discordgo.MessageEmbedField{
			Name: "🔔 Your review has written",
		}).MessageEmbed

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

	msg, err := s.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return
	}
	db.UpdateDMMessageInfoByID(review.ID, channelID, msg.ID)
}
