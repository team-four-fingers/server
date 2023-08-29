package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

type HTTPServer struct {
	*echo.Echo
}

type HTTPHandler interface {
	Method() string
	Path() string
	HandleFunc() func(c echo.Context) error
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		echo.New(),
	}
}

func (s *HTTPServer) Start(port string) error {
	logrus.Info("starting server on port: ", port)

	return s.Echo.Start(fmt.Sprintf(":%s", port))
}

func (s *HTTPServer) RegisterHandler(httpHandler HTTPHandler, middlewareFunc ...echo.MiddlewareFunc) (
	*echo.Route,
	error,
) {
	var (
		method      = httpHandler.Method()
		path        = httpHandler.Path()
		handlerFunc = httpHandler.HandleFunc()
	)

	if method == http.MethodGet {
		return s.Echo.GET(path, handlerFunc, middlewareFunc...), nil
	}

	if method == http.MethodPost {
		return s.Echo.POST(path, handlerFunc, middlewareFunc...), nil
	}

	if method == http.MethodPut {
		return s.Echo.PUT(path, handlerFunc, middlewareFunc...), nil
	}

	if method == http.MethodDelete {
		return s.Echo.DELETE(path, handlerFunc, middlewareFunc...), nil
	}

	if method == http.MethodPatch {
		return s.Echo.PATCH(path, handlerFunc, middlewareFunc...), nil
	}

	return nil, errors.New("not supported method")
}

func (s *HTTPServer) RegisterMiddleware(middlewareFunc ...echo.MiddlewareFunc) {
	s.Echo.Use(middlewareFunc...)
}

func (s *HTTPServer) RegisterHTTPHandlers(httpHandlers ...HTTPHandler) error {
	for _, handler := range httpHandlers {
		route, err := s.RegisterHandler(handler)
		if err != nil {
			return err
		}
		logrus.Info("registered handler: ", route.Path)
	}

	return nil
}
