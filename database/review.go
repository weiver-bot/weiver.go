package database

import (
	"time"
)

func LoadReivewByID(id int) (*ReviewModel, error) {
	var reviews []ReviewModel

	err := db.Model(&ReviewModel{}).
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

func LoadReivewByInfo(authorID string, subjectID string) (*ReviewModel, error) {
	var reviews []ReviewModel

	err := db.Model(&ReviewModel{}).
		Where(ReviewModel{
			AuthorID:  authorID,
			SubjectID: subjectID,
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

func ModifyReviewByInfo(authorID string, subjectID string, score int, title string, content string) (*ReviewModel, error) {
	review, err := LoadReivewByInfo(authorID, subjectID)
	if err != nil {
		return nil, err
	}

	LoadUserByID(authorID)
	LoadUserByID(subjectID)

	if review == nil {
		review = &ReviewModel{
			Score:     score,
			Title:     title,
			Content:   content,
			SubjectID: subjectID,
			AuthorID:  authorID,
			TimeStamp: time.Now(),
		}
		return review, db.Create(review).Error
	}

	// remove old data
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

	return review, err
}

func UpdateMessageInfoByID(id int, guildID string, channelID string, messageID string) (*ReviewModel, error) {
	var review ReviewModel

	err := db.Model(&ReviewModel{ID: id}).
		Updates(ReviewModel{
			ChannelID: channelID,
			MessageID: messageID,
			GuildID:   guildID,
		}).
		Take(&review).Error

	return &review, err
}

func UpdateDMMessageInfoByID(id int, channelID string, messageID string) (*ReviewModel, error) {
	var review ReviewModel

	err := db.Model(&ReviewModel{ID: id}).
		Updates(ReviewModel{
			DMMessageID: messageID,
		}).
		Take(&review).Error

	return &review, err
}

func GetReviewBest(id string) (*ReviewModel, error) {
	var reviews []ReviewModel

	err := db.Model(&ReviewModel{}).
		Where(&ReviewModel{
			SubjectID: id,
		}).Order("Like_Total desc, Time_Stamp asc").Limit(1).
		Find(&reviews).Error
	if err != nil {
		return nil, err
	}

	if len(reviews) == 0 {
		return nil, nil
	}
	return &reviews[0], nil
}

var (
	ReviewButtonHandler = map[string]func(reviewID int, userID string) (*ReviewModel, error){
		"like": func(reviewID int, userID string) (*ReviewModel, error) {
			err := db.Model(&ReviewModel{ID: reviewID}).
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
			err := db.Model(&ReviewModel{ID: reviewID}).
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

	err := db.Model(&ReviewModel{ID: reviewID}).
		Updates(map[string]interface{}{
			"LikeTotal": likeCount - hateCount,
		}).
		Take(&review).Error

	return &review, err
}

func GetReviewsByUserID(id string) (*[]ReviewModel, error) {
	var reviews []ReviewModel

	err := db.Model(&ReviewModel{}).
		Where(&ReviewModel{
			SubjectID: id,
		}).Order("Like_Total desc, Time_Stamp asc").
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
	var avg float64 = 0

	err := db.Model(&ReviewModel{}).
		Select("avg(Score)").Row().
		Scan(&avg)

	return avg, err
}

func GetReviewsCount(relate string) (int64, error) {
	var count int64

	this := db.Model(&ReviewModel{})

	if relate != "" {
		this.Where("author_id = ?", relate).Or("subject_id = ?", relate)
	}

	return count, this.Count(&count).Error
}

func GetReviews(from int, count int, orderby string, relate string) (*[]ReviewModel, error) {
	var reviews []ReviewModel

	this := db.Model(&ReviewModel{})

	if relate != "" {
		this.Where("author_id = ?", relate).Or("subject_id = ?", relate)
	}

	err := this.Order(orderby).
		Offset(from).Limit(count).
		Find(&reviews).Error

	return &reviews, err
}
