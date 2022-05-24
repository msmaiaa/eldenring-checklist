package category

import (
	"github.com/labstack/echo/v4"
	"github.com/msmaiaa/eldenring-checklist/middleware"
)

type CategoryRouter struct{}

func (ctrl CategoryRouter) Init(g *echo.Group) {
	g.GET("/", ctrl.GetCategories)
	g.POST("/", ctrl.AddCategory, middleware.AdminMiddleware)
	g.PUT("/:id", ctrl.UpdateCategory, middleware.AdminMiddleware)
	g.DELETE("/:id", ctrl.DeleteCategory, middleware.AdminMiddleware)
}
