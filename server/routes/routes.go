package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/msmaiaa/eldenring-checklist/config"
	"github.com/msmaiaa/eldenring-checklist/routes/auth"
	"github.com/msmaiaa/eldenring-checklist/routes/category"
	"github.com/msmaiaa/eldenring-checklist/routes/check"
	"github.com/msmaiaa/eldenring-checklist/routes/entity"
	"github.com/msmaiaa/eldenring-checklist/routes/region"
	"github.com/msmaiaa/eldenring-checklist/routes/user"
)

func Routes(g *echo.Group) {
	auth.AuthRouter{}.Init(g.Group(("/auth")))
	g.Use(middleware.JWTWithConfig(config.JWT))
	region.RegionRouter{}.Init(g.Group("/region"))
	entity.EntityRouter{}.Init(g.Group("/entity"))
	category.CategoryRouter{}.Init(g.Group("/category"))
	check.CheckRouter{}.Init(g.Group(("/check")))
	user.UserRouter{}.Init(g.Group(("/user")))
}

////////////////////////
