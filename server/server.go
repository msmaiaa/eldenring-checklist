package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/msmaiaa/eldenring-checklist/pkg/routes"
)

func main() {
	dotenv_err := godotenv.Load()
	if dotenv_err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	e.GET("/api/v1/regions", routes.GetRegions)
	e.POST("/api/v1/regions", routes.AddRegion)
	e.PUT("/api/v1/regions/:id", routes.UpdateRegion)
	e.DELETE("/api/v1/regions/:id", routes.DeleteRegion)

	e.GET("/api/v1/categories", routes.GetCategories)
	e.POST("/api/v1/categories", routes.AddCategory)
	e.PUT("/api/v1/categories/:id", routes.UpdateCategory)
	e.DELETE("/api/v1/categories/:id", routes.DeleteCategory)

	e.GET("/api/v1/enemies", routes.GetEnemies)
	e.POST("/api/v1/enemies", routes.AddEnemy)
	e.PUT("/api/v1/enemies/:id", routes.UpdateEnemy)
	e.DELETE("/api/v1/enemies/:id", routes.DeleteEnemy)

	e.Logger.Fatal(e.Start(":1337"))
}