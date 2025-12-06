package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TUUsers struct {
	UsersId        int       `gorm:"primaryKey;AUTO_INCREMENT;uniqueIndex" json:"users_id"`
	UsersUID       string    `gorm:"type:varchar" json:"users_uid"`
	RoleUID        string    `gorm:"type:varchar" json:"role_uid"`
	FullName       string    `gorm:"type:varchar" json:"full_name"`
	Email          string    `gorm:"type:varchar" json:"email"`
	Password       string    `gorm:"type:varchar" json:"password"`
	ProfilePicture string    `gorm:"type:text" json:"profile_picture"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (users *TUUsers) BeforeSaveUsers(*gorm.DB) error {
	UniqueID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	users.UsersUID = UniqueID.String()

	return nil
}
