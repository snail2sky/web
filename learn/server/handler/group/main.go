package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/login", loginEndpoint)
		v1.GET("/submit", submitEndpoint)
		v1.GET("/read", readEndpoint)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.GET("/login", loginEndpoint)
		v2.GET("/submit", submitEndpoint)
		v2.GET("/read", readEndpoint)
	}

	router.Run(":9000")
}

func loginEndpoint(c *gin.Context) {
	log.Println("in login")
	c.String(http.StatusOK, "login")
}

func submitEndpoint(c *gin.Context) {
	log.Println("in submit")
	c.String(http.StatusOK, "submit")
}
func readEndpoint(c *gin.Context) {
	log.Println("in read")
	c.String(http.StatusOK, "read")
}
