package hello

import (
	"github.com/HelloMrShu/grpc_client/proto"
	"github.com/HelloMrShu/grpc_client/services"
	"github.com/gin-gonic/gin"
)

func Say(c *gin.Context) {
	req := &proto.HelloReq{}
	req.Name = c.DefaultQuery("name", "guest")

	res, _ := services.SayHello(c, req)
	c.JSON(0, res.Message)
}
