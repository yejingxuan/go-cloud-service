package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go-cloud-service/internal/handler"
	logger "go-cloud-service/internal/log"
	"go-cloud-service/internal/tasks"
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

func TaskRun(cmd *cobra.Command, args []string) {
	for _, arg := range args {
		switch arg {
		case "crawler":
			tasks.StartCrawlerTask()
		default:
			logger.Log.Info("error task arg", zap.Any("arg", arg))
		}
	}

	//监听 ctrl+c 退出命令，监听到执行退出
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	logger.Log.Info("signal received", zap.Any("signal", <-sigChan))
}

func StartGrpcServer() {
	logger.Log.Info("start grpc server success")
}

func StartHttpServer() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	//V1版本接口定义
	v1 := engine.Group("/service/api/v1/base")
	{
		v1.GET("/healthCheck", handler.HealthCheck)
		//获取首页数据
		v1.GET("/getIndexData", handler.GetIndexData)
	}
	port := viper.Get("general.app_port")
	logger.Log.Info("start http server success on port: ", zap.Any("port", port))
	engine.Run(fmt.Sprintf(":%d", port))
}
