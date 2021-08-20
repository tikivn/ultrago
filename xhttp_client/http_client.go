package xhttp_client

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"

	"github.com/tikivn/ultrago/xlogaff"
)

func NewHttpClient(timeout time.Duration) *HttpClient {
	return &HttpClient{
		retry: 0,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func NewRetryHttpClient(timeout time.Duration, retry uint64) *HttpClient {
	return &HttpClient{
		retry: retry,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

type HttpClient struct {
	url     string
	headers map[string]string
	params  interface{}
	retry   uint64

	client *http.Client
}

func (c *HttpClient) WithUrl(uri string, params map[string]string) *HttpClient {
	if params != nil {
		requestParams := url.Values{}
		for key, value := range params {
			requestParams.Set(key, value)
		}
		c.url = fmt.Sprintf("%s?%s", uri, requestParams.Encode())
	} else {
		c.url = uri
	}
	return c
}

func (c *HttpClient) WithHeaders(headers map[string]string) *HttpClient {
	if c.headers == nil {
		c.headers = make(map[string]string)
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

func (c *HttpClient) WithBearerAuth(token string) *HttpClient {
	auth := "Bearer " + token
	return c.WithHeaders(map[string]string{"Authorization": auth})
}

func (c *HttpClient) Do(ctx context.Context, method string) ([]byte, error) {
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
	for key, value := range c.headers {
		r.Header.Set(key, value)
	}

	return c.execute(r)
}

func (c *HttpClient) DoFormEncoding(ctx context.Context, method string) ([]byte, error) {
	var buffer *strings.Reader
	if c.params != nil {
		buffer = strings.NewReader(c.params.(string))
	}

	r, err := http.NewRequestWithContext(ctx, method, c.url, buffer)
	if err != nil {
		return nil, err
	}
	for key, value := range c.headers {
		r.Header.Set(key, value)
	}

	return c.execute(r)
}

func (c *HttpClient) DoFormMultipart(ctx context.Context) ([]byte, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url, c.params.(io.Reader))
	if err != nil {
		return nil, err
	}
	for key, value := range c.headers {
		r.Header.Set(key, value)
	}

	return c.execute(r)
}

func (c *HttpClient) execute(r *http.Request) ([]byte, error) {
	var res []byte
	op := func() error {
		httpRes, err := c.client.Do(r)
		if err != nil {
			return err
		}
		res, err = ioutil.ReadAll(httpRes.Body)
		if httpRes.StatusCode > 299 {
			return fmt.Errorf("non 2xx status code return: code %v", httpRes.StatusCode)
		}
		return err
	}

	retry := backoff.WithContext(backoff.WithMaxRetries(backoff.NewConstantBackOff(50*time.Millisecond), c.retry), r.Context())
	err := backoff.Retry(op, retry)
	if err != nil {
		logger := xlogaff.GetNewLogger()
		logger.Errorf("%s, response data: %s", err.Error(), string(res))
		return nil, err
	}
	return res, nil
}
