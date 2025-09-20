package Models

import "time"

type TUPermissions struct {
	PermissionId     int       `gorm:"primary_key;AUTO_INCREMENT;uniqueIndex" json:"permission_id"`
	PermissionUID    string    `gorm:"type:varchar" json:"permission_uid"`
	PermissionDate   string    `gorm:"type:date" json:"permission_date"`
	PermissionStatus int       `gorm:"type:int" json:"permission_status"`
	ProcessedBy      string    `gorm:"type:varchar" json:"processed_by"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
}
