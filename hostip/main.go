package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHostIP(c *gin.Context) {
	hostIP := c.ClientIP()
	c.IndentedJSON(http.StatusOK, hostIP)
}

func main() {
	router := gin.Default()
	router.GET("/hostip", getHostIP)
	router.Run("0.0.0.0:8080")
}
