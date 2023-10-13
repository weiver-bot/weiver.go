package database

import (
	"fmt"
	"log"
	"runtime/debug"
)

func GetRoleByInfo(guildID string, display string) *RoleModel {
	var roles []RoleModel

	err = db.Model(&RoleModel{}).
		Where(&RoleModel{
			GuildID: guildID,
			Display: display,
		}).Limit(1).
		Find(&roles).Error
	if err != nil {
		log.Println(fmt.Sprintf("Error: %v\n%v", err, string(debug.Stack())))
	}

	if len(roles) == 0 {
		return nil
	}
	return &roles[0]
}

func GetRoleByID(id string) *RoleModel {
	var roles []RoleModel

	err = db.Model(&RoleModel{}).
		Where(&RoleModel{
			ID: id,
		}).Limit(1).
		Find(&roles).Error
	if err != nil {
		log.Println(fmt.Sprintf("Error: %v\n%v", err, string(debug.Stack())))
	}

	if len(roles) == 0 {
		return nil
	}
	return &roles[0]
}

func CreateRole(roleID string, guildID string, display string) *RoleModel {
	role := &RoleModel{
		GuildID: guildID,
		RoleID:  roleID,
		Display: display,
		ID:      fmt.Sprintf("%s#%s", guildID, roleID),
	}
	err = db.Create(role).Error
	if err != nil {
		log.Println(fmt.Sprintf("Error: %v\n%v", err, string(debug.Stack())))
	}

	return role
}

func AddRoleOnUser(id string, userID string) {
	err = db.Model(&RoleModel{ID: id}).
		Association("User").
		Append(&UserModel{
			ID: userID,
		})
	if err != nil {
		log.Println(fmt.Sprintf("Error: %v\n%v", err, string(debug.Stack())))
	}
}

func RemoveRoleOnUser(id string, userID string) bool {
	err = db.Model(&RoleModel{ID: id}).
		Association("User").
		Delete(&UserModel{
			ID: userID,
		})
	if err != nil {
		log.Println(fmt.Sprintf("Error: %v\n%v", err, string(debug.Stack())))
	}

	count := db.Model(&RoleModel{ID: id}).Association("User").Count()
	if count == 0 {
		db.Delete(&RoleModel{ID: id})
		return false
	}
	return true
}

func GetRoleOnUser(userID string) []*RoleModel {
	var roles []*RoleModel

	err = db.Model(&UserModel{ID: userID}).
		Association("Role").
		Find(&roles)
	if err != nil {
		log.Println(fmt.Sprintf("Error: %v\n%v", err, string(debug.Stack())))
	}

	return roles
}
