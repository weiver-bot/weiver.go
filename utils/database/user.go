package database

func LoadUserByID(id string) (*UserModel, error) {
	var users []UserModel

	err = db.Model(&UserModel{}).
		Where(&UserModel{
			ID: id,
		}).Limit(1).
		Find(&users).Error
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		var user = &UserModel{
			ID: id,
		}
		return user, db.Create(user).Error
	} else {
		return &users[0], nil
	}
}

func GetUserReviewCount(id string) (int64, error) {
	var count int64

	err = db.Model(&ReviewModel{}).
		Where(ReviewModel{
			ToID: id,
		}).Count(&count).Error

	return count, err
}

func GetUserScore(id string) (float64, error) {
	var res float64

	err = db.Model(&ReviewModel{}).
		Where(ReviewModel{
			ToID: id,
		}).Select("avg(Score)").Row().
		Scan(&res)

	return res, err
}
