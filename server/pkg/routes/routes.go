package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


func GetRegions(c echo.Context) error {
	return c.String(http.StatusOK, "get regions")
}

func AddRegion(c echo.Context) error {
	return c.String(http.StatusOK, "add region")
}

func UpdateRegion(c echo.Context) error {
	return c.String(http.StatusOK, "update region")
}

func DeleteRegion(c echo.Context) error {
	return c.String(http.StatusOK, "delete region")
}
////////////////////////

func GetCategories(c echo.Context) error {
	return c.String(http.StatusOK, "get categories")
}

func AddCategory(c echo.Context) error {
	return c.String(http.StatusOK, "add Category")
}

func UpdateCategory(c echo.Context) error {
	return c.String(http.StatusOK, "update Category")
}

func DeleteCategory(c echo.Context) error {
	return c.String(http.StatusOK, "delete Category")
}
////////////////////////

func GetEnemies(c echo.Context) error {
	return c.String(http.StatusOK, "get enemies")
}

func AddEnemy(c echo.Context) error {
	return c.String(http.StatusOK, "add Enemy")
}

func UpdateEnemy(c echo.Context) error {
	return c.String(http.StatusOK, "update Enemy")
}

func DeleteEnemy(c echo.Context) error {
	return c.String(http.StatusOK, "delete Enemy")
}