package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
func main() {
	const (
		host   = "roach"
		dbUser = "root"
		dbPort = "26257"
	)
	openDB(host, dbUser, dbPort, "postgres", "")
	addHost()
	var port = "8080"
	mux := echo.New()
	//activePlugins := []string{}
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.gohtml")),
	}
	mux.Use(middleware.Logger())
	mux.Use(middleware.Recover())
	mux.Static("/static", "static")
	mux.GET("/", homePage)
	mux.PUT("/", login)

	mux.Renderer = renderer
	//plugins.SetupPlugins(mux, nil, activePlugins)
	mux.Logger.Fatal(mux.Start(":" + port))
}
func homePage(c echo.Context) error {

	return c.Render(http.StatusOK, "index", nil)
}
