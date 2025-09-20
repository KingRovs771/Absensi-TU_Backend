package Models

import "time"

type TUAccessPasswords struct {
	AccessPasswordsId  int       `gorm:"primary_key;AUTO_INCREMENT;uniqueIndex" json:"access_passwords_id"`
	AccessPasswordsUID string    `gorm:"type:varchar" json:"access_passwords"`
	PasswordAccess     string    `gorm:"type:varchar" json:"password_access"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
}
