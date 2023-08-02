package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func listen1() {
	router := gin.Default()
	http.ListenAndServe(":9000", router)
}

func listen2() {
	router := gin.Default()

	s := &http.Server{
		Addr:           ":9000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func main() {
	//listen1()
	listen2()
}
