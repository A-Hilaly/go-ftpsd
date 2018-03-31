package core

import (

    "github.com/a-hilaly/supfile-api/core/data"
    "github.com/a-hilaly/supfile-api/core/system"
)

// Create supfile user (Storage + system user)
func CreateUser(level int, username, email, authtype, pass, type_ string) (*data.User, bool, error){
    if !AllowCreateUser {
        return nil, false, ErrorRuleNotAllowed
    }
    var user *data.User
    var err error
    cdb, csys := false, false
    if AllowCreateUserDb {
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
    if AllowCreateUserSys {
        err := system.AddUser(FTPGroup, user.Username, user.Password)
        if err != nil {
            return user, cdb, err
        }
        csys = true
    }
    return user, cdb || csys, nil
    //_, _ := models.NewUser(first, last, email, password, type_)
}

func DropUser(email string) {
    //_ := models.DropUser()
}

func UpdateUserData() {

}

func BanUser() {

}

func AuthUser() {

}

func AnalyzeUserPartition() {

}

func GetUserQuora() {

}
