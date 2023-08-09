package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Root(c *gin.Context) {
	c.Redirect(http.StatusFound, "/ip")
}
