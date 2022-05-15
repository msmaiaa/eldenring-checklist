package check

import (
	"github.com/labstack/echo/v4"
)

func (CheckRouter) AddCheck(c echo.Context) error {
	type AddCheckDTO struct {}
	return nil
}


func (CheckRouter) DeleteCheck(c echo.Context) error {
	return nil
}
