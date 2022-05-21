package auth

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
	"github.com/msmaiaa/eldenring-checklist/db"
	"github.com/msmaiaa/eldenring-checklist/db/models"
	"github.com/msmaiaa/eldenring-checklist/lib/steam"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func (AuthRouter) Login(c echo.Context) error {
	url, err := steam.GenerateUrl()
	if err != nil {
		c.Logger().Fatal("Error creating redirect URL: %q\n", err)
		return c.String(http.StatusInternalServerError, "Error creating redirect URL")
	}
	return c.Redirect(303, url)
}

type JWTPayload struct {
	Id   string `json:"id"`
	Role string `json:"role"`
}

type JWTClaim struct {
	JWTPayload
	jwt.StandardClaims
}

type LoginResponse struct {
	Token string `json:"token"`
}

//TODO: move these functions to a repository?
func GetUser(id string) (models.User, error) {
	user := models.User{}
	result := db.GetDB().First(&user, id)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func CreateUser(id string, role string) (models.User, error) {
	user := models.User{
		Steamid64: id,
		Role:      role,
	}
	if err := db.GetDB().Create(&user).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return user, echo.NewHTTPError(http.StatusConflict, "User already exists")
			}
		}
		return user, echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return user, nil
}

func signJwt(payload JWTPayload) (string, error) {
	claims := JWTClaim{
		JWTPayload: payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Unix() + 864000,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (AuthRouter) SteamReturn(c echo.Context) error {
	id, err := steam.VerifyURL(c)
	if err != nil {
		c.Logger().Error("Error verifying: %q\n", err)
		return c.String(http.StatusInternalServerError, "Error while trying to authenticate")
	}
	steamUser := steam.QuerySteamUser(&id)
	var role string
	steamid := steamUser.Steamid
	foundUser, foundUserErr := GetUser(steamid)
	if foundUserErr != nil {
		created, err := CreateUser(steamid, "user")
		if err != nil {
			return err
		}
		role = created.Role
	} else {
		role = foundUser.Role
	}

	tokenString, tokenErr := signJwt(JWTPayload{Id: steamid, Role: role})
	if tokenErr != nil {
		return c.JSON(http.StatusInternalServerError, tokenErr)
	}
	return c.JSON(http.StatusOK, LoginResponse{
		Token: tokenString,
	})
}
