package check

import "github.com/labstack/echo/v4"

type CheckRouter struct {}
func (ctrl CheckRouter) Init(g *echo.Group) {
	g.POST("/", ctrl.AddCheck)
	g.DELETE("/:id", ctrl.DeleteCheck)
}