package check

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/msmaiaa/eldenring-checklist/db"
	"github.com/msmaiaa/eldenring-checklist/db/models"
	"github.com/msmaiaa/eldenring-checklist/routes/auth"
)

func getUserFromContext(c echo.Context) auth.JWTPayload {
	return c.Get("user").(*jwt.Token).Claims.(*auth.JWTClaim).JWTPayload
}

func (CheckRouter) AddCheck(c echo.Context) error {
	userId := getUserFromContext(c).Id
	type AddCheckDTO struct {
		EntityId uint `json:"entityId" validate:"required"`
	}
	var body AddCheckDTO
	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Validate(&body); err != nil {
		return err
	}
	if err := db.GetDB().Where("entity_id = ? AND user_id = ?", body.EntityId, userId).First(&models.Check{}).Error; err == nil {
		return c.NoContent(http.StatusConflict)
	}
	check := models.Check{
		UserId:   userId,
		EntityId: body.EntityId,
	}
	db.GetDB().Create(&check)
	return c.JSON(http.StatusCreated, check)
}

func (CheckRouter) DeleteCheck(c echo.Context) error {
	userId := getUserFromContext(c).Id
	type DeleteCheckDTO struct {
		EntityId uint `json:"entityId" validate:"required"`
	}
	var body DeleteCheckDTO
	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	if err := c.Validate(&body); err != nil {
		return err
	}
	if err := db.GetDB().Where("entity_id = ? AND user_id = ?", body.EntityId, userId).First(&models.Check{}).Error; err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	db.GetDB().Where("entity_id = ? AND user_id = ?", body.EntityId, userId).Delete(&models.Check{})
	return c.NoContent(http.StatusOK)
}
