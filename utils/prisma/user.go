package prisma

import (
	"github.com/y2hO0ol23/weiver/db"
)

func LoadUserById(id string) *db.UserModel {
	users, err := client.User.FindMany(
		db.User.ID.Equals(id),
	).Take(1).Exec(ctx)
	if err != nil {
		panic(err)
	}

	if len(users) == 0 {
		user, err := client.User.CreateOne(
			db.User.ID.Set(id),
		).Exec(ctx)
		if err != nil {
			panic(err)
		}
		return user
	} else {
		return &users[0]
	}
}

func GetUserScoreAverage(id string) (float64, int) {
	users, err := client.User.FindMany(
		db.User.ID.Equals(id),
	).With(
		db.User.Written.Fetch(),
	).Take(1).Exec(ctx)
	if err != nil {
		panic(err)
	}

	if len(users) == 0 {
		return 0.0, 0
	}

	user := users[0]
	sum := 0
	count := len(user.Written())
	for _, review := range user.Written() {
		sum += review.Score
	}
	return float64(sum) / float64(count), count
}

func GetBestReview(id string) *db.ReviewModel {
	user, err := client.User.FindUnique(
		db.User.ID.Equals(id),
	).With(
		db.User.Written.Fetch().OrderBy(
			db.Review.LikeTotal.Order(db.DESC),
		).Take(1),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}

	return &user.Written()[0]
}
