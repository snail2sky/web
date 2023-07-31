package health

import "github.com/gin-gonic/gin"

func init() {
	HealthyRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
