package prisma

import (
	"context"

	"github.com/y2hO0ol23/weiver/db"
)

var (
	client *db.PrismaClient
	ctx    context.Context
)

type Review struct {
}

func init() {
	client = db.NewClient()
	ctx = context.Background()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
}

func LoadUserById(id string) *db.UserModel {
	user, err := client.User.FindUnique(
		db.User.ID.Equals(id),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}

	if user == nil {
		user, err = client.User.CreateOne(
			db.User.ID.Set(id),
		).Exec(ctx)
		if err != nil {
			panic(err)
		}
	}

	return user
}

func LoadReivewByIds(fromId string, toId string) *db.ReviewModel {
	review, err := client.Review.FindMany(
		db.Review.FromID.Equals(fromId),
		db.Review.ToID.Equals(toId),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}

	if len(review) == 0 {
		return nil
	}
	return &review[0]
}
