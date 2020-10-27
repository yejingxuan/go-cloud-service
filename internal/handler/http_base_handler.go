package handler

import "github.com/gin-gonic/gin"

func HealthCheck(c *gin.Context) {
	rep := gin.H{
		"message": "ok",
		"code":    200,
	}
	c.JSON(200, rep)
}

func GetIndexData(c *gin.Context) {
	rep := gin.H{
		"message": "ok",
		"code":    200,
	}
	c.JSON(200, rep)
}
