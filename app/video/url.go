package video

import (
	"github.com/gin-gonic/gin"
	"web/app/video/handlers"
)

func RegisterURL(r *gin.RouterGroup) *gin.RouterGroup {
	r.GET("/", handlers.Download)
	return r
}
