package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginGet(c *gin.Context) {
	c.String(http.StatusOK, "In login get func")
}

func LoginPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "login successful", "data": gin.H{}})
}
