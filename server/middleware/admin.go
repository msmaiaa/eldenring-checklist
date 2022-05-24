package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/msmaiaa/eldenring-checklist/routes/auth"
)

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userToken := c.Get("user").(*jwt.Token)
		claims := userToken.Claims.(*auth.JWTClaim)
		payload := claims.JWTPayload
		if payload.Role != "admin" {
			return echo.ErrUnauthorized
		}
		next(c)
		return nil
	}
}
