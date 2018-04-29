package handler

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
)

// HandleNotFound is the generator of the NotFound handler.
func HandleNotFound(lgc *logic.Logic) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lgc.Logger().WithFields(log.Fields{
			"path": r.URL.Path, "method": r.Method,
			"message": "404 Page Not Found",
		}).Info("request start")
		w.WriteHeader(404)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf(
			`{"message":"Page Not Found %s %s"}`, r.Method, r.URL.Path)))
	}
}
