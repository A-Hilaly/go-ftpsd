package data

import (
    "github.com/jinzhu/gorm"

    "github.com/a-hilaly/supfile-api/core/data/engine"
)

// User model
// table create as 'user'
type User struct {
	gorm.Model
    // fields become low cased strings
    // Username -> 'username'
    // AccountType -> 'account_type'
    // NOTE: id field is create by default
    //       Generated

    // User specifications
	Username    string `gorm:"size:64"`
	Firstname   string `gorm:"size:64"`
	Lastname    string `gorm:"size:64"`

    // Authentification fields
	Email       string `gorm:"size:64"`
	Password    string `gorm:"size:64"`

    // Account specifications
    AccountId   string `gorm:"size:64"`  // Generated
	AccountType string `gorm:"size:64"`  // Admin  | Default
    AuthType    string `gorm:"size:64"`  // Normal | Facebook | Google
    AuthToken   string `gorm:"size:64"`  // Goten from Facebook/Google

    // Other Specifications
    State       string `gorm:"size:64"`
    MaxStorage  int // in Mb
}

// Create user table if dosent exist else pass
func autoMigrateUserTable() {
    engine.DB.AutoMigrate(User{})
}

// Create a new user in gorm database engine, mysql in our case
func newUser(un,
             email,
             password,
             type_,
             authtype,
             authtoken string) (*User, error) {
    id := HashUQ(un)
    if password == "" && authtoken == "" {
        return nil, ErrorAuthMethodNotSet
    }
    if exist, _ := userExistBy("account_id", id); exist {
        return nil, ErrorUserAlreadyExist
    }
    if exist, _ := userExistBy("email", email); exist {
        return nil, ErrorUserEmailExist
    }
    user := User{
        Username    :             un,
		Email       :          email,
		Password    : Hash(password),
		AccountType :          type_,
        AccountId   :             id,
        AuthType    :       authtype,
        AuthToken   :      authtoken,
	}
	if err := engine.DB.Create(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
