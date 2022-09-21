# Weather
使用高德API获取天气
# example

```go
func TestGetWeather(t *testing.T) {
words := []string{}
key := ""
city := "110107"
extensions := "all"
output := "JSON"
result := GetWeather(key, city, extensions, output)
date := result.reporttime.Format("2006年1月2日 15时04分05发布")
title := strings.Join([]string{result.province, result.city, date}, "")
t.Log(title)
_date := result.Forecasts[0].date
today_date := strings.Join([]string{strings.Split(_date, "-")[0], "年", strings.Split(_date, "-")[1], "月", strings.Split(_date, "-")[2], "日"}, "")
today_week := constant.Week[result.Forecasts[0].week]
today_day_weather := result.Forecasts[0].dayweather
today_night_weather := result.Forecasts[0].nightweather
today_day_temp := result.Forecasts[0].daytemp
today_night_temp := result.Forecasts[0].nighttemp
today_day_wind := result.Forecasts[0].daywind
today_night_wind := result.Forecasts[0].nightwind
today_day_power := result.Forecasts[0].daypower
today_day_power = strings.Replace(today_day_power, "\u2264", "小于", -1)
today_day_power = strings.Replace(today_day_power, "\u2267", "大于", -1)
today_night_power := result.Forecasts[0].nightpower
today_night_power = strings.Replace(today_night_power, "\u2264", "小于", -1)
today_night_power = strings.Replace(today_night_power, "\u2267", "大于", -1)
today_day_word := strings.Join([]string{today_date, today_week, "白天天气", today_day_weather, today_day_temp, "\u2103", today_day_wind, today_day_power, "级"}, "")
today_night_word := strings.Join([]string{"夜间天气", today_night_weather, today_night_temp, "\u2103", today_night_wind, today_night_power, "级"}, "")
words = append(words, today_day_word, today_night_word)
t.Log(words)
}
```