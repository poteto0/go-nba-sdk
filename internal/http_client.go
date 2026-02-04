package internal

import (
	"net/http"
	"time"
)

var DefaultHeaders = http.Header{
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

type HttpClient struct {
	client        *http.Client
	defaultHeader http.Header
}

type IHttpClient interface {
	DefaultHeader() http.Header
	Get(path string, headers *http.Header) (*http.Response, error)
}

func NewHttpClient() IHttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		defaultHeader: DefaultHeaders,
	}
}

func (c *HttpClient) DefaultHeader() http.Header {
	return c.defaultHeader
}

func (c *HttpClient) Get(path string, headers *http.Header) (*http.Response, error) {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		req.Header = *headers
	} else {
		req.Header = c.defaultHeader
	}

	return c.client.Do(req)
}
