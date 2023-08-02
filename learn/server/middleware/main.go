package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Logger1() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("log1 before")

		c.Next()

		log.Println("log1 after")

		status := c.Writer.Status()
		log.Println(status)
	}
}

func Logger2() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("log2 before")
		c.Next()
		log.Println("log2 after")
	}
}

func main() {
	r := gin.New()
	r.Use(Logger1())
	r.Use(Logger2())

	r.GET("/test", func(c *gin.Context) {
		// it would print: "12345"
		log.Println("in test")
	})

	// Listen and serve on 0.0.0.0:9000
	r.Run(":9000")
}
