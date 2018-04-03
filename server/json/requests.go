package json

import (
    //"fmt"
    "github.com/gin-gonic/gin"
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
    Args        []string `json:"args"`
}


type expReqOption struct {
    AccountId  string   `json:"username"`
    System     bool     `json:"system"`
    Data       bool     `json:"data"`
    Actions    []string `json:"actions"`
}


type Requester interface {
    FromContext(c *gin.Context, data interface{}) error
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
