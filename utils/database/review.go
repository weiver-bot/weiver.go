package database

import (
	"time"
)

func LoadReivewByID(id int) (*ReviewModel, error) {
	var reviews []ReviewModel

	err = db.Model(&ReviewModel{}).
		Where(ReviewModel{
			ID: id,
		}).Limit(1).
		Find(&reviews).Error
	if err != nil {
		return nil, err
	}

	if len(reviews) == 0 {
		return nil, nil
	}
	return &reviews[0], nil
}

func LoadReivewByInfo(fromID string, toID string) (*ReviewModel, error) {
	var reviews []ReviewModel

	err = db.Model(&ReviewModel{}).
		Where(ReviewModel{
			FromID: fromID,
			ToID:   toID,
		}).Limit(1).
		Find(&reviews).Error
	if err != nil {
		return nil, err
	}

	if len(reviews) == 0 {
		return nil, nil
	}
	return &reviews[0], nil
}

func ModifyReviewByInfo(fromID string, toID string, score int, title string, content string) (*ReviewModel, error) {
	review, err := LoadReivewByInfo(fromID, toID)
	if err != nil {
		return nil, err
	}

	LoadUserByID(fromID)
	LoadUserByID(toID)

	if review == nil {
		review = &ReviewModel{
			Score:     score,
			Title:     title,
			Content:   content,
			ToID:      toID,
			FromID:    fromID,
			TimeStamp: time.Now(),
		}
		err = db.Create(review).Error
		if err != nil {
			return nil, err
		}
	} else {
		err = db.Model(&ReviewModel{ID: review.ID}).
			Association("Like").
			Clear()
		if err != nil {
			return nil, err
		}

		err = db.Model(&ReviewModel{ID: review.ID}).
			Association("Hate").
			Clear()
		if err != nil {
			return nil, err
		}

		err = db.Model(&ReviewModel{ID: review.ID}).
			Updates(map[string]interface{}{
				"Score":     score,
				"Title":     title,
				"Content":   content,
				"LikeTotal": 0,
				"TimeStamp": time.Now(),
			}).
			Take(review).Error
		if err != nil {
			return nil, err
		}
	}

	return review, nil
}

func UpdateMessageInfoByID(id int, guildID string, channelID string, messageID string) (*ReviewModel, error) {
	var review ReviewModel

	err = db.Model(&ReviewModel{ID: id}).
		Updates(ReviewModel{
			ChannelID: channelID,
			MessageID: messageID,
			GuildID:   guildID,
		}).
		Take(&review).Error
	if err != nil {
		return nil, err
	}

	return &review, nil
}

func UpdateDMMessageInfoByID(id int, channelID string, messageID string) (*ReviewModel, error) {
	var review ReviewModel

	err = db.Model(&ReviewModel{ID: id}).
		Updates(ReviewModel{
			DMChannelID: channelID,
			DMMessageID: messageID,
		}).
		Take(&review).Error
	if err != nil {
		return nil, err
	}

	return &review, nil
}

func GetReviewBest(id string) (*ReviewModel, error) {
	var reviews []ReviewModel

	err = db.Model(&ReviewModel{}).
		Where(&ReviewModel{
			ToID: id,
		}).Order("Score desc").Limit(1).
		Find(&reviews).Error
	if err != nil {
		return nil, err
	}

	return &reviews[0], nil
}

var (
	ReviewButtonHandler = map[string]func(reviewID int, userID string) (*ReviewModel, error){
		"like": func(reviewID int, userID string) (*ReviewModel, error) {
			err = db.Model(&ReviewModel{ID: reviewID}).
				Association("Like").
				Append(&UserModel{
					ID: userID,
				})
			if err != nil {
				return nil, err
			}

			err = db.Model(&ReviewModel{ID: reviewID}).
				Association("Hate").
				Delete(&UserModel{
					ID: userID,
				})
			if err != nil {
				return nil, err
			}

			return reviewButtonHandlerFianl(reviewID)
		},
		"hate": func(reviewID int, userID string) (*ReviewModel, error) {
			err = db.Model(&ReviewModel{ID: reviewID}).
				Association("Hate").
				Append(&UserModel{
					ID: userID,
				})
			if err != nil {
				return nil, err
			}

			err = db.Model(&ReviewModel{ID: reviewID}).
				Association("Like").
				Delete(&UserModel{
					ID: userID,
				})
			if err != nil {
				return nil, err
			}

			return reviewButtonHandlerFianl(reviewID)
		},
	}
)

func reviewButtonHandlerFianl(reviewID int) (*ReviewModel, error) {
	var review ReviewModel

	likeCount := db.Model(&ReviewModel{ID: reviewID}).Association("Like").Count()
	hateCount := db.Model(&ReviewModel{ID: reviewID}).Association("Hate").Count()

	err = db.Model(&ReviewModel{ID: reviewID}).
		Updates(ReviewModel{
			LikeTotal: likeCount - hateCount,
		}).
		Take(&review).Error
	if err != nil {
		return nil, err
	}

	return &review, nil
}

func GetReviewsByUserID(id string) (*[]ReviewModel, error) {
	var reviews []ReviewModel

	err = db.Model(&ReviewModel{}).
		Where(&ReviewModel{
			ToID: id,
		}).Order("Score desc").
		Find(&reviews).Error
	if err != nil {
		return nil, err
	}

	if len(reviews) == 0 {
		return nil, nil
	}
	return &reviews, nil
}

func GetReviewsScoreAvg() (float64, error) {
	var avg float64
	err := db.Model(&ReviewModel{}).
		Select("avg(score)").Row().
		Scan(&avg)
	if err != nil {
		return 0, err
	}

	return avg, nil
}

func GetReviewsCount() (int64, error) {
	var count int64
	err := db.Model(&ReviewModel{}).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
