package handlers

import (
    //"fmt"
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/a-hilaly/supfile-api/server/json"
    "github.com/a-hilaly/supfile-api/server/middlewares"

)

func CreateUserHandler(c *gin.Context) {
    req, err := json.Operate(c)
    //fmt.Println(req.Data)
    if err != nil {
        //fmt.Println(req, err)
        c.JSON(http.StatusOK, json.FailErrors(err.Error()))
        return
    }
    if valid := middlewares.ValidateToken(req.Token); valid != true {
        c.JSON(http.StatusOK, json.FailErrors(middlewares.ErrorUnknownToken.Error()))
        return
    }
    //fmt.Println(req)
    d := req.Data
    tr := validateNonEmpty(d.Username,
                           d.Email,
                           d.Password,
                           d.AuthType,
                           d.AccountType)
    if !tr {
        c.JSON(http.StatusOK, json.FailErrors(ErrorNonValidData.Error()))
        return
    }
    option := req.Option
    //fmt.Println(option)
    if option.Data {
        user, err := Manager.Data().CreateUser(d.Username, d.Email, d.AccountType, "simple", d.Password)
        if err != nil {
            c.JSON(http.StatusOK, json.FailErrors(err.Error()))
            return
        }
        if option.System {
            err2 := Manager.System().AddUserGroup("sftp", d.Username, d.Password)
            if err != nil {
                c.JSON(http.StatusOK, json.FailErrors(err.Error(), err2.Error()))
                return
            }
        }
        c.JSON(http.StatusOK, json.SuccessAID(user.AccountId))
        return
    }
    c.JSON(http.StatusOK, json.FailErrors("Missing options data | sys ?"))
}

func AuthentificateUserHandler(c *gin.Context) {
    // Fetch json
    req, err := json.Operate(c)
    if err != nil {
        //fmt.Println(req, err)
        c.JSON(http.StatusOK, json.FailErrors(err.Error()))
        return
    }
    // Validation
    if valid := middlewares.ValidateToken(req.Token); valid != true {
        c.JSON(http.StatusOK, json.FailErrors(middlewares.ErrorUnknownToken.Error()))
        return
    }
    d := req.Data
    o := req.Option
    tr := validateNonEmpty(d.Password,
                           d.AuthType)
    if !tr  {
        c.JSON(http.StatusOK, json.FailErrors(ErrorNonValidData.Error()))
        return
    }
    //fmt.Println(req)
    switch d.AuthType {
    // Simple Authentification
    case "simple":
        var un, value string
        // Find unique and value
        if nonNull(d.Username) {
            un, value = "username", d.Username
        } else if nonNull(d.Email) {
            un, value = "email", d.Email
        } else if nonNull(o.AccountId) {
            un, value = "account_id", o.AccountId
        } else {
            c.JSON(http.StatusOK, json.FailErrors(ErrorNonValidData.Error()))
            return
        }
        // Authentificate
        user, ok, err := Manager.Data().BasicAuthUser(un, value, d.Password)
        //
        if ok {
            c.JSON(http.StatusOK, json.SuccessAID(user.AccountId))
            return
        }
        c.JSON(http.StatusOK, json.FailErrors(err.Error()))
        return
    // Tokenized Authentification
    case "facebook", "twitter", "google":
        if nonNull(d.AuthToken) && nonNull(d.Email) {
            // A
            user, ok, err := Manager.Data().TokenizedAuthUser(d.Email, d.AuthType, d.AuthToken)
            //
            if ok {
                c.JSON(http.StatusOK, json.SuccessAID(user.AccountId))
                return
            }
            c.JSON(http.StatusOK, json.FailErrors(err.Error()))
            return
        }
        c.JSON(http.StatusOK, json.FailErrors(ErrorNonValidData.Error()))
        return
    default:
        c.JSON(http.StatusOK, json.FailErrors("Unkown authentification type"))
    }
}

func UpdateUserHandler(c *gin.Context) {
    req, err := json.Operate(c)
    if err != nil {
        //fmt.Println(req, err)
        c.JSON(http.StatusOK, json.FailErrors(err.Error()))
        return
    }
    if valid := middlewares.ValidateToken(req.Token); valid != true {
        c.JSON(http.StatusOK, json.FailErrors(middlewares.ErrorUnknownToken.Error()))
        return
    }
    //fmt.Println(req)
    option := req.Option
    tr := validateNonEmpty(option.AccountId)
    if !tr {
        c.JSON(http.StatusOK, json.FailErrors(ErrorNonValidData.Error()))
        return
    }
    data := req.Data
    tmap := data.MakeMap()
    err = Manager.Data().UpdateUserMap(option.AccountId, tmap)
    if err != nil {
        c.JSON(http.StatusOK, json.FailErrors(err.Error()))
    }
    user, errl := Manager.Data().GetUser(option.AccountId)
    if errl != nil {
        c.JSON(http.StatusOK, json.FailErrors(errl.Error()))
    }
    if nonNull(data.Username) {
        if err1 := Manager.System().ChangeUserName(user.Username, data.Username); err1 != nil {
            c.JSON(http.StatusOK, json.SuccessUserDataErrors(*user, err1.Error()))
            return
        }
    }
    if nonNull(data.Password) {
        if err2 := Manager.System().ChangeUserPassword(user.Username, data.Password); err2 != nil {
            c.JSON(http.StatusOK, json.SuccessUserDataErrors(*user, err2.Error()))
            return
        }
    }
    c.JSON(http.StatusOK, json.SuccessUserData(*user))
}

func UserDataHandler(c *gin.Context) {
    req, err := json.Operate(c)
    if err != nil {
        //fmt.Println(req, err)
        c.JSON(http.StatusOK, json.FailErrors(err.Error()))
        return
    }
    if valid := middlewares.ValidateToken(req.Token); valid != true {
        c.JSON(http.StatusOK, json.FailErrors(middlewares.ErrorUnknownToken.Error()))
        return
    }
    //fmt.Println(req)
    option := req.Option
    tr := validateNonEmpty(option.AccountId)
    if !tr {
        c.JSON(http.StatusOK, json.FailErrors(ErrorNonValidData.Error()))
        return
    }
    user, err := Manager.Data().GetUser(option.AccountId)
    if err != nil {
        c.JSON(http.StatusOK, json.FailErrors(err.Error()))
        return
    }
    c.JSON(http.StatusOK, json.SuccessUserData(*user))
}

func UserStatsHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "Hello" : "world",
    })
}

func DropUserHandler(c *gin.Context) {
    req, err := json.Operate(c)
    if err != nil {
        //fmt.Println(req, err)
        c.JSON(http.StatusOK, json.FailErrors(err.Error()))
        return
    }
    if valid := middlewares.ValidateToken(req.Token); valid != true {
        c.JSON(http.StatusOK, json.FailErrors(middlewares.ErrorUnknownToken.Error()))
        return
    }
    //fmt.Println(req)
    option := req.Option
    tr := validateNonEmpty(option.AccountId)
    if !tr {
        c.JSON(http.StatusOK, json.FailErrors(ErrorNonValidData.Error()))
        return
    }
    if option.Data {
        err := Manager.Data().DropUser(option.AccountId)
        if err != nil {
            c.JSON(http.StatusOK, json.FailErrors(err.Error()))
            return
        }
        if option.System {
            user, err := Manager.Data().GetUser(option.AccountId)
            if err != nil {
                c.JSON(http.StatusOK, json.FailErrors(err.Error()))
                return
            }
            err2 := Manager.System().DelUser(user.Username)
            if err != nil {
                c.JSON(http.StatusOK, json.FailErrors(err.Error(), err2.Error()))
                return
            }
            c.JSON(http.StatusOK, json.SuccessAID(user.AccountId))
            return
        }
        c.JSON(http.StatusOK, json.SuccessAID(option.AccountId))
        return
    }
    c.JSON(http.StatusOK, json.FailErrors("Missing request option"))
}
