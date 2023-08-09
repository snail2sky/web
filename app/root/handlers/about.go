package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func About(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintln("This is a gin test api"))
}
