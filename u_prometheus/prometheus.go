package u_prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tikivn/ultrago/u_logger"
)

const (
	metric_incoming_http_request = "uo_incoming_http_request"
	metric_outgoing_http_request = "uo_outgoing_http_request"
)

var (
	MetricIncomingHttpRequest *prometheus.HistogramVec
	MetricOutgoingHttpRequest *prometheus.HistogramVec
)

func init() {
	MetricIncomingHttpRequest = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    metric_incoming_http_request,
		Help:    "How long it took to process the request, partitioned by status code, method and HTTP path.",
		Buckets: []float64{0.01, 0.03, 0.05, 0.1, 0.3, 0.5, 1, 3, 5, 10},
	}, []string{"code", "method", "path"})

	MetricOutgoingHttpRequest = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    metric_outgoing_http_request,
		Help:    "How long it took to make a request and has response, partitioned by status code, method and HTTP path.",
		Buckets: []float64{0.01, 0.03, 0.05, 0.1, 0.3, 0.5, 1, 3, 5, 10},
	}, []string{"code", "method", "path", "host"})

	defer func() {
		if re := recover(); re != nil {
			logger := u_logger.NewLogger()
			logger.Errorf("register prometheus metric failed: %v", re)
		}
	}()
	prometheus.MustRegister(MetricIncomingHttpRequest, MetricOutgoingHttpRequest)
}
