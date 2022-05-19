package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAzureBuildAgentIP(c *gin.Context) {
	azureBuildAgentIP := c.ClientIP()
	c.IndentedJSON(http.StatusOK, azureBuildAgentIP)
}

func main() {
	router := gin.Default()
	router.GET("/getazurebuildagentip", getAzureBuildAgentIP)
	router.Run("0.0.0.0:8080")
}
