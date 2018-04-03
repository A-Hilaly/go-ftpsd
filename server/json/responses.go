package json

import (
    //"github.com/gin-gonic/gin"
    us "github.com/a-hilaly/supfile-api/core/data"
)

type Response struct {
    Success bool        `json:"success"`
    Errors  []string    `json:"error"`
    Data    interface{} `json:"data"`
}

type reqAID struct {
    AccountId string `json:"account_id"`
}

type Reponser interface {
    MakeResponse(success bool, errors []string, data interface{}) Response
}

func MakeResponse(s bool, data interface{}, errs ...string) Response {
    return Response{
        Success : s,
        Errors  : errs,
        Data    : data,
    }
}

func SuccessN(data interface{}, errs ...string) Response {
    return MakeResponse(true, data, errs...)
}

func SuccessOnly() Response {
    return MakeResponse(true, nil)
}

func SuccessUserData(ud us.User) Response {
    return Response{}
}

func SuccessData(data interface{}) Response {
    return MakeResponse(true, data)
}

func SuccessAID(id string) Response {
    return SuccessData(reqAID{id})
}

func FailN(data interface{}, errs ...string) Response {
    return MakeResponse(false, data, errs...)
}

func FailErrors(errs ...string) Response {
    return MakeResponse(false, nil, errs...)
}

func FailOnly() Response {
    return MakeResponse(false, nil)
}
