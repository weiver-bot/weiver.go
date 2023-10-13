package database

import (
	"fmt"
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
		log.Println(fmt.Sprintf("Error: %v\n%v", err, string(debug.Stack())))
	}

	if len(guilds) == 0 {
		guild := GuildModel{
			ID: id,
		}
		err = db.Create(&guild).Error
		if err != nil {
			log.Println(fmt.Sprintf("Error: %v\n%v", err, string(debug.Stack())))
		}

		return guild
	} else {
		return guilds[0]
	}
}

func UpdateGuildRoleOption(id string, value bool) *GuildModel {
	var guild GuildModel
	err := db.Model(&GuildModel{ID: id}).
		Updates(map[string]interface{}{
			"AllowRole": value,
		}).
		Take(&guild).Error
	if err != nil {
		log.Println(fmt.Sprintf("Error: %v\n%v", err, string(debug.Stack())))
	}

	return &guild
}
