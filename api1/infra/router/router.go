package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Attach(rg *gin.RouterGroup) {
	rg.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This is App 1 root"})
	})

	rg.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from App 1"})
	})
}
