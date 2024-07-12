package client

import (
	"context"
	"io"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type HTTPClient struct {
	client *http.Client
}

type HTTPClientConfig struct {
	Timeout             time.Duration
	MaxIdleConns        int
	MaxConnsPerHost     int
	MaxIdleConnsPerHost int
}

func NewHTTPClient(config *HTTPClientConfig) *HTTPClient {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = config.MaxIdleConns
	t.MaxConnsPerHost = config.MaxConnsPerHost
	t.MaxIdleConnsPerHost = config.MaxIdleConnsPerHost

	return &HTTPClient{
		client: &http.Client{
			Timeout:   config.Timeout,
			Transport: t,
		},
	}
}

func NewDefaultHTTPClient() *HTTPClient {
	return NewHTTPClient(&HTTPClientConfig{
		Timeout:             10 * time.Second,
		MaxIdleConns:        100,
		MaxConnsPerHost:     100,
		MaxIdleConnsPerHost: 100,
	})
}

func (c *HTTPClient) Get(ctx context.Context, url string) ([]byte, int, error) {
	resp, err := otelhttp.Get(ctx, url)
	if err != nil {
		return nil, 0, err
	}

	bytes, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return bytes, resp.StatusCode, nil
}

func (c *HTTPClient) GetWithTrace(ctx context.Context, url string) ([]byte, int, error) {
	// ctx = httptrace.WithClientTrace(ctx, otelhttptrace.NewClientTrace(ctx))
	resp, err := otelhttp.Get(ctx, url)
	if err != nil {
		return nil, 0, err
	}

	bytes, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return bytes, resp.StatusCode, nil
}
