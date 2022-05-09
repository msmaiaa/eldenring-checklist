package entity

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/msmaiaa/eldenring-checklist/db"
	"github.com/msmaiaa/eldenring-checklist/db/models"
	"gorm.io/gorm"
)

func (EntityRouter) GetEntity(c echo.Context) error {
	entities := []models.Entity{}
	db.GetDB().Find(&entities)
	return c.JSON(http.StatusOK, entities)
}

func (EntityRouter) AddEntity(c echo.Context) error {
	type AddEntityDTO struct {
		Name       string `json:"name" validate:"required"`
		CategoryID uint   `json:"categoryId" validate:"required"`
		RegionID   uint   `json:"regionId" validate:"required"`
		X          int16  `json:"x" validate:"required"`
		Y          int16  `json:"y" validate:"required"`
	}
	var body AddEntityDTO
	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusNoContent)
	}
	if err := c.Validate(&body); err != nil {
		return err
	}
	entity := models.Entity{
		Name:       body.Name,
		X:          body.X,
		Y:          body.X,
		CategoryID: body.CategoryID,
		RegionID:   body.RegionID,
	}
	db.GetDB().Create(&entity)
	return c.JSON(http.StatusOK, entity)
}

func (EntityRouter) UpdateEntity(c echo.Context) error {
	id := c.Param("id")

	result := db.GetDB().First(&models.Entity{}, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "Couldn't find an entity with the provided id")
	}
	var body models.Entity
	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	db.GetDB().Model(&body).Updates(body)
	return c.JSON(http.StatusOK, body)
}

func (EntityRouter) DeleteEntity(c echo.Context) error {
	id := c.Param("id")

	result := db.GetDB().First(&models.Entity{}, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "Couldn't find an entity with the provided id")
	}
	db.GetDB().Where("id = ?", id).Delete(&models.Entity{})
	return c.NoContent(http.StatusOK)
}