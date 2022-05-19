package config

import (
	"os"

	"github.com/labstack/echo/v4/middleware"
	"github.com/msmaiaa/eldenring-checklist/routes/auth"
)

var JWT = middleware.JWTConfig{
	SigningKey: []byte(os.Getenv("JWT_SECRET")),
	Claims:     &auth.JWTClaim{},
}
