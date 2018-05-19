package data

import (
	"fmt"

	"github.com/a-hilaly/gears/crypto"
	"github.com/a-hilaly/go-ftpsd/core/data/engine"
)

// Hash algorithm
var (
	Hash   = crypto.Md5
	HashUQ = crypto.Sha256
)

var uniqueFields []string = []string{"username", "email", "account_id", "auth_token"}

func isUnique(s string) bool {
	for _, elem := range uniqueFields {
		if elem == s {
			return true
		}
	}
	return false
}

// Init
func initDatabase() {
	autoMigrateUserTable()
	fmt.Println("Init: Automigrate user table")
}

func selectUserBy(by string, value string) (*User, error) {
	if !isUnique(by) {
		return nil, ErrorNotUnique
	}
	user := User{}
	if err := engine.DB.Where(by+"= ?", value).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func selectUsersBy(by string, value string) (*[]User, error) {
	var users []User
	if err := engine.DB.Where(by+"= ?", value).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func userExistBy(by string, value string) (bool, error) {
	if _, err := selectUserBy(by, value); err != nil {
		return false, err
	}
	return true, nil
}

// Remove User
func dropUserBy(by, value string) error {
	// Drop by email
	user, err := selectUserBy(by, value)
	if err != nil {
		return err
	}
	err = engine.DB.Unscoped().Delete(user).Error
	return err
}

func selectAllUsers() (*[]User, error) {
	return selectUsersBy("state", "Active")
}
