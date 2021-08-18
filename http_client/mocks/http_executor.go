package mocks

import (
	"net/http"
	"time"

	"github.com/tikivn/tially/pkg/util/logaff"
	"moul.io/http2curl"

	"github.com/tikivn/tially/pkg/util/http_client"
)

func NewHttpExecutor() http_client.HttpExecutor {
	return &mockHttpExecutor{
		res: make(map[string][]byte, 0),
		err: make(map[string]error, 0),
	}
}

func NewHttpExecutorWithRes(mockRes map[string][]byte, mockErr map[string]error) http_client.HttpExecutor {
	executor := &mockHttpExecutor{
		res: make(map[string][]byte, 0),
		err: make(map[string]error, 0),
	}
	if mockRes != nil {
		executor.res = mockRes
	}
	if mockErr != nil {
		executor.err = mockErr
	}
	return executor
}

type mockHttpExecutor struct {
	res map[string][]byte
	err map[string]error
}

func (svc *mockHttpExecutor) Execute(r *http.Request, timeout time.Duration, retry uint64) (int, []byte, error) {
	_, logger := logaff.GetLogger(r.Context())
	command, err := http2curl.GetCurlCommand(r)
	if err != nil {
		logger.Errorf("build curl command failed: %v", err)
	} else {
		logger.Infof("execute request with timeout = %s and retry = %d with curl:\n%s", timeout.String(), retry, command)
	}
	return http.StatusOK, svc.res[r.URL.Path], svc.err[r.URL.Path]
}
