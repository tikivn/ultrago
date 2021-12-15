package u_handler

import (
	"fmt"
	"net/http"
)

type baseProcessor struct{}

func (baseProcessor) BeforeSuccess(w http.ResponseWriter, r *http.Request, data interface{}) {
	return
}

func (baseProcessor) BeforeError(w http.ResponseWriter, r *http.Request, err error) {
	return
}

func (baseProcessor) FormatErr(w http.ResponseWriter, r *http.Request, err error) string {
	return fmt.Sprintf("%v", err)
}
