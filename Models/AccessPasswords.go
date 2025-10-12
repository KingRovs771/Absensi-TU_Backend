package Models

import (
	"absensiTU-Backend/Database"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type TUAccessPasswords struct {
	AccessPasswordsId  int       `gorm:"primary_key;AUTO_INCREMENT;uniqueIndex" json:"access_passwords_id"`
	AccessPasswordsUID string    `gorm:"type:varchar" json:"access_passwords_uid"`
	PasswordAccess     string    `gorm:"type:varchar" json:"password_access"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func hashPA(passwordacess string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwordacess), 12)
	return string(bytes), err
}

func (AP *TUAccessPasswords) beforeSaveAccessPasswords(*gorm.DB) error{
	UniqueUID, err := uuid.NewRandom()

	if err != nil {
		return err
	}

	AP.AccessPasswordsUID = UniqueUID.String()
	
	hashPasswordAccess, err := hashPA(AP.PasswordAccess)
	if err != nil{
		return err
	}
	AP.PasswordAccess = hashPasswordAccess

	return nil
}

func (AP *TUAccessPasswords) saveAccessPassword() (*TUAccessPasswords, error){
	err := Database.DB.Create(&AP).Error

	if err != nil{
		return &TUAccessPasswords{}, err
	}

	return AP, nil
}

func (AP *TUAccessPasswords) validateAccessPassword(PasswordAccess string)error {
	return bcrypt.CompareHashAndPassword([]byte(AP.PasswordAccess), []byte(PasswordAccess))
}
