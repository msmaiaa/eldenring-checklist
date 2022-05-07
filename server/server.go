package main

import (
	"fmt"
	"log"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/msmaiaa/eldenring-checklist/pkg/db"
	"github.com/msmaiaa/eldenring-checklist/pkg/models"
	"github.com/msmaiaa/eldenring-checklist/pkg/routes"
	"github.com/msmaiaa/eldenring-checklist/pkg/util"
)

func main() {
	dotenv_err := godotenv.Load()
	if dotenv_err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connect()
	fmt.Println("\033[32m Connected to the database")

	db.GetDB().AutoMigrate(
		&models.Category{}, 
		&models.Entity{}, 
		&models.Region{},
		&models.User{},
	)

	fmt.Println("\033[32m Schemas migrated successfully")

	e := echo.New()
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	e.GET("/api/v1/region", routes.GetRegions)
	e.POST("/api/v1/region", routes.AddRegion)
	e.PUT("/api/v1/region/:id", routes.UpdateRegion)
	e.DELETE("/api/v1/region/:id", routes.DeleteRegion)

	e.GET("/api/v1/category", routes.GetCategories)
	e.POST("/api/v1/category", routes.AddCategory)
	e.PUT("/api/v1/category/:id", routes.UpdateCategory)
	e.DELETE("/api/v1/category/:id", routes.DeleteCategory)

	e.GET("/api/v1/entity", routes.GetEntity)
	e.POST("/api/v1/entity", routes.AddEntity)
	e.PUT("/api/v1/entity/:id", routes.UpdateEntity)
	e.DELETE("/api/v1/entity/:id", routes.DeleteEntity)

	fmt.Println("\033[32m The routes have been set up")

	e.Logger.Fatal(e.Start("127.0.0.1:1337"))
}