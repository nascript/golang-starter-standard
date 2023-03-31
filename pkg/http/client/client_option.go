package client

import (
	"context"
	"time"
)

type HTTPClientOption func(*HTTPClient)

func Ctx(ctx context.Context) HTTPClientOption {
	return func(httpclient *HTTPClient) {
		httpclient.Ctx = ctx
	}
}

func Endpoint(endpoint string) HTTPClientOption {
	return func(httpclient *HTTPClient) {
		httpclient.Endpoint = endpoint
	}
}

func Timeout(timeout time.Duration) HTTPClientOption {
	return func(httpclient *HTTPClient) {
		httpclient.Timeout = timeout
	}
}

func Method(method string) HTTPClientOption {
	return func(httpclient *HTTPClient) {
		httpclient.Method = method
	}
}
