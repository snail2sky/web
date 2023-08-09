package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Download(c *gin.Context) {
	url, ok := c.GetQuery("url")
	if !ok {
		c.String(http.StatusBadRequest, "Must give a url\n")
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("your url is %s\n", url))
}
