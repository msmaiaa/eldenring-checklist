package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/msmaiaa/eldenring-checklist/routes/category"
	"github.com/msmaiaa/eldenring-checklist/routes/entity"
	"github.com/msmaiaa/eldenring-checklist/routes/region"
)

func Routes(g *echo.Group) {
	region.RegionRouter{}.Init(g.Group("/api/v1/region"))
	entity.EntityRouter{}.Init(g.Group("/api/v1/entity"))
	category.CategoryRouter{}.Init(g.Group("/api/v1/category"))
}

////////////////////////