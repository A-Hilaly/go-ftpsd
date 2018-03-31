package core

import (
    "github.com/a-hilaly/supfile-api/core/data"
)

type DataInterface interface {
    // Data Interface config
    Init()
    SetConfig()
    GetConfig()

    // Account data
    GetUserAccountID()
    GetUserData()

    // Login and subscription
    CreateUser()
    UserExist()
    BasicAuthUser()
    TokenizedAuthUser()

    // Deletion and management
    DropUser()
    ActivateUser()
    DeactivateUser()

    // User updates
    UpdateUser()
    ChangeUserPassword()
    ChangeUserAuthMethod()
    ChangeUserName()
    ChangeUserEmail()

    // User control
    UpdateUserQuota()
}

type DataManager struct {

}

// Create supfile user (Storage + system user)
func CreateUser(level int, username, email, authtype, pass, type_ string) (*data.User, bool, error){
    // level 1 create at db
    // level 2 create at sys
    // level 3 create at db and sys
    if !AllowCreateUser {
        return nil, false, ErrorRuleNotAllowed
    }
    var user *data.User
    var err error
    cdb, csys := false, false
    if AllowCreateUserDb && (level == 3 or level == 1) {
        if authtype == "simple" {
            user, err = data.NewUser(username, email, pass, type_, authtype, "")
        } else {
            user, err = data.NewUser(username, email, "", type_, authtype, pass)
        }
        if err != nil {
            return nil, false, err
        }
        cdb = true
    }
    if AllowCreateUserSys && (level == 3 or level = 2) {
        err = system.AddUser(FTPGroup, user.Username, user.Password)
        if err != nil {
            return user, cdb, err
        }
        csys = true
    }
    return user, cdb || csys, nil
    //_, _ := models.NewUser(first, last, email, password, type_)
}

func UserExist(level int, username, email, authtype, pass, type_ string) {
    // level param:
    //       1 : check at sys
    //       2 : check at db
    //       3 : check at sys and db
}
