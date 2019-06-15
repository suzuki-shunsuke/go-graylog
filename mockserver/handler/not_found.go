package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
)

// HandleNotFound is the generator of the NotFound handler.
func HandleNotFound(lgc *logic.Logic) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		lgc.Logger().WithFields(log.Fields{
			"path": req.URL.Path, "method": req.Method,
			"message": "404 Page Not Found",
		}).Info("request start")
		return c.JSON(404, map[string]string{
			"message": fmt.Sprintf("Page Not Found %s %s", req.Method, req.URL.Path),
		})
	}
}
