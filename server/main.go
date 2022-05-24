package main

import (
	"fmt"
	"os"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/msmaiaa/eldenring-checklist/db"
	"github.com/msmaiaa/eldenring-checklist/lib"
	"github.com/msmaiaa/eldenring-checklist/routes"
	"github.com/rs/zerolog"
	"github.com/ziflex/lecho/v3"
)

func setupLogger(e *echo.Echo) {
	output := zerolog.ConsoleWriter{Out: os.Stdout}
	logger := lecho.New(
		output,
		lecho.WithLevel(log.DEBUG),
		lecho.WithTimestamp(),
	)
	e.Logger = logger
	e.Use(middleware.RequestID())
	e.Use(lecho.Middleware(lecho.Config{
		Logger: logger,
	}))
}

func main() {
	dotenv_err := godotenv.Load()
	if dotenv_err != nil {
		panic(dotenv_err)
	}

	db.Connect()

	e := echo.New()
	setupLogger(e)

	e.Validator = &lib.CustomValidator{Validator: validator.New()}

	routes.Routes(e.Group("/api/v1"))

	e.Logger.Fatal(e.Start(fmt.Sprintf("127.0.0.1:%s", os.Getenv("PORT"))))
}
