package handler

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/server"
)

func HandleNotFound(ms *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ms.Logger().WithFields(log.Fields{
			"path": r.URL.Path, "method": r.Method,
			"message": "404 Page Not Found",
		}).Info("request start")
		w.WriteHeader(404)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf(
			`{"message":"Page Not Found %s %s"}`, r.Method, r.URL.Path)))
	}
}