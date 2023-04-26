package global

import (
	"fmt"
	"github.com/HelloMrShu/grpc_client/proto"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SrvConnService() {
	consul := ServerConfig.Consul
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%s/%s?wait=14s", consul.Host, consul.Port, ServerName),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		//grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)

	if err != nil {
		Logger.Fatal("[InitSrvConn] connection error")
	}

	RpcConn = proto.NewGreeterClient(conn)
}
