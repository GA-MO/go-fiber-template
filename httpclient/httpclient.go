package httpclient

import (
	"time"

	"github.com/go-resty/resty/v2"
)

type HttpClient = *resty.Client

func NewHttpClient() HttpClient {
	client := resty.New()
	client.SetDebug(true)
	client.SetTimeout(10 * time.Second)
	client.SetRetryCount(3)
	client.SetRetryWaitTime(1 * time.Second)
	client.SetRetryMaxWaitTime(5 * time.Second)
	client.SetDebug(true)

	return client
}
