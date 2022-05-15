package user

import "github.com/labstack/echo/v4"

type UserRouter struct {}
func (ctrl UserRouter) Init(g *echo.Group) {	
	g.POST("/", ctrl.AddUser)
	g.GET("/:id/checks", ctrl.GetUserChecks)
}