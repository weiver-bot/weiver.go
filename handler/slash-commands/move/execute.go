package move

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/database"
	"github.com/y2hO0ol23/weiver/localization"
	ReviewUtils "github.com/y2hO0ol23/weiver/utils/bot/review"
	"github.com/y2hO0ol23/weiver/utils/builder"
)

func Execute(s *discordgo.Session, i *discordgo.InteractionCreate, locale discordgo.Locale, subjectID string) {
	authorID := i.Interaction.Member.User.ID

	review, err := db.LoadReivewByInfo(authorID, subjectID)
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
	}
	if review == nil {
		err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
			Content: fmt.Sprintf("`%v`", localization.Load(locale, "#move.IsNone")),
			Flags:   discordgo.MessageFlagsEphemeral,
		}))
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		}
		return
	}

	subject, err := s.GuildMember(i.GuildID, subjectID)
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return
	}

	_, err = s.ChannelMessage(review.ChannelID, review.MessageID)
	if err != nil {
		review, err := ReviewUtils.SendReview(s, i, review)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
		if review != nil {
			ReviewUtils.ModifyReviewDM(s, review, locale)
		}
		return
	}

	move := builder.Button().
		SetCustomID("move-review").
		SetLable(localization.Load(locale, "#move.Move")).
		SetStyle(discordgo.SuccessButton)

	err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
		Embeds: []*discordgo.MessageEmbed{
			ReviewUtils.BaseEmbed(review, subject.AvatarURL("")).
				SetDescription(fmt.Sprintf("https://discord.com/channels/%v/%v/%v", review.GuildID, review.ChannelID, review.MessageID)).
				MessageEmbed,
		},
		Components: []discordgo.MessageComponent{
			builder.ActionRow().AddComponents(move).ActionsRow,
		},
		Flags: discordgo.MessageFlagsEphemeral,
	}))
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return
	}

	msg, err := s.InteractionResponse(i.Interaction)
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return
	}

	var handler func(*discordgo.Session, *discordgo.InteractionCreate)
	handler = func(s *discordgo.Session, iter *discordgo.InteractionCreate) {
		if iter.Type != discordgo.InteractionMessageComponent || i.Interaction.Member.User.ID != iter.Interaction.Member.User.ID {
			s.AddHandlerOnce(handler)
			return
		}
		if iter.Interaction.Message.ID != msg.ID {
			return
		}

		data := iter.MessageComponentData()
		if data.ComponentType != discordgo.ButtonComponent || data.CustomID != "move-review" {
			return
		}

		s.InteractionResponseDelete(i.Interaction)
		reviewNow, err := db.LoadReivewByInfo(authorID, subjectID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
		if review.TimeStamp != reviewNow.TimeStamp {
			err := s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					builder.Embed().
						SetDescription(fmt.Sprintf("❌ %v", localization.Load(locale, "$review.IsEdited"))).
						MessageEmbed,
				},
			}))
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			}
			return
		}
		err = ReviewUtils.DeleteLastMessage(s, authorID, subjectID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		}
		review, err = ReviewUtils.SendReview(s, iter, review)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		}
		if review != nil {
			ReviewUtils.ModifyReviewDM(s, review, locale)
		}
	}
	s.AddHandlerOnce(handler)
}
