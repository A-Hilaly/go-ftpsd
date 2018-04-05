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
	Username    string `gorm:"size:64" json:"username"`
	Firstname   string `gorm:"size:64" json:"firstname"`
	Lastname    string `gorm:"size:64" json:"lastname"`

    // Authentification fields
	Email       string `gorm:"size:64" json:"email"`
	Password    string `gorm:"size:64" json:"-"`

    // Account specifications
    AccountId   string `gorm:"size:64" json:"account_id"`  // Generated
	AccountType string `gorm:"size:64" json:"account_type"`  // Admin  | Default
    AuthType    string `gorm:"size:64" json:"auth_type"`  // Normal | Facebook | Google
    AuthToken   string `gorm:"size:64" json:"auth_token"`  // Goten from Facebook/Google

    // Other Specifications
    State       string `gorm:"size:64" json:"state"`
    MaxStorage  int    `json:"auth_token"`  // in Mb
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
