package auth

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/yohcop/openid-go"
)

// NoOpDiscoveryCache implements the DiscoveryCache interface and doesn't cache anything.
// For a simple website, I'm not sure you need a cache.
type NoOpDiscoveryCache struct{}

// Put is a no op.
func (n *NoOpDiscoveryCache) Put(id string, info openid.DiscoveredInfo) {}

// Get always returns nil.
func (n *NoOpDiscoveryCache) Get(id string) openid.DiscoveredInfo {
	return nil
}

var nonceStore = openid.NewSimpleNonceStore()
var discoveryCache = &NoOpDiscoveryCache{}

func (AuthRouter) Login(c echo.Context) error {

	url, err := openid.RedirectURL(
		"http://steamcommunity.com/openid",
		os.Getenv("STEAM_RETURN_URL"),
		os.Getenv("BASE_URL"))

	if err != nil {
		log.Printf("Error creating redirect URL: %q\n", err)
		return c.String(http.StatusInternalServerError, "Error creating redirect URL")
	} else {
		return c.Redirect(303, url)
	}
}

func (AuthRouter) SteamReturn(c echo.Context) error {
	id, err := openid.Verify(os.Getenv("BASE_URL") + c.Request().URL.String(), discoveryCache, nonceStore)
	if err != nil {
		log.Printf("Error verifying: %q\n", err)
		return c.String(http.StatusInternalServerError, "Error verifying")
	} else {
		log.Printf("NonceStore: %+v\n", nonceStore)
		data := make(map[string]string)
		data["user"] = id
		return c.JSON(http.StatusOK, data)
	}
}