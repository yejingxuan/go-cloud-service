package cmd

import (
	"fmt"
	"github.com/gin-contrib/cors"
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
	"time"
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

/*func CallPythonRun(cmd *cobra.Command, args []string) {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}

	gostr := "foo"
	pystr := python.PyString_FromString(gostr)
	str := python.PyString_AsString(pystr)
	fmt.Println("hello [", str, "]")
}*/

// 开启天气预报
func WeatherReportRun(cmd *cobra.Command, args []string) {
	tasks.StartWeatherReport()
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

	engine.Use(getCors())
	engine.Use(gin.Recovery())
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

func getCors() gin.HandlerFunc {
	// 支持跨域
	mwCORS := cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type",
			"X-Requested-With", "X-Request-ID", "X-HTTP-Method-Override",
			"Content-Type", "Upload-Length", "Upload-Offset", "Tus-Resumable",
			"Upload-Metadata", "Upload-Defer-Length", "Upload-Concat"},
		ExposeHeaders: []string{"Content-Type", "Upload-Offset", "Location",
			"Upload-Length", "Tus-Version", "Tus-Resumable", "Tus-Max-Size",
			"Tus-Extension", "Upload-Metadata", "Upload-Defer-Length", "Upload-Concat"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 2400 * time.Hour,
	})
	return mwCORS
}
