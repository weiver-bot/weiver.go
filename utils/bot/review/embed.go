package ReviewUtils

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/database"
	"github.com/y2hO0ol23/weiver/localization"
	"github.com/y2hO0ol23/weiver/utils/builder"
)

func BaseEmbed(review *db.ReviewModel, thubnail string) *builder.EmbedStructure {
	return builder.Embed().
		SetFields(&discordgo.MessageEmbedField{
			Name:  fmt.Sprintf("📝 %v [%v%v]", review.Title, "★★★★★"[:review.Score*3], "☆☆☆☆☆"[review.Score*3:]),
			Value: fmt.Sprintf("```%v```", review.Content),
		}).
		SetThumbnail(&discordgo.MessageEmbedThumbnail{
			URL: thubnail,
		})
}

func BaseEmbedWithFooter(review *db.ReviewModel, thubnail string) *builder.EmbedStructure {
	return BaseEmbed(review, thubnail).
		SetFooter(&discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("👍 %v", review.LikeTotal),
		}).
		SetTimeStamp(review.TimeStamp)
}

func DMEmbed(review *db.ReviewModel, locale discordgo.Locale) *discordgo.MessageEmbed {
	return builder.Embed().
		SetDescription(fmt.Sprintf(
			"https://discord.com/channels/%v/%v/%v → <@%v>",
			review.GuildID,
			review.ChannelID,
			review.MessageID,
			review.AuthorID,
		)).
		AddFields(&discordgo.MessageEmbedField{
			Name: fmt.Sprintf("🔔 %v", localization.Load(locale, "$review.DM")),
		}).MessageEmbed
}
