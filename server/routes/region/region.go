package region

import "github.com/labstack/echo/v4"

type RegionRouter struct{}

func (ctrl RegionRouter) Init(g *echo.Group) {
	g.GET("/", ctrl.GetRegions)
	g.POST("/", ctrl.AddRegion)
	g.PUT("/:id", ctrl.UpdateRegion)
	g.DELETE("/:id", ctrl.DeleteRegion)
}

