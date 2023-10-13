package database

import (
	"log"
	"runtime/debug"
)

func LoadGuildByID(id string) GuildModel {
	var guilds []GuildModel

	err = db.Model(&GuildModel{}).
		Where(&GuildModel{
			ID: id,
		}).Limit(1).
		Find(&guilds).Error
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
	}

	if len(guilds) == 0 {
		guild := GuildModel{
			ID: id,
		}
		err = db.Create(&guild).Error
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		}

		return guild
	} else {
		return guilds[0]
	}
}

func UpdateGuildRoleOption(id string, value bool) {
	var guild GuildModel
	err := db.Model(&GuildModel{ID: id}).
		Updates(map[string]interface{}{
			"AllowRole":  value,
			"InProgress": true,
		}).
		Take(&guild).Error
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
	}
}

func EndOFGuildProgress(id string) {
	var guild GuildModel
	err := db.Model(&GuildModel{ID: id}).
		Updates(map[string]interface{}{
			"InProgress": false,
		}).
		Take(&guild).Error
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
	}
}

func GetGuildInProgress() *[]GuildModel {
	var guilds []GuildModel
	err := db.Model(&GuildModel{}).
		Where(map[string]interface{}{
			"InProgress": true,
		}).Find(&guilds).Error
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
	}

	return &guilds
}
