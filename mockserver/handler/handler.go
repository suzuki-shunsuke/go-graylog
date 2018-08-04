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
type Handler func(user *graylog.User, lgc *logic.Logic, r *http.Request, ps httprouter.Params) (interface{}, int, error)

func wrapHandle(lgc *logic.Logic, handler Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// logging
		lgc.Logger().WithFields(log.Fields{
			"path": r.URL.Path, "method": r.Method,
		}).Info("request start")
		// set header
		w.Header().Set("Content-Type", "application/json")
		// authentication
		var user *graylog.User
		if lgc.Auth() {
			authName, authPass, ok := r.BasicAuth()
			if !ok {
				lgc.Logger().WithFields(log.Fields{
					"path": r.URL.Path, "method": r.Method,
				}).Warn("request basic authentication header is not set")
				w.WriteHeader(401)
				return
			}
			var (
				sc  int
				err error
			)
			user, sc, err = lgc.Authenticate(authName, authPass)
			if err != nil {
				w.WriteHeader(sc)
				if sc == 401 {
					return
				}
				ae := NewAPIError(err.Error())
				b, err := json.Marshal(ae)
				if err != nil {
					w.Write([]byte(`{"message":"failed to authenticate"}`))
					return
				}
				w.Write(b)
				return
			}
			lgc.Logger().WithFields(log.Fields{
				"path": r.URL.Path, "method": r.Method,
				"user_name": user.Username,
			}).Info("request user name")
		}

		// call handler
		body, sc, err := handler(user, lgc, r, ps)
		// set status code
		// write response body
		if err != nil {
			w.WriteHeader(sc)

			ae := NewAPIError(err.Error())
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
