package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		name := c.PostForm("name")
		message := c.PostForm("message")
		age := c.Query("age")
		gender := c.DefaultQuery("gender", "1")

		c.JSON(http.StatusOK, gin.H{"name": name, "message": message, "age": age, "gender": gender})
	})
	router.Run(":9000")
}
