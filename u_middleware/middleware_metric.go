package u_middleware

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/tikivn/ultrago/u_prometheus"
)

func NewMetricMiddleware(pathConfig PathConfig) *MetricMiddleware {
	if pathConfig == nil {
		pathConfig = NewDefaultPathConfig()
	}

	return &MetricMiddleware{
		pathCleanUpMap: pathConfig.PathCleanUp(),
		pathIgnoredMap: pathConfig.PathIgnored(),
	}
}

type MetricMiddleware struct {
	pathCleanUpMap map[*regexp.Regexp]string
	pathIgnoredMap map[string]bool
}

func (a *MetricMiddleware) Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			next.ServeHTTP(ww, r)
			if !a.isIgnorePath(r.URL.Path) {
				u_prometheus.MetricIncomingHttpRequest.
					WithLabelValues(fmt.Sprintf("%d", ww.Status()), r.Method, a.cleanUpPath(r.URL.Path)).
					Observe(time.Since(start).Seconds())
			}
		}
		return http.HandlerFunc(fn)
	}
}

func (a *MetricMiddleware) cleanUpPath(path string) string {
	for regex, alt := range a.pathCleanUpMap {
		path = regex.ReplaceAllString(path, alt)
	}
	return path
}

func (a *MetricMiddleware) isIgnorePath(path string) bool {
	_, ok := a.pathIgnoredMap[path]
	return ok
}
