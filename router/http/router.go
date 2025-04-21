package http_router

import (
	"go_server/router/http/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/ping", handler.Ping)
	rg := r.Group("/v1")
	{
		rg.GET("/helloworld", handler.Helloworld)
	}
}
