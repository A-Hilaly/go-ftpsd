package handlers

import (
	//"fmt"
	"net/http"

	"github.com/a-hilaly/go-ftpsd/server/json"
	"github.com/gin-gonic/gin"
)

func DevHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Hello": "world",
	})
}

func TestHandler(c *gin.Context) {
	_, err := json.Operate(c)
	if err == nil {
		c.JSON(http.StatusOK, json.Response{
			Success: true,
			Errors:  []string{"OK"},
			Data: struct {
				Test string `json:"test"`
			}{"ON"},
		})
	}
	c.JSON(http.StatusOK, json.Response{
		Success: false,
		Errors:  []string{"Test error"},
		Data: struct {
			Test string `json:"test"`
		}{"OFF"},
	})
}

func BenchmarksHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Benchmarks": "Not implemented",
	})
}

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"HealthCheck": "OK",
	})
}
