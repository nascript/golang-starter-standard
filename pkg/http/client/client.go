package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type (
	HTTPClient struct {
		Ctx      context.Context
		Endpoint string
		Timeout  time.Duration
		Method   string
	}

	IHttpClient interface {
		NewClient(opts ...HTTPClientOption) *HTTPClient
		MakeRequest(obj interface{}) error
	}
)

func (c *HTTPClient) NewClient(opts ...HTTPClientOption) *HTTPClient {
	httpclient := &HTTPClient{
		Ctx:     context.Background(),
		Timeout: 10 * time.Second,
		Method:  http.MethodGet,
	}

	for _, opt := range opts {
		opt(httpclient)
	}

	return httpclient
}

func (c *HTTPClient) MakeRequest(obj interface{}, body io.Reader) error {
	req, err := http.NewRequestWithContext(c.Ctx, c.Method, c.Endpoint, body)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: c.Timeout}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) { _ = Body.Close() }(resp.Body)

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return err
	}
	data := buf.Bytes()

	return json.Unmarshal(data, &obj)
}
