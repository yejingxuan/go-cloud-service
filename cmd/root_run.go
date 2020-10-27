package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-cloud-service/internal/handler"
	logger "go-cloud-service/internal/log"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func RootRun(cmd *cobra.Command, args []string) {
	logger.Log.Info("execute RootRun")
	go StartHttpServer()
	go StartGrpcServer()

	//监听 ctrl+c 退出命令，监听到执行退出
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	logger.Log.Info("signal received", zap.Any("signal", <-sigChan))

}

func StartGrpcServer() {
	logger.Log.Info("start grpc server success")
}

func StartHttpServer() {
	logger.Log.Info("start http server success")
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	//V1版本接口定义
	v1 := engine.Group("/service/api/v1/base")
	{
		v1.GET("/healthCheck", handler.HealthCheck)
		//获取首页数据
		v1.GET("/getIndexData", handler.GetIndexData)
	}
	engine.Run(":8081")
}
