package u_middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tikivn/ultrago/u_logger"
)

func NewHttpLogMiddleware() *HttpLogMiddleware {
	return &HttpLogMiddleware{}
}

type HttpLogMiddleware struct{}

func (a *HttpLogMiddleware) Middleware() func(next http.Handler) http.Handler {
	return chi.Chain(
		middleware.RequestID,
		Handler(),
		middleware.Recoverer,
	).Handler
}

func Handler() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// Setup request logger with formatter
			ctx, logger := u_logger.GetLogger(r.Context())
			logger.Infof("Request: %s %s", r.Method, r.URL.Path)

			logFields := requestLogFields(r, DefaultOptions.Concise, DefaultOptions.Body)
			for field, value := range logFields {
				logger.Infof("%v: %v", field, value)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func requestLogFields(r *http.Request, concise bool, body bool) map[string]interface{} {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	requestURL := fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)

	requestFields := map[string]interface{}{
		"requestURL":    requestURL,
		"requestMethod": r.Method,
		"requestPath":   r.URL.Path,
		"remoteIP":      r.RemoteAddr,
		"proto":         r.Proto,
	}

	if body {
		// temporary buffer
		b := bytes.NewBuffer(make([]byte, 0))

		// TeeReader returns a Reader that writes to b what it reads from r.Body.
		reader := io.TeeReader(r.Body, b)

		bytes, err := io.ReadAll(reader)
		defer r.Body.Close()
		if err != nil {
			panic(err)
		}
		requestFields["body"] = string(bytes)
		r.Body = io.NopCloser(b)
	}

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		requestFields["requestID"] = reqID
	}

	if concise {
		return map[string]interface{}{
			"httpRequest": requestFields,
		}
	}

	requestFields["scheme"] = scheme
	if len(r.Header) > 0 {
		requestFields["header"] = headerLogField(r.Header)
	}

	return map[string]interface{}{
		"HttpRequest": requestFields,
	}
}

func headerLogField(header http.Header) map[string]string {
	headerField := map[string]string{}
	for k, v := range header {
		k = strings.ToLower(k)
		switch {
		case len(v) == 0:
			continue
		case len(v) == 1:
			headerField[k] = v[0]
		default:
			headerField[k] = fmt.Sprintf("[%s]", strings.Join(v, "], ["))
		}
		if k == "authorization" || k == "cookie" || k == "set-cookie" {
			headerField[k] = "***"
		}

		for _, skip := range DefaultOptions.SkipHeaders {
			if k == skip {
				headerField[k] = "***"
				break
			}
		}
	}
	return headerField
}
