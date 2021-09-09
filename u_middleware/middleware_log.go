package u_middleware

import (
	"net/http"

	"github.com/tikivn/ultrago/u_logger"
)

func NewLogMiddleware() *LogMiddleware {
	return &LogMiddleware{}
}

type LogMiddleware struct {
}

func (a *LogMiddleware) Middleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx, _ := u_logger.GetLogger(r.Context())
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
