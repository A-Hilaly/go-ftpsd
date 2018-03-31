package handlers

import (
    "github.com/gin-gonic/gin"
)

type Requester interface {
    FromContext(c *gin.Context) error
}

type Reponser interface {
    OnContext(c *gin.Context) error
}

type Response struct {
    success int         `json:"id"`
    Errors  []string    `json:"error"`
    Data    interface{} `json:"data"`
}


type RequestJson struct {
    Token string      `json:"id"`
    Data  interface{} `json:"data"`
}
