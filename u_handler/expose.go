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

type Processor interface {
	BeforeSuccess(w http.ResponseWriter, r *http.Request, data interface{})
	BeforeError(w http.ResponseWriter, r *http.Request, err error)
	FormatErr(w http.ResponseWriter, r *http.Request, err error) string
}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{
		processor: new(baseProcessor),
		logConfig: map[int]bool{
			http.StatusBadRequest:          true,
			http.StatusForbidden:           true,
			http.StatusTooManyRequests:     true,
			http.StatusInternalServerError: true,
		},
	}
}

type BaseHandler struct {
	processor Processor
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

func (h *BaseHandler) WithProcessor(proc Processor) *BaseHandler {
	if proc != nil {
		h.processor = proc
	}
	return h
}
