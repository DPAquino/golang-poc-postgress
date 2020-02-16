package router

import (
	"golang-poc-postgress/registry"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, r registry.Registry) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(context echo.Context) error { return r.UserController().GetUsers(context) })
	e.GET("/roles", func(context echo.Context) error { return r.RoleController().GetRoles(context) })

	return e
}
