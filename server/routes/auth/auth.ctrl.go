package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/imroc/req/v3"
	"github.com/labstack/echo/v4"
	"github.com/yohcop/openid-go"
)

type NoOpDiscoveryCache struct{}
func (n *NoOpDiscoveryCache) Put(id string, info openid.DiscoveredInfo) {}
func (n *NoOpDiscoveryCache) Get(id string) openid.DiscoveredInfo {
	return nil
}

var nonceStore = openid.NewSimpleNonceStore()
var discoveryCache = &NoOpDiscoveryCache{}

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
		Players [] SteamUser `json:"players"`
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

func GetSteamUser (url *string) SteamUser {
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

func (AuthRouter) SteamReturn(c echo.Context) error {
	id, err := openid.Verify(os.Getenv("BASE_URL") + c.Request().URL.String(), discoveryCache, nonceStore)
	if err != nil {
		log.Printf("Error verifying: %q\n", err)
		return c.String(http.StatusInternalServerError, "Error verifying")
	}
	user := GetSteamUser(&id)
	return c.JSON(http.StatusOK, user)
}