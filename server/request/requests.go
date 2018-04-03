package request

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

type Requester interface {
    FromContext(c *gin.Context, data interface{}) error
}


type RequestJson struct {
    Token string      `json:"token"`
    Data  interface{} `json:"data"`
}


type Req struct {}

func (r *Req) FromContext(c *gin.Context, data interface{}) (*RequestJson, error) {
    req := &RequestJson{Data:data}
    fmt.Println(req, "hello")
    err := c.ShouldBindJSON(req)
    fmt.Println(err, "hello")
    if err != nil {
        return nil, err
    }
    return req, nil
}

var Request = Req{}
