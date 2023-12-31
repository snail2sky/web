package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		ct := c.Copy()
		go func(c *gin.Context) {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(2 * time.Second)
			//log.Printf("%p", c)
			//log.Printf("%p", ct)

			// note that you are using the copied context "cCp", IMPORTANT
			log.Println("Done! in path " + ct.Request.URL.Path)
		}(ct)
	})

	r.GET("/long_sync", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(2 * time.Second)

		// since we are NOT using a goroutine, we do not have to copy the context
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":9000")
}
