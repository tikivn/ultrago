package u_http_client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/tikivn/ultrago/u_logger"
	"github.com/tikivn/ultrago/u_prometheus"
)

func NewHttpExecutor() HttpExecutor {
	return &httpExecutor{
		prometheusHttpConfig: u_prometheus.NewDefaultOutgoingHttpConfig(),
	}
}

type httpExecutor struct {
	prometheusHttpConfig *u_prometheus.HttpConfig
}

func (a *httpExecutor) WithPrometheusHttpConfig(conf *u_prometheus.HttpConfig) HttpExecutor {
	if conf == nil {
		return a
	}

	a.prometheusHttpConfig.WithHttpConfig(*conf)
	return a
}

func (c *httpExecutor) Execute(r *http.Request, timeout time.Duration, retry uint64) (int, []byte, error) {
	ctx, logger := u_logger.GetLogger(r.Context())

	var res []byte
	var statusCode int
	op := func() error {
		client := &http.Client{
			Timeout: timeout,
		}

		start := time.Now()
		httpRes, err := client.Do(r)
		if err != nil {
			return err
		}

		// Close the connection to reuse it (keep-alive connection)
		defer func() {
			statusCode = httpRes.StatusCode
			httpRes.Body.Close()
		}()

		u_prometheus.MetricOutgoingHttpRequest.
			WithLabelValues(fmt.Sprintf("%d", httpRes.StatusCode), r.Method, c.cleanUpPath(r.URL.Path), r.URL.Host).
			Observe(time.Since(start).Seconds())

		res, err = ioutil.ReadAll(httpRes.Body)
		if httpRes.StatusCode > 299 {
			return fmt.Errorf(string(res))
		}
		return err
	}

	retryFn := backoff.WithContext(backoff.WithMaxRetries(backoff.NewConstantBackOff(50*time.Millisecond), retry), ctx)
	err := backoff.Retry(op, retryFn)
	if err != nil {
		logger.Errorf("%v, response data: %s", err, string(res))
		return statusCode, nil, err
	}
	return statusCode, res, nil
}

func (c *httpExecutor) cleanUpPath(path string) string {
	if c.prometheusHttpConfig == nil {
		return path
	}

	for regex, alt := range c.prometheusHttpConfig.PathCleanUpMap {
		path = regex.ReplaceAllString(path, alt)
	}
	return path
}
