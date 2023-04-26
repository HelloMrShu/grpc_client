package services

import (
	"context"
	"github.com/HelloMrShu/grpc_client/global"
	"github.com/HelloMrShu/grpc_client/proto"
)

func SayHello(ctx context.Context, req *proto.HelloReq) (*proto.HelloResp, error) {
	resp, err := global.RpcConn.SayHello(ctx, req)
	if err != nil {
		global.Logger.Error(err.Error())
	}

	return resp, nil
}
