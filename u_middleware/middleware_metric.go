package u_middleware

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/tikivn/ultrago/u_prometheus"
)

func NewMetricMiddleware() *MetricMiddleware {
	return &MetricMiddleware{
		pathCleanUpMap: map[*regexp.Regexp]string{
			regexp.MustCompile("\\/([0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12})\\/"): "/<id>/",
			regexp.MustCompile("\\/([0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12})"):    "/<id>",
		},
		pathIgnoredMap: map[string]bool{
			"/":            true,
			"/healthcheck": true,
			"/heartbeat":   true,
			"/metrics":     true,
		},
	}
}

type MetricMiddleware struct {
	pathCleanUpMap map[*regexp.Regexp]string
	pathIgnoredMap map[string]bool
}

func (a *MetricMiddleware) WithPathConfig(conf PathConfig) *MetricMiddleware {
	if conf != nil {
		pathCleanUp := conf.PathCleanUp()
		if len(pathCleanUp) > 0 {
			a.pathCleanUpMap = pathCleanUp
		}

		pathIgnored := conf.PathIgnored()
		if len(pathIgnored) > 0 {
			a.pathIgnoredMap = pathIgnored
		}
	}
	return a
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
