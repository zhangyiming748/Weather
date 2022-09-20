package Weather

import (
	"encoding/json"
	"github.com/zhangyiming748/Weather/log"
	"io"
	"net/http"
	"strings"
	"time"
)

type Report struct {
	Status    string `json:"status"`   // 返回状态
	Count     string `json:"count"`    // 返回结果总数目
	Info      string `json:"info"`     // 返回的状态信息
	Infocode  string `json:"infocode"` // 返回状态说明,10000代表正确
	Forecasts []struct {
		City       string `json:"city"`       // 城市名称
		Adcode     string `json:"adcode"`     // 城市编码
		Province   string `json:"province"`   // 省份名称
		Reporttime string `json:"reporttime"` // 预报发布时间
		Casts      []struct {
			Date         string `json:"date"`         // 日期
			Week         string `json:"week"`         // 星期
			Dayweather   string `json:"dayweather"`   // 白天天气现象
			Nightweather string `json:"nightweather"` // 晚上天气现象
			Daytemp      string `json:"daytemp"`      // 白天温度
			Nighttemp    string `json:"nighttemp"`    // 晚上温度
			Daywind      string `json:"daywind"`      // 白天风向
			Nightwind    string `json:"nightwind"`    // 晚上风向
			Daypower     string `json:"daypower"`     // 白天风力
			Nightpower   string `json:"nightpower"`   // 晚上风力
		} `json:"casts"` // 预报数据list结构，元素cast,按顺序为当天、第二天、第三天的预报数据
	} `json:"forecasts"` // 预报天气信息数据
}
type Fusion struct {
	province   string    // 省份
	city       string    // 市
	reporttime time.Time // 预报发布时间
	Forecasts  []Forecast
}
type Forecast struct {
	date         string
	week         string
	dayweather   string
	nightweather string
	daytemp      string
	nighttemp    string
	daywind      string
	nightwind    string
	daypower     string
	nightpower   string
}

func GetWeather(key, city, extensions, output string) Fusion {
	key = strings.Join([]string{"key", key}, "=")                      // key=<用户key>
	city = strings.Join([]string{"city", city}, "=")                   // city=110101
	extensions = strings.Join([]string{"extensions", extensions}, "=") // extensions=base|all
	output = strings.Join([]string{"output", "JSON"}, "=")             // output=JSON|XML
	args := strings.Join([]string{key, city, extensions, output}, "&")
	prefix := "https://restapi.amap.com/v3/weather/weatherInfo"
	url := strings.Join([]string{prefix, args}, "?")
	log.Debug.Println(url)
	var report Report
	response, _ := http.Get(url)
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	_ = json.Unmarshal(body, &report)
	log.Debug.Println(report)
	// "https://restapi.amap.com/v3/weather/weatherInfo?city=110101&key=<用户key>"
	log.Debug.Printf("%T\n", report.Forecasts[0].Reporttime)
	//loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.Parse("2006-01-02 15:04:05", report.Forecasts[0].Reporttime)
	var forecasts []Forecast
	today_forecast := &Forecast{
		date:         report.Forecasts[0].Casts[0].Date,
		week:         report.Forecasts[0].Casts[0].Week,
		dayweather:   report.Forecasts[0].Casts[0].Dayweather,
		nightweather: report.Forecasts[0].Casts[0].Nightweather,
		daytemp:      report.Forecasts[0].Casts[0].Daytemp,
		nighttemp:    report.Forecasts[0].Casts[0].Nighttemp,
		daywind:      report.Forecasts[0].Casts[0].Daywind,
		nightwind:    report.Forecasts[0].Casts[0].Nightwind,
		daypower:     report.Forecasts[0].Casts[0].Daypower,
		nightpower:   report.Forecasts[0].Casts[0].Nightpower,
	}
	forecasts = append(forecasts, *today_forecast)
	tomorrow_forecast := &Forecast{
		date:         report.Forecasts[0].Casts[1].Date,
		week:         report.Forecasts[0].Casts[1].Week,
		dayweather:   report.Forecasts[0].Casts[1].Dayweather,
		nightweather: report.Forecasts[0].Casts[1].Nightweather,
		daytemp:      report.Forecasts[0].Casts[1].Daytemp,
		nighttemp:    report.Forecasts[0].Casts[1].Nighttemp,
		daywind:      report.Forecasts[0].Casts[1].Daywind,
		nightwind:    report.Forecasts[0].Casts[1].Nightwind,
		daypower:     report.Forecasts[0].Casts[1].Daypower,
		nightpower:   report.Forecasts[0].Casts[1].Nightpower,
	}
	forecasts = append(forecasts, *tomorrow_forecast)
	after_tomorrow_forecast := &Forecast{
		date:         report.Forecasts[0].Casts[2].Date,
		week:         report.Forecasts[0].Casts[2].Week,
		dayweather:   report.Forecasts[0].Casts[2].Dayweather,
		nightweather: report.Forecasts[0].Casts[2].Nightweather,
		daytemp:      report.Forecasts[0].Casts[2].Daytemp,
		nighttemp:    report.Forecasts[0].Casts[2].Nighttemp,
		daywind:      report.Forecasts[0].Casts[2].Daywind,
		nightwind:    report.Forecasts[0].Casts[2].Nightwind,
		daypower:     report.Forecasts[0].Casts[2].Daypower,
		nightpower:   report.Forecasts[0].Casts[2].Nightpower,
	}
	forecasts = append(forecasts, *after_tomorrow_forecast)
	fusion := &Fusion{
		province:   report.Forecasts[0].Province,
		city:       report.Forecasts[0].City,
		reporttime: t,
		Forecasts:  forecasts,
	}
	log.Debug.Printf("%+v\n", fusion)
	return *fusion
}
