package region

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/msmaiaa/eldenring-checklist/db"
	"github.com/msmaiaa/eldenring-checklist/db/models"
	"gorm.io/gorm"
)

func (RegionRouter) GetRegions(c echo.Context) error {
	regions := []models.Region{}
	db.GetDB().Find(&regions)
	return c.JSON(http.StatusOK, regions)
}

func (RegionRouter) AddRegion(c echo.Context) error {
	type AddRegionDTO struct {
		Name string `json:"name" validate:"required"`
	}
	var body AddRegionDTO
	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	if err := c.Validate(&body); err != nil {
		return err
	}
	region := models.Region{
		Name: body.Name,
	}
	if err := db.GetDB().Create(&region).Error; err != nil {
		if pgErr, isPgError := db.GetPostgresError(&err); isPgError {
			if pgErr.Code == "23505" {
				return c.JSON(http.StatusConflict, echo.Map{
					"error": "A region with the provided name already exists",
				})
			}
		}
	}
	return c.JSON(http.StatusOK, region)
}

func (RegionRouter) UpdateRegion(c echo.Context) error {
	id := c.Param("id")

	result := db.GetDB().First(&models.Region{}, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "Couldn't find a region with the provided id")
	}
	var body models.Region
	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	db.GetDB().Model(&body).Updates(body)
	return c.JSON(http.StatusOK, body)
}

func (RegionRouter) DeleteRegion(c echo.Context) error {
	id := c.Param("id")

	result := db.GetDB().First(&models.Region{}, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "Couldn't find a region with the provided id")
	}
	db.GetDB().Where("id = ?", id).Delete(&models.Region{})
	return c.NoContent(http.StatusOK)
}
