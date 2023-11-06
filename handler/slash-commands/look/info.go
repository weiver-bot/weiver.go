package look

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/database"
	"github.com/y2hO0ol23/weiver/localization"
	"github.com/y2hO0ol23/weiver/utils/builder"
)

func Info(s *discordgo.Session, i *discordgo.InteractionCreate, subjectID string) {
	locale := i.Locale

	// load subject
	subject, err := s.GuildMember(i.GuildID, subjectID)
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return
	}

	// get average of subject score
	var avg float64 = 0.0
	count, err := db.GetUserReviewCount(subjectID)
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return
	}
	if count > 0 {
		avg, err = db.GetUserScore(subjectID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
	}
	countOutput := func() string {
		if count >= 100 {
			return "100+"
		} else if count >= 50 {
			return "50+"
		}
		return fmt.Sprintf("%v", count)
	}()

	// build embed
	embed := builder.Embed().
		SetDescription(fmt.Sprintf("<@%v> **"+os.Getenv("SCORE_UI_SHORT")+" (%v)**", subjectID, avg, countOutput)).
		SetThumbnail(&discordgo.MessageEmbedThumbnail{
			URL: subject.AvatarURL(""),
		})

	if count == 0 { // if no reviews
		embed.SetFields(&discordgo.MessageEmbedField{
			Name:  fmt.Sprintf("📑 %v", localization.Load(locale, "#look.info.IsNone")),
			Value: "``` ```",
		})
	} else { // show best review
		review, err := db.GetReviewBest(subjectID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}
		embed.
			SetFields(&discordgo.MessageEmbedField{
				Name:  fmt.Sprintf("📑 %v 〔%v%v〕", review.Title, "★★★★★"[:review.Score*3], "☆☆☆☆☆"[review.Score*3:]),
				Value: fmt.Sprintf("```%v```", review.Content),
			}).
			SetFooter(&discordgo.MessageEmbedFooter{
				Text: fmt.Sprintf("👍 %v", review.LikeTotal),
			})
	}

	// send embed
	err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
		Embeds: []*discordgo.MessageEmbed{
			embed.MessageEmbed,
		},
		Flags:           discordgo.MessageFlagsEphemeral,
		AllowedMentions: &discordgo.MessageAllowedMentions{},
	}))
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return
	}
}
