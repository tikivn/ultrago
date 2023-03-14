package u_http_client

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tikivn/ultrago/u_env"
	"github.com/tikivn/ultrago/u_logger"
	"moul.io/http2curl"
)

func NewHttpClient(httpExecutor HttpExecutor, timeout time.Duration) *HttpClient {
	return &HttpClient{
		retry:        0,
		timeout:      timeout,
		httpExecutor: httpExecutor,
	}
}

func NewRetryHttpClient(httpExecutor HttpExecutor, timeout time.Duration, retry uint64) *HttpClient {
	return &HttpClient{
		retry:        retry,
		timeout:      timeout,
		httpExecutor: httpExecutor,
	}
}

type HttpClient struct {
	url          string
	headers      map[string]string
	params       interface{}
	retry        uint64
	timeout      time.Duration
	httpExecutor HttpExecutor
}

func (c *HttpClient) URL() string {
	return c.url
}

func (c *HttpClient) WithUrl(uri string, params map[string][]string) *HttpClient {
	if params != nil {
		requestParams := url.Values{}
		for key, values := range params {
			for _, value := range values {
				if value != "" {
					requestParams.Add(key, value)
				}
			}
		}
		c.url = fmt.Sprintf("%s?%s", uri, requestParams.Encode())
	} else {
		c.url = uri
	}
	return c
}

func (c *HttpClient) WithHeaders(headers map[string]string) *HttpClient {
	if c.headers == nil {
		c.headers = make(map[string]string, 0)
	}
	for key, value := range headers {
		c.headers[key] = value
	}
	return c
}

func (c *HttpClient) WithPayload(params interface{}) *HttpClient {
	c.params = params
	return c
}

func (c *HttpClient) WithBasicAuth(username string, password string) *HttpClient {
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))
	return c.WithHeaders(map[string]string{"Authorization": auth})
}

func (c *HttpClient) WithBearerAuth(token string, addPrefix bool) *HttpClient {
	var auth string
	if addPrefix {
		auth = "Bearer " + token
	} else {
		auth = token
	}
	return c.WithHeaders(map[string]string{"Authorization": auth})
}

func (c *HttpClient) Do(ctx context.Context, method string) (*HttpResponse, error) {
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

	statusCode, resp, err := c.httpExecutor.Execute(r, c.timeout, c.retry)

	logger.WithFields(logrus.Fields{
		"url":     c.url,
		"payload": c.params,
		"header":  c.headers,
	}).Infof("call api with code=%d res=%s, err=%v", statusCode, string(resp), err)
	return &HttpResponse{
		Code:    statusCode,
		Payload: resp,
	}, err
}

func (c *HttpClient) DoFormEncoding(ctx context.Context, method string) (*HttpResponse, error) {
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

	statusCode, resp, err := c.httpExecutor.Execute(r, c.timeout, c.retry)

	logger.WithFields(logrus.Fields{
		"url":     c.url,
		"payload": c.params,
		"header":  c.headers,
	}).Infof("call api with code=%d res=%s, err=%v", statusCode, string(resp), err)
	return &HttpResponse{
		Code:    statusCode,
		Payload: resp,
	}, err
}

func (c *HttpClient) DoFormMultipart(ctx context.Context) (*HttpResponse, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url, c.params.(io.Reader))
	if err != nil {
		return nil, err
	}
	for key, value := range c.headers {
		r.Header.Set(key, value)
	}

	statusCode, resp, err := c.httpExecutor.Execute(r, c.timeout, c.retry)
	// comment because only use for upload image to cdn, will consider in the feature
	return &HttpResponse{
		Code:    statusCode,
		Payload: resp,
	}, err
}
