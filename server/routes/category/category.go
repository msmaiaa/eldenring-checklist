package category

import "github.com/labstack/echo/v4"

type CategoryRouter struct{}

func (ctrl CategoryRouter) Init(g *echo.Group) {
	g.GET("/", ctrl.GetCategories)
	g.POST("/", ctrl.AddCategory)
	g.PUT("/:id", ctrl.UpdateCategory)
	g.DELETE("/:id", ctrl.DeleteCategory)
}

