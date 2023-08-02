package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/301", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.google.com/")
	})

	router.GET("/302", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://www.google.com/")
	})

	router.GET("/307", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "https://www.google.com/")
	})

	router.Run(":9000")
}
