package database

import (
	"fmt"
)

func GetRoleByInfo(guildID string, display string) (*RoleModel, error) {
	var roles []RoleModel

	err := db.Model(&RoleModel{}).
		Where(&RoleModel{
			GuildID: guildID,
			Display: display,
		}).Limit(1).
		Find(&roles).Error
	if err != nil {
		return nil, err
	}

	if len(roles) == 0 {
		return nil, nil
	}
	return &roles[0], nil
}

func GetRoleByID(id string) (*RoleModel, error) {
	var roles []RoleModel

	err := db.Model(&RoleModel{}).
		Where(&RoleModel{
			ID: id,
		}).Limit(1).
		Find(&roles).Error
	if err != nil {
		return nil, err
	}

	if len(roles) == 0 {
		return nil, nil
	}
	return &roles[0], nil
}

func CreateRole(roleID string, guildID string, display string) (*RoleModel, error) {
	role := &RoleModel{
		GuildID: guildID,
		RoleID:  roleID,
		Display: display,
		ID:      fmt.Sprintf("%v#%v", guildID, roleID),
	}
	return role, db.Create(role).Error
}

func AddRoleOnUser(id string, userID string) error {
	return db.Model(&RoleModel{ID: id}).
		Association("User").
		Append(&UserModel{
			ID: userID,
		})
}

func RemoveRoleOnUser(id string, userID string) (bool, error) {
	err := db.Model(&RoleModel{ID: id}).
		Association("User").
		Delete(&UserModel{
			ID: userID,
		})
	if err != nil {
		return false, err
	}

	count := db.Model(&RoleModel{ID: id}).Association("User").Count()
	if count == 0 {
		return false, db.Delete(&RoleModel{ID: id}).Error
	}
	return true, nil
}

func GetRoleOnUser(userID string) ([]*RoleModel, error) {
	var roles []*RoleModel

	err := db.Model(&UserModel{ID: userID}).
		Association("Role").
		Find(&roles)

	return roles, err
}
