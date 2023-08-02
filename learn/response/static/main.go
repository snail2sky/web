package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"net/http"
)

var root = flag.String("root", ".", "root dir")

func main() {
	flag.Parse()
	router := gin.Default()
	//router.Static("/learn1", "../../../")
	router.StaticFS("/", http.Dir(*root))
	//router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	router.Run(":9000")
}
