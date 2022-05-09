package entity

import "github.com/labstack/echo/v4"

type EntityRouter struct{}

func (ctrl EntityRouter) Init(g *echo.Group) {
	g.GET("/", ctrl.GetEntity)
	g.POST("/", ctrl.AddEntity)
	g.PUT("/:id", ctrl.UpdateEntity)
	g.DELETE("/:id", ctrl.DeleteEntity)
}

