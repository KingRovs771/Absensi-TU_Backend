package Models

import (
	"absensiTU-Backend/Database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TURole struct {
	RoleId   int    `gorm:"primary_key;AUTO_INCREMENT;uniqueIndex" json:"role_id"`
	RoleUID  string `gorm:"type:varchar;uniqueIndex" json:"role_uid"`
	NameRole string `gorm:"type:varchar" json:"name_role"`
}

func (role *TURole) BeforeSave(*gorm.DB) error {
	if role.RoleUID == "" {
		uid, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		role.RoleUID = uid.String()
	}
	return nil
}

func (role *TURole) SaveRole() (*TURole, error) {
	err := Database.DB.Create(&role).Error
	if err != nil {
		return &TURole{}, err
	}
	return role, nil
}
