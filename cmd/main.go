package main

import (
	"context"
	"fmt"
	. "github.com/HelloMrShu/grpc_client/global"
	"github.com/HelloMrShu/grpc_client/router"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// 初始化
	Initialize()
	//初始化rpc server 连接
	SrvConnService()
	// 加载路由
	engine := gin.New()
	engine.Use(LoggerMiddleware(Logger), RecoveryMiddleware(Logger, false))
	router.InitRouter(engine)

	// 启动服务
	srv := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: engine,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			Logger.Error("服务启动失败", zap.String("error ", err.Error()))
		}
	}()

	// 关闭服务器
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("\n开始关闭服务器 ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("关闭服务器异常:", err)
	}

	fmt.Println("服务器已关闭")
}
