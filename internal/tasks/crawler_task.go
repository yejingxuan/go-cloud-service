package tasks

import (
	"github.com/robfig/cron"
	logger "go-cloud-service/internal/log"
)

func StartCrawlerTask() {
	logger.Log.Info("start CrawlerTask")
	c := cron.New()
	c.AddFunc("0/5 * * * * ?", doCrawlerTask)
	c.Start()
}

func doCrawlerTask() {
	logger.Log.Info("5秒执行一次")
}
