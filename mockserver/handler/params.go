package handler

import (
	"github.com/labstack/echo/v4"
)

type (
	// Params is an interface of request parameters.
	Params interface {
		PathParam(name string) string
		QueryParam(name string) string
	}

	// EchoParams is an implementation of Params for echo.
	echoParams struct {
		c echo.Context
	}
)

// NewParams returns a Params.
func NewParams(c echo.Context) Params {
	return NewEchoParams(c)
}

// NewEchoParams returns a Params.
func NewEchoParams(c echo.Context) Params {
	return &echoParams{c: c}
}

// PathParam returns a path parameter.
func (params *echoParams) PathParam(name string) string {
	return params.c.Param(name)
}

// QueryParam returns a query parameter.
func (params *echoParams) QueryParam(name string) string {
	return params.c.QueryParam(name)
}
