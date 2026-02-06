package constants

import "net/http"

var DefaultStatsHeaders = http.Header{
	"Host":             {"stats.nba.com"},
	"User-Agent":       {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36"},
	"Accept":           {"application/json, text/plain, */*"},
	"Accept-Language":  {"en-US,en;q=0.5"},
	"Accept-Encoding":  {"gzip, deflate, br"},
	"Connection":       {"keep-alive"},
	"Referer":          {"https://stats.nba.com/"},
	"Pragma":           {"no-cache"},
	"Cache-Control":    {"no-cache"},
	"Sec-Ch-Ua":        {"\"Chromium\";v=\"140\", \"Google Chrome\";v=\"140\", \"Not;A=Brand\";v=\"99\""},
	"Sec-Ch-Ua-Mobile": {"?0"},
	"Sec-Fetch-Dest":   {"empty"},
}

var DefaultLiveStatsHeaders = http.Header{
	"Accept":          {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"},
	"Accept-Encoding": {"gzip, deflate, br"},
	"Accept-Language": {"en-US,en;q=0.9"},
	"Cache-Control":   {"max-age=0"},
	"Connection":      {"keep-alive"},
	"Host":            {"cdn.nba.com"},
	"User-Agent":      {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"},
}
