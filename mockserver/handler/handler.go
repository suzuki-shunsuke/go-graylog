package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
)

// Handler is the graylog REST API's handler.
// the argument `user` is the authenticated user and are mainly used for the authorization.
type Handler func(user *graylog.User, ms *logic.Logic, w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, int, error)

func wrapHandle(ms *logic.Logic, handler Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ms.Logger().WithFields(log.Fields{
			"path": r.URL.Path, "method": r.Method,
		}).Info("request start")
		w.Header().Set("Content-Type", "application/json")
		// authentication
		var user *graylog.User
		if ms.Auth() {
			authName, authPass, ok := r.BasicAuth()
			if !ok {
				ms.Logger().WithFields(log.Fields{
					"path": r.URL.Path, "method": r.Method,
				}).Warn("request basic authentication header is not set")
				w.WriteHeader(401)
				return
			}
			var (
				sc  int
				err error
			)
			user, sc, err = ms.Authenticate(authName, authPass)
			if err != nil {
				w.WriteHeader(sc)
				if sc == 401 {
					return
				}
				ae := graylog.NewAPIError(err.Error())
				b, err := json.Marshal(ae)
				if err != nil {
					w.Write([]byte(`{"message":"failed to authenticate"}`))
					return
				}
				w.Write(b)
				return
			}
			ms.Logger().WithFields(log.Fields{
				"path": r.URL.Path, "method": r.Method,
				"user_name": user.Username,
			}).Info("request user name")
		}

		body, sc, err := handler(user, ms, w, r, ps)
		if err != nil {
			w.WriteHeader(sc)

			ae := graylog.NewAPIError(err.Error())
			b, err := json.Marshal(ae)
			if err != nil {
				w.Write([]byte(`{"message":"failed to marshal an APIError"}`))
				return
			}
			w.Write(b)
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
