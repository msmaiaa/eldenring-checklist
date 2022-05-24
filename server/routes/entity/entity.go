package entity

import (
	"github.com/labstack/echo/v4"
	"github.com/msmaiaa/eldenring-checklist/middleware"
)

type EntityRouter struct{}

func (ctrl EntityRouter) Init(g *echo.Group) {
	g.GET("/", ctrl.GetEntity)
	g.POST("/", ctrl.AddEntity, middleware.AdminMiddleware)
	g.PUT("/:id", ctrl.UpdateEntity, middleware.AdminMiddleware)
	g.DELETE("/:id", ctrl.DeleteEntity, middleware.AdminMiddleware)
}
