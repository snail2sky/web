package root

import (
	"github.com/gin-gonic/gin"
	"web/app/root/handlers"
)

func RegisterURL(r *gin.Engine) *gin.Engine {
	r.GET("/", handlers.Root)
	r.GET("/about", handlers.About)
	r.GET("/ip", handlers.IP)
	return r
}
