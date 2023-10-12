package events

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	parse "github.com/y2hO0ol23/weiver/handler/events/modal"
	"github.com/y2hO0ol23/weiver/utils/builder"
	db "github.com/y2hO0ol23/weiver/utils/database"
)

func init() {
	events = append(events, func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionModalSubmit {
			return
		}

		data := i.ModalSubmitData()
		fromID, toID, ok := parse.Review.CustomID(data.CustomID)
		if !ok {
			log.Println("Error on parse CustomID")
			return
		}
		score, title, content, ok := parse.Review.ModalComponents(data.Components)
		if !ok {
			log.Println("Error on parse ModalComponents")
			return
		}

		// remove old reivew
		review := db.LoadReivewByInfo(fromID, toID)
		if review != nil {
			_, err := s.ChannelMessage(review.ChannelID, review.MessageID)
			if err == nil {
				s.ChannelMessageDelete(review.ChannelID, review.MessageID)
			}
		}

		// set db
		review = db.ModifyReviewByInfo(fromID, toID, score, title, content)
		ResendReview(s, i, review, "written")
	})
}

func ResendReview(s *discordgo.Session, i *discordgo.InteractionCreate, review *db.ReviewModel, comment string) bool {
	to, err := s.GuildMember(i.GuildID, review.ToID)
	if err != nil {
		log.Println(err)
		return false
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
			builder.Embed().
				SetDescription(fmt.Sprintf("<@%s> → <@%s>", review.FromID, review.ToID)).
				SetFields(&discordgo.MessageEmbedField{
					Name:  fmt.Sprintf("📝 %s [%s%s]", review.Title, "★★★★★"[:review.Score*3], "☆☆☆☆☆"[review.Score*3:]),
					Value: fmt.Sprintf("```%s```", review.Content),
				}).
				SetFooter(&discordgo.MessageEmbedFooter{
					Text: fmt.Sprintf("👍 %d", review.LikeTotal),
				}).
				SetThumbnail(&discordgo.MessageEmbedThumbnail{
					URL: to.User.AvatarURL(""),
				}).
				SetTimeStamp(review.TimeStamp).
				MessageEmbed,
		},
		Components: []discordgo.MessageComponent{
			builder.ActionRow().SetComponents(button_good, button_bad).ActionsRow,
		},
		AllowedMentions: &discordgo.MessageAllowedMentions{},
	}))
	if err != nil {
		log.Println(err)
		return false
	}

	// add msg data on db
	msg, err := s.InteractionResponse(i.Interaction)
	if err != nil {
		log.Println(err)
		return false
	}
	db.UpdateMessageInfoByID(review.ID, i.GuildID, msg.ChannelID, msg.ID)

	// send dm to subject
	channel, err := s.UserChannelCreate(review.ToID)
	if err != nil {
		//log.Println(err)
		// if bot, blocked, etc...
	} else if channel != nil {
		s.ChannelMessageSendEmbeds(channel.ID, []*discordgo.MessageEmbed{
			builder.Embed().
				SetFields(&discordgo.MessageEmbedField{
					Name:  fmt.Sprintf("🔔 Your review has %s", comment),
					Value: fmt.Sprintf("➥ https://discord.com/channels/%s/%s/%s", i.GuildID, msg.ChannelID, msg.ID),
				}).
				MessageEmbed,
		})
	}

	return true
}
