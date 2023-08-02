package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Person struct {
	ID   string `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

func handle(c *gin.Context) {
	var person Person
	if err := c.ShouldBindUri(&person); err != nil {
		log.Print(err)
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
}
func main() {
	route := gin.Default()
	route.GET("/:name/:id", handle)
	route.Run(":9000")
}
