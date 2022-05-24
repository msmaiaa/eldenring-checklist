package region

import (
	"github.com/labstack/echo/v4"
	"github.com/msmaiaa/eldenring-checklist/middleware"
)

type RegionRouter struct{}

func (ctrl RegionRouter) Init(g *echo.Group) {
	g.GET("/", ctrl.GetRegions)
	g.POST("/", ctrl.AddRegion, middleware.AdminMiddleware)
	g.PUT("/:id", ctrl.UpdateRegion, middleware.AdminMiddleware)
	g.DELETE("/:id", ctrl.DeleteRegion, middleware.AdminMiddleware)
}
