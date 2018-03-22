package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func CreateUserHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "Hello" : "world",
    })
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
