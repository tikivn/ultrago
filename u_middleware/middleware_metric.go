package u_middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/tikivn/ultrago/u_prometheus"
)

func NewMetricMiddleware() *MetricMiddleware {
	return &MetricMiddleware{
		prometheusHttpConfig: u_prometheus.NewDefaultIncomingHttpConfig(),
	}
}

type MetricMiddleware struct {
	prometheusHttpConfig *u_prometheus.HttpConfig
}

func (a *MetricMiddleware) WithPrometheusHttpConfig(conf *u_prometheus.HttpConfig) *MetricMiddleware {
	if conf == nil {
		return a
	}

	a.prometheusHttpConfig.WithHttpConfig(*conf)
	return a
}

func (a *MetricMiddleware) Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			next.ServeHTTP(ww, r)
			if a.isIgnorePath(r.URL.Path) || a.isIgnoreStatus(ww.Status()) || a.isIgnoreMethod(r.Method) {
				return
			}

			u_prometheus.MetricIncomingHttpRequest.
				WithLabelValues(fmt.Sprintf("%d", ww.Status()), r.Method, a.cleanUpPath(r.URL.Path)).
				Observe(time.Since(start).Seconds())
		}
		return http.HandlerFunc(fn)
	}
}

func (a *MetricMiddleware) cleanUpPath(path string) string {
	for regex, alt := range a.prometheusHttpConfig.PathCleanUpMap {
		path = regex.ReplaceAllString(path, alt)
	}
	return path
}

func (a *MetricMiddleware) isIgnorePath(path string) bool {
	if a.prometheusHttpConfig == nil {
		return false
	}

	_, ok := a.prometheusHttpConfig.PathIgnoredMap[path]
	return ok
}

func (a *MetricMiddleware) isIgnoreStatus(status int) bool {
	if a.prometheusHttpConfig == nil {
		return false
	}

	_, ok := a.prometheusHttpConfig.StatusIgnoredMap[status]
	return ok
}

func (a *MetricMiddleware) isIgnoreMethod(method string) bool {
	if a.prometheusHttpConfig == nil {
		return false
	}

	_, ok := a.prometheusHttpConfig.MethodIgnoreMap[method]
	return ok
}
