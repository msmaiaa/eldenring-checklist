package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/msmaiaa/eldenring-checklist/db"
	"github.com/msmaiaa/eldenring-checklist/db/models"
)

func (UserRouter) AddUser(c echo.Context) error {
	type AddUserDTO struct {
		Steamid64 string `json:"steamid64" validate:"required"`
		Role      string `json:"role" validate:"required"`
	}
	var body AddUserDTO
	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	if err := c.Validate(&body); err != nil {
		return err
	}
	if body.Role != "admin" && body.Role != "user" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid role")
	}
	user := models.User{
		Steamid64: body.Steamid64,
		Role:      body.Role,
	}
	if err := db.GetDB().Create(&user).Error; err != nil {
		if pgErr, isPgError := db.GetPostgresError(&err); isPgError {
			if pgErr.Code == "23505" {
				return c.JSON(http.StatusConflict, echo.Map{
					"error": "User already exists",
				})
			}
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, user)
}

func (UserRouter) GetUserChecks(c echo.Context) error {
	id := c.Param("id")
	entities := []string{}
	_db := db.GetDB()
	_db.Table("checks").Where("user_id = ?", id).Pluck("entity_id", &entities)
	return c.JSON(http.StatusOK, entities)
}
