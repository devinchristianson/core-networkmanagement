package root

import (
	"core-networkmanager/frontend/plugins"
	"net/http"

	"github.com/labstack/echo"
)

func setup() bool {
	plugins.RegisterEndpoint(plugins.GET, "/", homePage)
	return true
}

//Page struct to hold per-page data
type page struct {
	Location string
	Name     string
}

func homePage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
