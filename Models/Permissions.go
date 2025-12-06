package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TUPermissions struct {
	PermissionId     int       `gorm:"primary_key;AUTO_INCREMENT;uniqueIndex" json:"permission_id"`
	PermissionUID    string    `gorm:"type:varchar" json:"permission_uid"`
	UsersUID         string    `gorm:"type:varchar;uniqueIndex" json:"users_uid"`
	PermissionDate   string    `gorm:"type:date" json:"permission_date"`
	PermissionStatus int       `gorm:"type:int" json:"permission_status"`
	ProcessedBy      string    `gorm:"type:varchar" json:"processed_by"`
	Description      string    `gorm:"type:text" json:"description"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`

	User TUUsers `gorm:"foreignKey:UsersUID;references:UsersUID"`
}

func (permissions *TUPermissions) BeforeSavePermissions(*gorm.DB) error {
	UniqueID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	permissions.PermissionUID = UniqueID.String()

	return nil
}
