package steam

import (
	"fmt"
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

func parseSteamUrl(url *string) string {
	return strings.Replace(*url, "http://steamcommunity.com/openid/id/", "", 1)
}

func VerifyURL(c echo.Context) (string, error) {
	return openid.Verify(os.Getenv("BASE_URL")+c.Request().URL.String(), discoveryCache, nonceStore)
}

func QuerySteamUser(url *string) SteamUser {
	response := SteamApiResponse{}
	steamId := parseSteamUrl(url)

	req.
		SetResult(&response).
		Get(fmt.Sprintf("https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", os.Getenv("STEAM_KEY"), steamId))
	return response.Response.Players[0]
}

func GenerateUrl() (string, error) {
	return openid.RedirectURL(
		"http://steamcommunity.com/openid",
		os.Getenv("STEAM_RETURN_URL"),
		os.Getenv("BASE_URL"))
}

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
