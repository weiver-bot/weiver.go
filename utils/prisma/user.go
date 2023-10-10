package prisma

import (
	"github.com/y2hO0ol23/weiver/db"
)

func LoadUserById(id string) *db.UserModel {
	users, err := client.User.FindMany(
		db.User.ID.Equals(id),
	).Exec(ctx)
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
