package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
)

// Handler is the graylog REST API's handler.
// the argument `user` is the authenticated user and are mainly used for the authorization.
type Handler func(user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params) (interface{}, int, error)

func wrapEchoHandle(lgc *logic.Logic, handler Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		// logging
		lgc.Logger().WithFields(log.Fields{
			"path": req.URL.Path, "method": req.Method,
		}).Info("request start")
		// authentication
		var user *graylog.User
		if lgc.Auth() {
			authName, authPass, ok := req.BasicAuth()
			if !ok {
				lgc.Logger().WithFields(log.Fields{
					"path": req.URL.Path, "method": req.Method,
				}).Warn("request basic authentication header is not set")
				return c.JSON(401, map[string]string{})
			}
			var (
				sc  int
				err error
			)
			user, sc, err = lgc.Authenticate(authName, authPass)
			if err != nil {
				if sc == 401 {
					return c.JSON(401, map[string]string{})
				}
				return c.JSON(sc, NewAPIError(err.Error()))
			}
			lgc.Logger().WithFields(log.Fields{
				"path": req.URL.Path, "method": req.Method,
				"user_name": user.Username,
			}).Info("request user name")
		}

		// call handler
		body, sc, err := handler(user, lgc, req, NewParams(c))
		// set status code
		// write response body
		if err != nil {
			return c.JSON(sc, NewAPIError(err.Error()))
		}
		return c.JSON(sc, body)
	}
}
