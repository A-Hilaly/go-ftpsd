package data

import "github.com/jinzhu/gorm"

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
