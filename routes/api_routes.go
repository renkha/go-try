package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/renkha/go-try/app/user"
	"github.com/renkha/go-try/helper"
)

func DefineApiRoutes(e *echo.Echo) {
	handlers := []helper.Handler{
		user.UserRoutes{},
	}
	var routes []helper.Route

	for _, handler := range handlers {
		routes = append(routes, handler.Route()...)
	}

	api := e.Group("/api/v1")

	for _, route := range routes {
		switch route.Method {
		case echo.POST:
			api.POST(route.Path, route.Handler, route.Middlerware...)
		case echo.GET:
			api.GET(route.Path, route.Handler, route.Middlerware...)
		case echo.PUT:
			api.PUT(route.Path, route.Handler, route.Middlerware...)
		case echo.DELETE:
			api.DELETE(route.Path, route.Handler, route.Middlerware...)
		}
	}
}
