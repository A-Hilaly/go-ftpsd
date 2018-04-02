package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func DevHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "Hello" : "world",
    })
}

func TestHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "Hello" : "world",
    })
}

func BenchmarksHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "Hello" : "world",
    })
}

func HealthCheckHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "Hello" : "world",
    })
}
