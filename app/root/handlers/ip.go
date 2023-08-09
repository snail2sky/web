package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IP(c *gin.Context) {
	c.String(http.StatusOK, c.RemoteIP())
}
