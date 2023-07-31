package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

var RootRouter *gin.Engine

func init() {
	RootRouter = gin.Default()
}

func main() {
	err := RootRouter.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatal(err)
	}
}
