package handlers

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/a-hilaly/supfile-api/server/json"
    "github.com/a-hilaly/supfile-api/server/middlewares"

)

func CreateUserHandler(c *gin.Context) {
    req, err := json.Operate(c)
    if err != nil {
        fmt.Println(req, err)
        c.JSON(http.StatusOK, json.FailErrors(err.Error()))
        return
    }
    if valid := middlewares.ValidateToken(req.Token); valid != true {
        c.JSON(http.StatusOK, json.FailErrors(middlewares.ErrorUnknownToken.Error()))
        return
    }
    fmt.Println(req)
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
    fmt.Println(option)
    if option.Data {
        user, err := Manager.Data().CreateUser(d.Username, d.Email, d.AccountType, "simple", d.Password)
        if err != nil {
            c.JSON(http.StatusOK, json.FailErrors(err.Error()))
            return
        }
        if option.System {
            err = Manager.System().AddUserGroup("sftp", d.Username, d.Password)
            if err != nil {
                c.JSON(http.StatusOK, json.FailErrors(err.Error()))
                return
            }
        }
        c.JSON(http.StatusOK, json.SuccessAID(user.AccountId))
    }
    c.JSON(http.StatusOK, json.FailErrors("why the fk do you ping me ?"))
}

func DropUserHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "Hello" : "world",
    })
}

func AuthentificateUserHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "Hello" : "world",
    })
}

func UpdateUserHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "Hello" : "world",
    })
}

func UserDataHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "Hello" : "world",
    })
}

func UserStatsHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "Hello" : "world",
    })
}
