package http_client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"

	"github.com/tikivn/tially/pkg/util/logaff"
)

func NewDefaultHttpExecutor() HttpExecutor {
	return &httpExecutor{}
}

type httpExecutor struct{}

// only use if you don't have prometheus or logging
func (c *httpExecutor) Execute(r *http.Request, timeout time.Duration, retry uint64) (int, []byte, error) {
	logger := logaff.GetNewLogger()
	var res []byte
	var statusCode int
	op := func() error {
		client := &http.Client{
			Timeout: timeout,
		}
		httpRes, err := client.Do(r)
		if err != nil {
			return err
		}

		// Close the connection to reuse it (keep-alive connection)
		defer httpRes.Body.Close()

		statusCode = httpRes.StatusCode

		res, err = ioutil.ReadAll(httpRes.Body)
		if httpRes.StatusCode > 299 {
			return fmt.Errorf(string(res))
		}
		return err
	}

	retryFn := backoff.WithContext(backoff.WithMaxRetries(backoff.NewConstantBackOff(50*time.Millisecond), retry), r.Context())
	err := backoff.Retry(op, retryFn)
	if err != nil {
		logger.Errorf("%s, response data: %s", err.Error(), string(res))
		return statusCode, nil, err
	}
	return statusCode, res, nil
}
