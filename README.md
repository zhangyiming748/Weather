# Weather
使用高德API获取天气
# example

```go
func TestGetWeather(t *testing.T) {
	key := ""
	city := "110107"
	extensions := "all"
	output := "JSON"
	result := GetWeather(key, city, extensions, output)
	date := result.reporttime.Format("2006年1月2日 15时04分05发布")
	title := strings.Join([]string{result.province, result.city, date}, "")
	t.Log(title)
}
```