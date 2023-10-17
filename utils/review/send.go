package reviewutil

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/utils/builder"
	db "github.com/y2hO0ol23/weiver/utils/database"
)

func Send(s *discordgo.Session, i *discordgo.InteractionCreate, review *db.ReviewModel) (*db.ReviewModel, error) {
	to, err := s.GuildMember(i.GuildID, review.ToID)
	if err != nil {
		return nil, err
	}

	button_good := builder.Button().
		SetCustomID(fmt.Sprintf("like_review_%d", review.ID)).
		SetLable("👍").
		SetStyle(discordgo.SecondaryButton)

	button_bad := builder.Button().
		SetCustomID(fmt.Sprintf("hate_review_%d", review.ID)).
		SetLable("👎").
		SetStyle(discordgo.SecondaryButton)

	err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
		Embeds: []*discordgo.MessageEmbed{
			EmbedMost(review, to.AvatarURL("")).
				SetDescription(fmt.Sprintf("<@%s> → <@%s>", review.FromID, review.ToID)).
				MessageEmbed,
		},
		Components: []discordgo.MessageComponent{
			builder.ActionRow().AddComponents(button_good, button_bad).ActionsRow,
		},
		AllowedMentions: &discordgo.MessageAllowedMentions{},
	}))
	if err != nil {
		return nil, err
	}

	// add msg data on db
	msg, err := s.InteractionResponse(i.Interaction)
	if err != nil {
		return nil, err
	}
	return db.UpdateMessageInfoByID(review.ID, i.GuildID, msg.ChannelID, msg.ID)
}
