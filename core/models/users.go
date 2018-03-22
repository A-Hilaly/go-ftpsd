package models

import (
	"fmt"
    "errors"

	"github.com/jinzhu/gorm"
    "github.com/a-hilaly/gears/crypto"
    "github.com/a-hilaly/supfile-api/core/engine"
)

// Init
func Init() {
    AutoMigrateUserTable()
    fmt.Println("Init: Automigrate user table")
}

// Hash algorithm
var Hash = crypto.Md5

// User model
// table create as 'user'
type User struct {
	gorm.Model
    // fields become low cased strings
    // Username -> 'username'
    // AccountType -> 'account_type'
    // NOTE: id field is create by default
    //       Generated
	Username    string `gorm:"size:64"`
	Firstname   string `gorm:"size:64"`
	Lastname    string `gorm:"size:64"`
	Email       string `gorm:"size:64"`
	Password    string `gorm:"size:64"`
	AccountType string `gorm:"size:64"`
}

// Create user table if dosent exist else pass
func AutoMigrateUserTable() {
    engine.DB.AutoMigrate(User{})
}


// Create a new user in gorm database engine, mysql in our case
func NewUser(un, first, last, email, password, type_ string) (*User, error) {
    user := User{
        Username    :             un,
		Firstname   :          first,
		Lastname    :           last,
		Email       :          email,
		Password    : Hash(password),
		AccountType :          type_,
	}
	if err := engine.DB.Create(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

// Authentificate user
func AuthentificateUser(email, password string) (*User, bool, error) {
	user := User{}
	if err := engine.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, false, err
	}
    // Authentification (MD5)
	if user.Password != "" && Hash(password) == user.Password {
		return &user, true, nil
	}
	return nil, false, errors.New("Authentification failed")
}

// Remove User
func DropUser(email string) error {
    // Drop by email
    //engine.DB.Where("email = ?", email).First(&user{})
    //engine.DB.Delete(user)
    return nil
}
