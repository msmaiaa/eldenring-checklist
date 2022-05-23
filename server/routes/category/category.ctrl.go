package category

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/msmaiaa/eldenring-checklist/db"
	"github.com/msmaiaa/eldenring-checklist/db/models"
	"gorm.io/gorm"
)

func (CategoryRouter) GetCategories(c echo.Context) error {
	categories := []models.Category{}
	db.GetDB().Find(&categories)
	return c.JSON(http.StatusOK, categories)
}

//TODO: return an error if a category with the same name already exists
func (CategoryRouter) AddCategory(c echo.Context) error {
	type AddCategoryDTO struct {
		Name string `json:"name" validate:"required"`
	}
	var body AddCategoryDTO
	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	if err := c.Validate(&body); err != nil {
		return err
	}
	category := models.Category{
		Name: body.Name,
	}
	db.GetDB().Create(&category)
	return c.JSON(http.StatusOK, category)
}

func (CategoryRouter) UpdateCategory(c echo.Context) error {
	id := c.Param("id")

	result := db.GetDB().First(&models.Category{}, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "Couldn't find a category with the provided id")
	}
	var body models.Category
	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	db.GetDB().Model(&body).Updates(body)
	return c.JSON(http.StatusOK, body)
}

func (CategoryRouter) DeleteCategory(c echo.Context) error {
	id := c.Param("id")

	result := db.GetDB().First(&models.Category{}, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "Couldn't find a category with the provided id")
	}
	db.GetDB().Where("id = ?", id).Delete(&models.Category{})
	return c.NoContent(http.StatusOK)
}
