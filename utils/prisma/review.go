package prisma

import (
	"github.com/y2hO0ol23/weiver/db"
)

func LoadReivewByIds(fromId string, toId string) *db.ReviewModel {
	reviews, err := client.Review.FindMany(
		db.Review.FromID.Equals(fromId),
		db.Review.ToID.Equals(toId),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}

	if len(reviews) == 0 {
		return nil
	}
	return &reviews[0]
}

func ModifyReviewByIds(fromId string, toId string, score int, title string, content string) *db.ReviewModel {
	review := LoadReivewByIds(fromId, toId)

	LoadUserById(fromId)
	LoadUserById(toId)

	review, err := func() (*db.ReviewModel, error) {
		if review == nil {
			return client.Review.CreateOne(
				db.Review.Score.Set(score),
				db.Review.Title.Set(title),
				db.Review.Content.Set(content),
				db.Review.To.Link(
					db.User.ID.Equals(toId),
				),
				db.Review.From.Link(
					db.User.ID.Equals(fromId),
				),
				db.Review.LikeTotal.Set(0),
			).Exec(ctx)
		} else {
			return client.Review.FindUnique(
				db.Review.ID.Equals(review.ID),
			).Update(
				db.Review.Score.Set(score),
				db.Review.Title.Set(title),
				db.Review.Content.Set(content),
				db.Review.Likes.Unlink(),
				db.Review.Hates.Unlink(),
			).Exec(ctx)
		}
	}()

	if err != nil {
		panic(err)
	}
	return review
}

func UpdateIdsById(id int, guildId string, channelId string, messageId string) *db.ReviewModel {
	review, err := client.Review.FindUnique(
		db.Review.ID.Equals(id),
	).Update(
		db.Review.GuildID.Set(guildId),
		db.Review.ChannelID.Set(channelId),
		db.Review.MessageID.Set(messageId),
	).Exec(ctx)

	if err != nil {
		panic(err)
	}
	return review
}

var (
	ReviewActionHandler = map[string]func(id string) []db.ReviewSetParam{
		"like": func(id string) []db.ReviewSetParam {
			return []db.ReviewSetParam{
				db.Review.Likes.Link(
					db.User.ID.Equals(id),
				),
				db.Review.Hates.Unlink(
					db.User.ID.Equals(id),
				),
			}
		},
		"hate": func(id string) []db.ReviewSetParam {
			return []db.ReviewSetParam{
				db.Review.Likes.Unlink(
					db.User.ID.Equals(id),
				),
				db.Review.Hates.Link(
					db.User.ID.Equals(id),
				),
			}
		},
	}
)

func ReviewAction(id int, values []db.ReviewSetParam) *db.ReviewModel {
	review, err := client.Review.FindUnique(
		db.Review.ID.Equals(id),
	).With(
		db.Review.Likes.Fetch(),
		db.Review.Hates.Fetch(),
	).Update(
		values...,
	).Exec(ctx)

	if err != nil {
		panic(err)
	}
	review, err = client.Review.FindUnique(
		db.Review.ID.Equals(id),
	).Update(
		db.Review.LikeTotal.Set(len(review.Likes()) - len(review.Hates())),
	).Exec(ctx)

	if err != nil {
		panic(err)
	}
	return review
}
