package database

func LoadGuildByID(id string) (*GuildModel, error) {
	var guilds []GuildModel

	err := db.Model(&GuildModel{}).
		Where(&GuildModel{
			ID: id,
		}).Limit(1).
		Find(&guilds).Error
	if err != nil {
		return nil, err
	}

	if len(guilds) == 0 {
		guild := GuildModel{
			ID: id,
		}
		return &guild, db.Create(&guild).Error
	}
	return &guilds[0], nil
}

func UpdateGuildRoleOption(id string, value bool) error {
	var guild GuildModel
	return db.Model(&GuildModel{ID: id}).
		Updates(map[string]interface{}{
			"AllowRole":  value,
			"InProgress": true,
		}).
		Take(&guild).Error
}

func EndOFGuildProgress(id string) error {
	var guild GuildModel
	return db.Model(&GuildModel{ID: id}).
		Updates(map[string]interface{}{
			"InProgress": false,
		}).
		Take(&guild).Error
}

func GetGuildInProgress() (*[]GuildModel, error) {
	var guilds []GuildModel
	err := db.Model(&GuildModel{}).
		Where(GuildModel{
			InProgress: true,
		}).Find(&guilds).Error

	return &guilds, err
}
