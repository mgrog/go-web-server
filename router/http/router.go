package http_router

import (
	"go_server/router/http/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/helloworld", handler.Helloworld)
	r.GET("/ping", handler.Ping)
}
