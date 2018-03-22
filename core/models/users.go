package models

import (
	//"fmt"

	"github.com/a-hilaly/gears/hash"
	"github.com/a-hilaly/supfile-api/engine"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"size:64"`
	Firstname   string `gorm:"size:64"`
	Lastname    string `gorm:"size:64"`
	Email       string `gorm:"size:64"`
	Password    string `gorm:"size:64"`
	AccountType string `gorm:"size:64"`
}

func NewUser(name, nick, email, password, t string) {
	engine.DB.Create(&User{
		Name:     name,
		Nick:     nick,
		Email:    email,
		Password: password,
		Type:     t,
	})
}

func AuthentificateUser(email, password string) (*User, bool) {
	user := User{}
	if err := engine.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, false
	}
	if user.Password != "" && crypto.Md5(password) == user.Password {
		return &user, true
	}
	return nil, false
}

func UsersList() *[]User {
	return nil
}

func DropUser() {

}
