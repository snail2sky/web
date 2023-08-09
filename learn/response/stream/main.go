package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	router.GET("/gopher-stream", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

	router.GET("/file-stream", func(c *gin.Context) {
		file := "group.go"
		reader, err := os.Open(file)
		defer reader.Close()
		if err != nil {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		info, err := reader.Stat()
		if err != nil {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		c.DataFromReader(
			http.StatusOK, info.Size(), "content/plain", reader, map[string]string{})
	})
	router.Run(":9000")
}
