package user

import (
	"github.com/gin-gonic/gin"
	"web/app/user/handlers"
)

func RegisterURL(r *gin.RouterGroup) *gin.RouterGroup {
	r.GET("/login", handlers.LoginGet)
	r.POST("/login", handlers.LoginPost)
	return r
}
