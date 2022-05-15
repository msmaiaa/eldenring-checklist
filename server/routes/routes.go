package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/msmaiaa/eldenring-checklist/routes/category"
	"github.com/msmaiaa/eldenring-checklist/routes/check"
	"github.com/msmaiaa/eldenring-checklist/routes/entity"
	"github.com/msmaiaa/eldenring-checklist/routes/region"
	"github.com/msmaiaa/eldenring-checklist/routes/user"
)

func Routes(g *echo.Group) {
	region.RegionRouter{}.Init(g.Group("/api/v1/region"))
	entity.EntityRouter{}.Init(g.Group("/api/v1/entity"))
	category.CategoryRouter{}.Init(g.Group("/api/v1/category"))
	check.CheckRouter{}.Init(g.Group(("/api/v1/check")))
	user.UserRouter{}.Init(g.Group(("/api/v1/user")))
}

////////////////////////