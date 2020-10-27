package logger

import "go.uber.org/zap"

var Log *zap.Logger

func init() {
	// 设置日志级别
	//Logger, _ = NewDevelopment()
	Log, _ = zap.NewDevelopment()
}
