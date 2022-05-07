package routes

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/msmaiaa/eldenring-checklist/pkg/db"
	"github.com/msmaiaa/eldenring-checklist/pkg/models"
	"gorm.io/gorm"
)


func GetRegions(c echo.Context) error {
	regions := []models.Region{}
	db.GetDB().Find(&regions)
	return c.JSON(http.StatusOK, regions)
}

func AddRegion(c echo.Context) error {
	type AddRegionDTO struct {
		Name string `json:"name" validate:"required"`
	}
	var body AddRegionDTO
	if err := c.Bind(&body); err != nil {			
		return c.NoContent(http.StatusNotFound)
	}
	if err:= c.Validate(&body); err != nil {
		return err
	}
	region := models.Region {
		Name: body.Name,
	}
	db.GetDB().Create(&region)
	return c.JSON(http.StatusOK, region)
}

func UpdateRegion(c echo.Context) error {
	id := c.Param("id")

	result := db.GetDB().First(&models.Region{}, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "Couldn't find a region with the provided id")
	}
	var body models.Region
	if err := c.Bind(&body); err != nil {			
		return c.NoContent(http.StatusNotFound)
	}
	if err:= c.Validate(&body); err != nil {
		return err
	}

	db.GetDB().Model(&body).Updates(body)
	return c.JSON(http.StatusOK, body)
}

func DeleteRegion(c echo.Context) error {
	id := c.Param("id")

	result := db.GetDB().First(&models.Region{}, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "Couldn't find a region with the provided id")
	}
	db.GetDB().Where("id = ?", id).Delete(&models.Region{})
	return c.NoContent(http.StatusOK)
}
////////////////////////

func GetCategories(c echo.Context) error {
	categories := []models.Category{}
	db.GetDB().Find(&categories)
	return c.JSON(http.StatusOK, categories)
}

func AddCategory(c echo.Context) error {
	type AddCategoryDTO struct {
		Name string `json:"name" validate:"required"`
	}
	var body AddCategoryDTO
	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	if err:= c.Validate(&body); err != nil {
		return err
	}
	category := models.Category {
		Name: body.Name,
	}
	db.GetDB().Create(&category)
	return c.JSON(http.StatusOK, category)
}

func UpdateCategory(c echo.Context) error {
	id := c.Param("id")

	result := db.GetDB().First(&models.Category{}, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "Couldn't find a category with the provided id")
	}
	var body models.Category
	if err := c.Bind(&body); err != nil {			
		return c.NoContent(http.StatusNotFound)
	}
	if err:= c.Validate(&body); err != nil {
		return err
	}

	db.GetDB().Model(&body).Updates(body)
	return c.JSON(http.StatusOK, body)
}

func DeleteCategory(c echo.Context) error {
	id := c.Param("id")

	result := db.GetDB().First(&models.Category{}, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "Couldn't find a category with the provided id")
	}
	db.GetDB().Where("id = ?", id).Delete(&models.Category{})
	return c.NoContent(http.StatusOK)
}
////////////////////////

func GetEntity(c echo.Context) error {
	entities := []models.Entity{}
	db.GetDB().Find(&entities)
	return c.JSON(http.StatusOK, entities)
}

func AddEntity(c echo.Context) error {
	type AddEntityDTO struct {
		Name string `json:"name" validate:"required"`
		CategoryID uint `json:"categoryId" validate:"required"`
		RegionID uint `json:"regionId" validate:"required"`
		X int16 `json:"x" validate:"required"`
		Y int16 `json:"y" validate:"required"`
	}
	var body AddEntityDTO
	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusNoContent)
	}
	if err := c.Validate(&body); err != nil {
		return err
	}
	entity := models.Entity {
		Name: body.Name,
		X: body.X,
		Y: body.X,
		CategoryID: body.CategoryID,
		RegionID: body.RegionID,
	}
	db.GetDB().Create(&entity)
	return c.JSON(http.StatusOK, entity)
}

func UpdateEntity(c echo.Context) error {
	id := c.Param("id")

	result := db.GetDB().First(&models.Entity{}, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "Couldn't find an entity with the provided id")
	}
	var body models.Entity
	if err := c.Bind(&body); err != nil {			
		return c.NoContent(http.StatusNotFound)
	}
	if err:= c.Validate(&body); err != nil {
		return err
	}

	db.GetDB().Model(&body).Updates(body)
	return c.JSON(http.StatusOK, body)
}

func DeleteEntity(c echo.Context) error {
	id := c.Param("id")

	result := db.GetDB().First(&models.Entity{}, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "Couldn't find an entity with the provided id")
	}
	db.GetDB().Where("id = ?", id).Delete(&models.Entity{})
	return c.NoContent(http.StatusOK)
}