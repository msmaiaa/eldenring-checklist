package auth

import "github.com/labstack/echo/v4"

type AuthRouter struct {}

func (ctrl AuthRouter) Init(g *echo.Group) {
	g.GET("/login", ctrl.Login)
	g.GET("/steam/return", ctrl.SteamReturn)
}