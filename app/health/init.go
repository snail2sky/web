package health

import (
	"github.com/gin-gonic/gin"
	"github.com/snail2sky/web"
	"log"
)

var HealthyRouter *gin.Engine

func init() {
	HealthyRouter = web.RootRouter.Group("/health")
	log.Println(web.RootRouter, HealthyRouter)
}
