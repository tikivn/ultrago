package http_client

import (
	"net/http"
	"time"
)

type HttpExecutor interface {
	Execute(r *http.Request, timeout time.Duration, retry uint64) (int, []byte, error)
}
