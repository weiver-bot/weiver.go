package database

import (
	"log"
	"runtime/debug"
)

func LoadUserByID(id string) *UserModel {
	var users []UserModel

	err = db.Model(&UserModel{}).
		Where(&UserModel{
			ID: id,
		}).Limit(1).
		Find(&users).Error
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
	}

	if len(users) == 0 {
		var user = &UserModel{
			ID: id,
		}
		err = db.Create(user).Error
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		}

		return user
	} else {
		return &users[0]
	}
}

func GetUserScoreAverage(id string) (float64, int) {
	var users []UserModel

	err = db.Model(&UserModel{}).
		Where(UserModel{
			ID: id,
		}).Limit(1).Preload("Written").
		Find(&users).Error
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
	}

	if len(users) == 0 || len(users[0].Written) == 0 {
		return 0.0, 0
	}

	user := users[0]
	sum := 0
	count := len(user.Written)
	for _, review := range user.Written {
		sum += review.Score
	}
	return float64(sum) / float64(count), count
}
