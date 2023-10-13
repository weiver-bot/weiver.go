package reviewutil

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/utils/builder"
	db "github.com/y2hO0ol23/weiver/utils/database"
)

func EmbedBody(review *db.ReviewModel, thubnail string) *builder.EmbedStructure {
	return builder.Embed().
		SetFields(&discordgo.MessageEmbedField{
			Name:  fmt.Sprintf("📝 %s [%s%s]", review.Title, "★★★★★"[:review.Score*3], "☆☆☆☆☆"[review.Score*3:]),
			Value: fmt.Sprintf("```%s```", review.Content),
		}).
		SetThumbnail(&discordgo.MessageEmbedThumbnail{
			URL: thubnail,
		})
}

func EmbedMost(review *db.ReviewModel, thubnail string) *builder.EmbedStructure {
	return EmbedBody(review, thubnail).
		SetFooter(&discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("👍 %d", review.LikeTotal),
		}).
		SetTimeStamp(review.TimeStamp)
}
