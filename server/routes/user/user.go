package user

import (
	"github.com/labstack/echo/v4"
	"github.com/msmaiaa/eldenring-checklist/middleware"
)

type UserRouter struct{}

func (ctrl UserRouter) Init(g *echo.Group) {
	g.POST("/", ctrl.AddUser, middleware.AdminMiddleware)
	g.GET("/:id/checks", ctrl.GetUserChecks)
}
