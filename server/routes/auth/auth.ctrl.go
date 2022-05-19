package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/imroc/req/v3"
	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
	"github.com/msmaiaa/eldenring-checklist/db"
	"github.com/msmaiaa/eldenring-checklist/db/models"
	"github.com/yohcop/openid-go"
)

type NoOpDiscoveryCache struct{}

func (n *NoOpDiscoveryCache) Put(id string, info openid.DiscoveredInfo) {}
func (n *NoOpDiscoveryCache) Get(id string) openid.DiscoveredInfo {
	return nil
}

var nonceStore = openid.NewSimpleNonceStore()
var discoveryCache = &NoOpDiscoveryCache{}
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

//TODO: move steam stuff to its own package
type SteamUser struct {
	Steamid                  string `json:"steamid"`
	Communityvisibilitystate int    `json:"communityvisibilitystate"`
	Profilestate             int    `json:"profilestate"`
	Personaname              string `json:"personaname"`
	Commentpermission        int    `json:"commentpermission"`
	Profileurl               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	Avatarmedium             string `json:"avatarmedium"`
	Avatarfull               string `json:"avatarfull"`
	Avatarhash               string `json:"avatarhash"`
	Lastlogoff               int    `json:"lastlogoff"`
	Personastate             int    `json:"personastate"`
	Realname                 string `json:"realname"`
	Primaryclanid            string `json:"primaryclanid"`
	Timecreated              int    `json:"timecreated"`
	Personastateflags        int    `json:"personastateflags"`
	Loccountrycode           string `json:"loccountrycode"`
}
type SteamApiResponse struct {
	Response struct {
		Players []SteamUser `json:"players"`
	} `json:"response"`
}

func (AuthRouter) Login(c echo.Context) error {
	url, err := openid.RedirectURL(
		"http://steamcommunity.com/openid",
		os.Getenv("STEAM_RETURN_URL"),
		os.Getenv("BASE_URL"))

	if err != nil {
		c.Logger().Fatal("Error creating redirect URL: %q\n", err)
		return c.String(http.StatusInternalServerError, "Error creating redirect URL")
	}
	return c.Redirect(303, url)
}

func GetSteamUser(url *string) SteamUser {
	response := SteamApiResponse{}
	steamId := strings.Replace(*url, "http://steamcommunity.com/openid/id/", "", 1)

	//TODO: start client elsewhere
	client := req.C().
		DevMode()

	//TODO: handle possible request errors
	client.R().
		SetResult(&response).
		Get(fmt.Sprintf("https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", os.Getenv("STEAM_KEY"), steamId))
	return response.Response.Players[0]
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

//TODO: Refactor this entire file
func GetUser(id string) (error, models.User) {
	user := models.User{}
	result := db.GetDB().First(&user, id)
	if result.Error != nil {
		return result.Error, user
	}
	return nil, user
}

func CreateUser(id string, role string) (error, models.User) {
	user := models.User{
		Steamid64: id,
		Role:      role,
	}
	if err := db.GetDB().Create(&user).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return echo.NewHTTPError(http.StatusConflict, "User already exists"), user
			}
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error"), user
	}
	return nil, user
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
	id, err := openid.Verify(os.Getenv("BASE_URL")+c.Request().URL.String(), discoveryCache, nonceStore)
	if err != nil {
		log.Printf("Error verifying: %q\n", err)
		return c.String(http.StatusInternalServerError, "Error verifying")
	}
	steamUser := GetSteamUser(&id)
	var role string
	steamid := steamUser.Steamid
	foundUserErr, foundUser := GetUser(steamid)
	if foundUserErr != nil {
		err, created := CreateUser(steamid, "user")
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
