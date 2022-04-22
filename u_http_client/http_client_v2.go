package u_http_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tikivn/ultrago/u_env"
	"github.com/tikivn/ultrago/u_logger"
	"moul.io/http2curl"
)

type BeforeRetryFn func(ctx context.Context) (headers map[string]string)

func NewHttpClientV2(httpExecutor HttpExecutor, timeout time.Duration) *HttpClientV2 {
	return &HttpClientV2{
		HttpClient: HttpClient{
			retry:        0,
			timeout:      timeout,
			httpExecutor: httpExecutor,
		},
	}
}

func NewRetryHttpClientV2(httpExecutor HttpExecutor, timeout time.Duration, retry uint64) *HttpClientV2 {
	return &HttpClientV2{
		HttpClient: HttpClient{
			retry:        retry,
			timeout:      timeout,
			httpExecutor: httpExecutor,
		},
	}
}

type HttpClientV2 struct {
	HttpClient
	retryFnMap map[int]BeforeRetryFn
}

func (c *HttpClientV2) WithRetryFn(statusCode int, fn BeforeRetryFn) *HttpClientV2 {
	if c.retryFnMap == nil {
		c.retryFnMap = make(map[int]BeforeRetryFn, 0)
	}
	c.retryFnMap[statusCode] = fn
	return c
}

func (c *HttpClientV2) Do(ctx context.Context, method string) ([]byte, error) {
	ctx, logger := u_logger.GetLogger(ctx)

	var buffer *bytes.Buffer
	if c.params != nil {
		payload, err := json.Marshal(c.params)
		if err != nil {
			return nil, err
		}
		buffer = bytes.NewBuffer(payload)
	} else {
		// prevent nil pointer bug when request lib dereference *bytes.Buffer in body request
		buffer = bytes.NewBuffer(nil)
	}

	for i := uint64(0); i <= c.retry; i++ {
		r, err := http.NewRequestWithContext(ctx, method, c.url, buffer)
		if err != nil {
			return nil, err
		}
		if c.headers == nil {
			c.headers = make(map[string]string, 0)
		}
		c.headers["Content-Type"] = "application/json"
		for key, value := range c.headers {
			r.Header.Set(key, value)
		}

		if u_env.IsDev() {
			command, _ := http2curl.GetCurlCommand(r)
			logger.Info(command.String())
		}

		statusCode, resp, err := c.httpExecutor.Execute(r, c.timeout, 0)
		logger.WithFields(logrus.Fields{
			"url":     c.url,
			"payload": c.params,
			"header":  c.headers,
		}).Infof("call api with code=%d res=%s, err=%v", statusCode, string(resp), err)

		if err == nil {
			return resp, nil
		} else if i == c.retry {
			return resp, err
		} else {
			time.Sleep(50 * time.Millisecond)
			fn, ok := c.retryFnMap[statusCode]
			if ok {
				headers := fn(ctx)
				if len(c.headers) > 0 {
					c.WithHeaders(headers)
				}
			}
		}
	}

	return nil, fmt.Errorf("this should not happen")
}

func (c *HttpClientV2) DoFormEncoding(ctx context.Context, method string) ([]byte, error) {
	ctx, logger := u_logger.GetLogger(ctx)

	var buffer *strings.Reader
	if c.params != nil {
		switch c.params.(type) {
		case map[string]string:
			payload := url.Values{}
			for key, value := range c.params.(map[string]string) {
				payload.Set(key, value)
			}
			buffer = strings.NewReader(payload.Encode())
		case url.Values:
			buffer = strings.NewReader(c.params.(url.Values).Encode())
		case string:
			buffer = strings.NewReader(c.params.(string))
		default:
			return nil, fmt.Errorf("invalid payload field type")
		}
	}

	for i := uint64(0); i <= c.retry; i++ {
		r, err := http.NewRequestWithContext(ctx, method, c.url, buffer)
		if err != nil {
			return nil, err
		}
		if c.headers == nil {
			c.headers = make(map[string]string, 0)
		}
		c.headers["Content-Type"] = "application/x-www-form-urlencoded"
		for key, value := range c.headers {
			r.Header.Set(key, value)
		}

		if u_env.IsDev() {
			command, _ := http2curl.GetCurlCommand(r)
			logger.Info(command.String())
		}

		statusCode, resp, err := c.httpExecutor.Execute(r, c.timeout, 0)
		logger.WithFields(logrus.Fields{
			"url":     c.url,
			"payload": c.params,
			"header":  c.headers,
		}).Infof("call api with code=%d res=%s, err=%v", statusCode, string(resp), err)
		if err == nil {
			return resp, nil
		} else if i == c.retry {
			return resp, err
		} else {
			time.Sleep(50 * time.Millisecond)
			fn, ok := c.retryFnMap[statusCode]
			if ok {
				headers := fn(ctx)
				if len(c.headers) > 0 {
					c.WithHeaders(headers)
				}
			}
		}
	}

	return nil, fmt.Errorf("this should not happen")
}
