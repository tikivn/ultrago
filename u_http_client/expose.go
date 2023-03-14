package u_http_client

import (
	"net/http"
	"time"

	"github.com/tikivn/ultrago/u_prometheus"
)

type HttpExecutor interface {
	Execute(r *http.Request, timeout time.Duration, retry uint64) (int, []byte, error)
	WithPrometheusHttpConfig(conf *u_prometheus.HttpConfig) HttpExecutor
}

type HttpResponse struct {
	Code    int
	Payload []byte
}
