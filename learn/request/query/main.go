package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/get", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		c.JSON(http.StatusOK, gin.H{"id": id, "page": page})

	})
	router.Run(":9000")
}
