package u_handler

import (
	"net/http"
)

const (
	LimitKey  string = "limit"
	OffsetKey string = "offset"
	SortByKey string = "sort_by"
	OrderKey  string = "order"
)

type LogConfig interface {
	StatusConfig() map[int]bool
}

type RequestUtils interface {
	RequestParamStr(r *http.Request, key string) string
	RequestParamStrWithDefault(r *http.Request, key string, defaultValue string) string
	RequestParamInt(r *http.Request, key string, defaultValue int64) int64
	RequestParamFloat(r *http.Request, key string, defaultValue float64) float64
	RequestParamBool(r *http.Request, key string, defaultValue bool) bool
	RequestParamArray(r *http.Request, key string) []string
}

type ResponseUtils interface {
	BadRequest(w http.ResponseWriter, r *http.Request, err error)
	Unauthorized(w http.ResponseWriter, r *http.Request, err error)
	Forbidden(w http.ResponseWriter, r *http.Request, err error)
	NotFound(w http.ResponseWriter, r *http.Request, err error)
	TooManyRequests(w http.ResponseWriter, r *http.Request, err error)
	Internal(w http.ResponseWriter, r *http.Request, err error)
	Success(w http.ResponseWriter, r *http.Request, data interface{})
	FileSuccess(w http.ResponseWriter, r *http.Request, data interface{}, fileName string)
	PaginateSuccess(w http.ResponseWriter, r *http.Request, data interface{}, total int64)
}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{
		logConfig: map[int]bool{
			http.StatusBadRequest:          true,
			http.StatusForbidden:           true,
			http.StatusTooManyRequests:     true,
			http.StatusInternalServerError: true,
		},
	}
}

type BaseHandler struct {
	logConfig map[int]bool
}

func (h *BaseHandler) WithLogConfig(conf LogConfig) *BaseHandler {
	if conf != nil {
		logConfig := conf.StatusConfig()
		if len(logConfig) > 0 {
			h.logConfig = logConfig
		}
	}
	return h
}
