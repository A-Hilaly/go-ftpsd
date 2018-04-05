package json

import (
    //"fmt"

    //"github.com/fatih/structs"
    "github.com/gin-gonic/gin"

    "github.com/a-hilaly/gears/crypto"
)

type RequestJson struct {
    Token  string       `json:"token"`
    Data   expReqData   `json:"data"`
    Option expReqOption `json:"option"`
}

type expReqData struct {
    Username    string   `json:"username"`
    Firstname   string   `json:"firstname"`
    Lastname    string   `json:"lastname"`
    Email       string   `json:"email"`
    Password    string   `json:"password"`
    AccountType string   `json:"account_type"`
    AuthType    string   `json:"auth_type"`
    AuthToken   string   `json:"auth_token"`
    State       string   `json:"state"`
    MaxStorage  int      `json:"max_storage"`
}

type expReqOption struct {
    AccountId  string   `json:"account_id"`
    System     bool     `json:"system"`
    Data       bool     `json:"data"`
    Actions    []string `json:"actions"`
}

type Requester interface {
    FromContext(c *gin.Context, data interface{}) error
}

func (erd *expReqData) MakeMap() (m map[string]interface{}) {
    m = make(map[string]interface{}, 10)
    if erd.Username != "" {
        m["username"] = erd.Username
    }
    if erd.Email != "" {
        m["email"] = erd.Email
    }
    if erd.Lastname != "" {
        m["lastname"] = erd.Lastname
    }
    if erd.Firstname != "" {
        m["firstname"] = erd.Firstname
    }
    if erd.AccountType != "" {
        m["account_type"] = erd.AccountType
    }
    if erd.Password != "" {
        m["password"] = crypto.Md5(erd.Password)
    }
    if erd.AuthType != "" {
        m["auth_type"] = erd.AuthType
    }
    if erd.AuthToken != "" {
        m["auth_token"] = erd.AuthToken
    }
    if erd.State != "" {
        m["state"] = erd.State
    }
    if erd.MaxStorage > 0 {
        m["max_storage"] = erd.MaxStorage
    }
    return
}

func fromContext(c *gin.Context, data expReqData, option expReqOption) (*RequestJson, error) {
    req := &RequestJson{
        Data:data,
        Option:option,
    }
    err := c.ShouldBindJSON(req)
    if err != nil {
        return req, err
    }
    return req, nil
}

func Operate(c *gin.Context) (*RequestJson, error){
    return fromContext(c, expReqData{}, expReqOption{})
}
