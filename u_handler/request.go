package u_handler

import (
	"net/http"
	"strconv"
	"strings"
)

func (h *BaseHandler) RequestParamStr(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

func (h *BaseHandler) RequestParamStrWithDefault(r *http.Request, key string, defaultValue string) string {
	value := r.URL.Query().Get(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func (h *BaseHandler) RequestParamInt(r *http.Request, key string, defaultValue int64) int64 {
	value := r.URL.Query().Get(key)
	valueInt, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defaultValue
	}
	return valueInt
}

func (h *BaseHandler) RequestParamFloat(r *http.Request, key string, defaultValue float64) float64 {
	value := r.URL.Query().Get(key)
	valueInt, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defaultValue
	}
	return valueInt
}

func (h *BaseHandler) RequestParamBool(r *http.Request, key string, defaultValue bool) bool {
	value := r.URL.Query().Get(key)
	valueBool, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return valueBool
}

func (h *BaseHandler) RequestParamArray(r *http.Request, key string) []string {
	value := r.URL.Query().Get(key)
	if value == "" {
		return []string{}
	} else {
		return strings.Split(value, ",")
	}
}
