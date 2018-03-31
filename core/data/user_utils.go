package data

import (
	"fmt"

    "github.com/a-hilaly/gears/crypto"
    "github.com/a-hilaly/supfile-api/core/data/engine"
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
func Init() {
    autoMigrateUserTable()
    fmt.Println("Init: Automigrate user table")
}

// Create user table if dosent exist else pass
func autoMigrateUserTable() {
    engine.DB.AutoMigrate(User{})
}

// Hash algorithm
var (
    Hash   = crypto.Md5
    HashUQ = crypto.Sha256
)

// Create a new user in gorm database engine, mysql in our case
func NewUser(un,
             email,
             password,
             type_,
             authtype,
             authtoken string) (*User, error) {
    id := HashUQ(un)
    if password == "" && authtoken == "" {
        return nil, ErrorAuthMethodNotSet
    }
    if exist, _ := UserExistBy("account_id", id); exist {
        return nil, ErrorUserAlreadyExist
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

func SelectUserBy(by string, value string) (*User, error) {
    if !isUnique(by) {
        return nil, ErrorNotUnique
    }
    user := User{}
	if err := engine.DB.Where(by + "= ?", value).First(&user).Error; err != nil {
		return nil, err
	}
    return &user, nil
}

func SelectUsersBy(by string, value string) (*[]User, error) {
    var users []User
    if err := engine.DB.Where(by + "= ?", value).Find(&users).Error; err != nil {
        return nil, err
    }
    return &users, nil
}

func UserExistBy(by string, value string) (bool, error) {
    if _, err := SelectUserBy(by, value); err != nil {
        return false, err
    }
    return true, nil
}

// Remove User
func DropUserBy(by ,value string) error {
    // Drop by email
    user, err := SelectUserBy(by, value)
    if err != nil {
        return err
    }
    err = engine.DB.Unscoped().Delete(user).Error
    return err
}


func SelectAllUsers() (*[]User, error) {
    return SelectUsersBy("state", "Active")
}
