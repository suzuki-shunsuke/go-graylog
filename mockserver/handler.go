package mockserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

type Handler func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (int, interface{}, error)

func (ms *MockServer) handleNotFound(w http.ResponseWriter, r *http.Request) {
	ms.Logger().WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
		"message": "404 Page Not Found",
	}).Info("request start")
	w.WriteHeader(404)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(
		`{"message":"Page Not Found %s %s"}`, r.Method, r.URL.Path)))
}

func wrapHandle(ms *MockServer, handler Handler) httprouter.Handle {
	// ms.Logger().WithFields(log.Fields{
	// 	"path": r.URL.Path, "method": r.Method,
	// }).Info("request start")
	// w.Header().Set("Content-Type", "application/json")
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		ms.Logger().WithFields(log.Fields{
			"path": r.URL.Path, "method": r.Method,
		}).Info("request start")
		sc, body, err := handler(w, r, ps)
		if err != nil {
			w.WriteHeader(sc)
			w.Write([]byte(fmt.Sprintf(
				`{"type": "ApiError", "message": "%s"}`, err.Error())))
			return
		}
		if body == nil {
			return
		}
		b, err := json.Marshal(body)
		if err == nil {
			w.Write(b)
			return
		}
		w.WriteHeader(500)
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
	}
}
