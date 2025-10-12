package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TUUsers struct {
	UsersId        int       `gorm:"primaryKey;AUTO_INCREMENT;uniqueIndex" json:"users_id"`
	UserUID        string    `gorm:"type:varchar" json:"user_uid"`
	RoleUID        string    `gorm:"type:varchar" json:"role_uid"`
	FullName       string    `gorm:"type:varchar" json:"full_name"`
	Email          string    `gorm:"type:varchar" json:"email"`
	Password       string    `gorm:"type:varchar" json:"password"`
	ProfilePicture string    `gorm:"type:text" json:"profile_picture"`
	CreateAt       time.Time `gorm:"autoCreateTime" json:"create_at"`
}

func (users *TUUsers) BeforeSaveUsers(*gorm.DB) error {
	UniqueID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	users.UserUID = UniqueID.String()

	return nil
}