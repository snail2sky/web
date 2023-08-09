package root

import (
	"github.com/gin-gonic/gin"
	"web/app/user"
	"web/app/video"
)

var Router *gin.Engine

func init() {
	Router = RegisterURL(gin.Default())
	user.Router = user.RegisterURL(Router.Group("/user"))
	video.Router = video.RegisterURL(Router.Group("/video"))
}
