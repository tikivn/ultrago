package u_handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/tikivn/ultrago/u_logger"
)

func (h *BaseHandler) BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	h.logging(r, http.StatusBadRequest)
	h.errorHandler(w, http.StatusBadRequest, err)
}

func (h *BaseHandler) Unauthorized(w http.ResponseWriter, r *http.Request, err error) {
	h.logging(r, http.StatusUnauthorized)
	h.errorHandler(w, http.StatusUnauthorized, err)
}

func (h *BaseHandler) Forbidden(w http.ResponseWriter, r *http.Request, err error) {
	h.logging(r, http.StatusForbidden)
	h.errorHandler(w, http.StatusForbidden, err)
}

func (h *BaseHandler) NotFound(w http.ResponseWriter, r *http.Request, err error) {
	h.logging(r, http.StatusNotFound)
	h.errorHandler(w, http.StatusNotFound, err)
}

func (h *BaseHandler) TooManyRequests(w http.ResponseWriter, r *http.Request, err error) {
	h.logging(r, http.StatusTooManyRequests)
	h.errorHandler(w, http.StatusTooManyRequests, err)
}

func (h *BaseHandler) Internal(w http.ResponseWriter, r *http.Request, err error) {
	h.logging(r, http.StatusInternalServerError)
	h.errorHandler(w, http.StatusInternalServerError, err)
}

func (h *BaseHandler) Success(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":   nil,
		"data":    data,
		"success": true,
	})
}

func (h *BaseHandler) FileSuccess(w http.ResponseWriter, r *http.Request, data interface{}, fileName string) {
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	byteData, _ := json.Marshal(data)
	w.Write(byteData)
}

func (h *BaseHandler) PaginateSuccess(w http.ResponseWriter, r *http.Request, data interface{}, total int64) {
	offset := h.RequestParamInt(r, OffsetKey, 0)
	limit := h.RequestParamInt(r, LimitKey, 10)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":   nil,
		"data":    data,
		"success": true,
		"pagination": map[string]interface{}{
			"total":  total,
			"offset": offset,
			"limit":  limit,
		},
	})
}

func (h *BaseHandler) logging(r *http.Request, statusCode int) {
	_, ok := h.logConfig[statusCode]
	if ok {
		_, logger := u_logger.GetLogger(r.Context())
		logger.WithFields(logrus.Fields{"url": r.URL.Path}).
			Errorf("got code=%d", statusCode)
	}
}

func (h *BaseHandler) errorHandler(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":   fmt.Sprintf("%v", err),
		"data":    nil,
		"success": false,
	})
}
