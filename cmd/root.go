package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	logger "go-cloud-service/internal/log"
	"go.uber.org/zap"
	"os"
)

// 根命令
var rootCmd = &cobra.Command{
	Use:   "cloud",
	Short: "cloud-service",
	Long:  `cloud-service`,
	//执行根命令
	Run: RootRun,
}

// version子命令，输出当前项目版本
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the go-cloud-service version",
	Run: func(cmd *cobra.Command, args []string) {
		println("1.0.0")
	},
}

// task子命令，执行task
var taskCmd = &cobra.Command{
	Use:   "task [the task name]",
	Short: "execute tasks",
	Args:  cobra.MinimumNArgs(1),
	Run:   TaskRun,
}

// 项目初始化
func init() {
	// 初始化配置文件
	InitConfig()
	// 添加version命令
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(taskCmd)
}

// 初始化配置文件
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	//path, _ := os.Getwd()
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Log.Error("read config failed: %v", zap.Any("err", err))
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
