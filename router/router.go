package router

import (
	"github.com/HelloMrShu/grpc_client/controllers/hello"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	Api := router.Group("api")
	{
		Api.GET("/hello", hello.Say)
	}
}
