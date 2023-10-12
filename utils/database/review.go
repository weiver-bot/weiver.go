package database

import (
	"log"
	"time"
)

func LoadReivewByID(id int) *ReviewModel {
	var reviews []ReviewModel

	err = db.Model(&ReviewModel{ID: id}).
		Limit(1).
		Find(&reviews).Error
	if err != nil {
		log.Println(err)
	}

	if len(reviews) == 0 {
		return nil
	}
	return &reviews[0]
}

func LoadReivewByInfo(fromID string, toID string) *ReviewModel {
	var reviews []ReviewModel

	err = db.Model(&ReviewModel{}).
		Where(ReviewModel{
			FromID: fromID,
			ToID:   toID,
		}).Limit(1).
		Find(&reviews).Error
	if err != nil {
		log.Println(err)
	}

	if len(reviews) == 0 {
		return nil
	}
	return &reviews[0]
}

func ModifyReviewByInfo(fromID string, toID string, score int, title string, content string) *ReviewModel {
	review := LoadReivewByInfo(fromID, toID)

	LoadUserByID(fromID)
	LoadUserByID(toID)

	if review == nil {
		err = db.Create(&ReviewModel{
			Score:     score,
			Title:     title,
			Content:   content,
			ToID:      toID,
			FromID:    fromID,
			TimeStamp: time.Now(),
		}).Error
		if err != nil {
			log.Println(err)
		}
	} else {
		err = db.Model(&ReviewModel{ID: review.ID}).
			Association("Like").
			Clear()
		if err != nil {
			log.Println(err)
		}

		err = db.Model(&ReviewModel{ID: review.ID}).
			Association("Hate").
			Clear()

		if err != nil {
			log.Println(err)
		}

		err = db.Model(&ReviewModel{ID: review.ID}).
			Updates(ReviewModel{
				Score:     score,
				Title:     title,
				Content:   content,
				LikeTotal: 0,
				TimeStamp: time.Now(),
			}).Error
		if err != nil {
			log.Println(err)
		}
	}

	return LoadReivewByInfo(fromID, toID)
}

func UpdateMessageInfoByID(id int, guildID string, channelID string, messageID string) *ReviewModel {
	var review ReviewModel

	err := db.Model(&ReviewModel{ID: id}).
		Updates(ReviewModel{
			GuildID:   guildID,
			ChannelID: channelID,
			MessageID: messageID,
		}).
		Take(&review).Error

	if err != nil {
		log.Println(err)
	}
	return &review
}

func GetReviewBest(id string) *ReviewModel {
	var reviews []ReviewModel

	err = db.Model(&ReviewModel{}).
		Where(&ReviewModel{
			ToID: id,
		}).Order("Score desc").Limit(1).
		Find(&reviews).Error
	if err != nil {
		log.Println(err)
	}

	return &reviews[0]
}

var (
	ReviewButtonHandler = map[string]func(reviewID int, userID string) *ReviewModel{
		"like": func(reviewID int, userID string) *ReviewModel {
			err = db.Model(&ReviewModel{ID: reviewID}).
				Association("Like").
				Append(&UserModel{
					ID: userID,
				})
			if err != nil {
				log.Println(err)
			}

			err = db.Model(&ReviewModel{ID: reviewID}).
				Association("Hate").
				Delete(&UserModel{
					ID: userID,
				})
			if err != nil {
				log.Println(err)
			}

			return reviewButtonHandlerFianl(reviewID)
		},
		"hate": func(reviewID int, userID string) *ReviewModel {
			err = db.Model(&ReviewModel{ID: reviewID}).
				Association("Hate").
				Append(&UserModel{
					ID: userID,
				})
			if err != nil {
				log.Println(err)
			}

			err = db.Model(&ReviewModel{ID: reviewID}).
				Association("Like").
				Delete(&UserModel{
					ID: userID,
				})
			if err != nil {
				log.Println(err)
			}

			return reviewButtonHandlerFianl(reviewID)
		},
	}
)

func reviewButtonHandlerFianl(reviewID int) *ReviewModel {
	var review ReviewModel

	likeCount := db.Model(&ReviewModel{ID: reviewID}).Association("Like").Count()
	hateCount := db.Model(&ReviewModel{ID: reviewID}).Association("Hate").Count()

	err = db.Model(&ReviewModel{ID: reviewID}).
		Updates(ReviewModel{
			LikeTotal: likeCount - hateCount,
		}).
		Take(&review).Error
	if err != nil {
		log.Println(err)
	}

	return &review
}

func GetReviewsByUserID(id string) *[]ReviewModel {
	var reviews []ReviewModel

	err = db.Model(&ReviewModel{}).
		Where(&ReviewModel{
			ToID: id,
		}).Order("Score desc").
		Find(&reviews).Error
	if err != nil {
		log.Println(err)
	}

	if len(reviews) == 0 {
		return nil
	}
	return &reviews
}
