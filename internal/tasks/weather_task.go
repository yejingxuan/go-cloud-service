package tasks

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/robfig/cron"
	logger "go-cloud-service/internal/log"
	"go-cloud-service/internal/utils"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
)

type WeatherInfo struct {
	Date        string
	Weather     string
	Temperature string
	Air         string
}

// 开启天气预报任务
func StartWeatherReport() {
	logger.Log.Info("start WeatherReportTask")
	c := cron.New()
	c.AddFunc("0 21 20 * * ?", doWeatherReportTask)
	c.Start()

}

func doWeatherReportTask() {
	info, _ := getWeatherInfo()
	logger.Log.Info("start WeatherReportTask", zap.Any("info", info))
	handleWeather(info)
}

func handleWeather(info WeatherInfo) {
	if !strings.Contains(info.Weather, "雨") {
		return
	}

	mailTo := []string{
		"710819495@qq.com",
		"962752834@qq.com",
		"1364675851@qq.com",
	}
	//邮件主题为"Hello"
	subject := "【天气预报】"
	// 邮件正文
	body := "亲爱的，武汉市明天天气信息如下：\n" +
		"温度："+info.Temperature+", 天气："+info.Weather+", 空气质量："+info.Air+"，出门记得带伞哦"
	err := utils.SendMail(mailTo, subject, body)
	if err != nil {
		logger.Log.Info("发送成功")
	}
}

func getWeatherInfo() (WeatherInfo, error) {
	var info WeatherInfo
	resp, err := http.Get("http://tianqi.2345.com/today-57494.htm")
	if err != nil {
		return info, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logger.Log.Info("解析失败")
	}

	tomorrow := time.Now().Unix() + 24*60*60
	formatTomorrow := time.Unix(tomorrow, 0).Format("01/02")
	println(formatTomorrow)

	doc.Find(".seven-day-item").Each(func(i int, selection *goquery.Selection) {
		date := selection.Find("em").Text()
		if date == formatTomorrow {
			weather := selection.Find("i").Text()
			temperature := selection.Find(".tem-show").Text()
			air := selection.Find(".wea-qulity").Text()
			info.Date = date
			info.Weather = weather
			info.Temperature = temperature
			info.Air = air
		}
	})
	return info, nil
}
