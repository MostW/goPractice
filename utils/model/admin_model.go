package model

import (
	"app-version-manager/utils/bcrypt"
	"github.com/jinzhu/gorm"
	"time"
	"app-version-manager/utils/initUtils"
)

type Admin struct {
	gorm.Model
	Name              string `gorm:"type:varchar(100);column:name"`
	AuthToken         string `gorm:"type:varchar(100);column:auth_token"`
	EncryptedPassword string `gorm:"type:varchar(100);column:encrypted_password'"`
}

type AdminSerializer struct {
	ID                uint       `json:"id"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at"`
	Name              string     `json:"name"`
	AuthToken         string     `json:"auth_token"`
	EncryptedPassword string     `json:"encrypted_password"`
}

func (a *Admin) Serializer() AdminSerializer {
	return AdminSerializer{
		ID:                a.ID,
		CreatedAt:         a.CreatedAt,
		UpdatedAt:         a.UpdatedAt,
		DeletedAt:         a.DeletedAt,
		Name:              a.Name,
		AuthToken:         a.AuthToken,
		EncryptedPassword: a.EncryptedPassword,
	}
}

func (a *Admin) UpdatePassWord(password string) {

	var newPassword string
	newPassword = password
	oldAuthToken := a.AuthToken
	a.EncryptedPassword = newPassword

	var admin Admin
	if dbError := initUtils.MYDB.Where("auth_token = ?", oldAuthToken).First(&admin).Error; dbError != nil {
		panic("no exists user")
	}

	if dbError := initUtils.MYDB.Save(&admin).Error; dbError != nil {
		panic("update error")
	}
	return

}

func (a *Admin) PasswordCheck(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}

