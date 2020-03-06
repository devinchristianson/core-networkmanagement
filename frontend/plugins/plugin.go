package plugins

import (
	"database/sql"
	"log"
	"sync"

	"github.com/labstack/echo"
)

//Global variables
var (
	pluginMu         sync.RWMutex
	plugins          = make(map[string]Plugin)
	activePlugins    = make(map[string]Plugin)
	universalPlugins []UniversalHandlerPlugin
	endpointMappings = make([]map[string]func(echo.Context) error, 4)
	mux              *echo.Echo
	db               *sql.DB
)

func init() {
	for key := range endpointMappings {
		endpointMappings[key] = make(map[string]func(echo.Context) error)
	}
}

//SetupPlugins activates plugins and loads endpoints
func SetupPlugins(m *echo.Echo, d *sql.DB, names []string) {
	mux = m
	db = d
	for _, n := range names {
		if _, exists := plugins[n]; !exists {
			log.Fatalf("Plugin %s does not exist", n)
		}
		plugins[n].Activate()
		if p, universal := plugins[n].(UniversalHandlerPlugin); universal {
			universalPlugins = append(universalPlugins, p)
		}
	}
	mux.Static("/assets", "assets ")
	for key, element := range endpointMappings[GET] {
		mux.GET(key, chainUniversalHandlers(element))
	}
	for key, element := range endpointMappings[PUT] {
		mux.GET(key, chainUniversalHandlers(element))
	}
	for key, element := range endpointMappings[POST] {
		mux.GET(key, chainUniversalHandlers(element))
	}
	for key, element := range endpointMappings[DELETE] {
		mux.GET(key, chainUniversalHandlers(element))
	}
}
func chainUniversalHandlers(h func(echo.Context) error) func(echo.Context) error {

	if len(universalPlugins) < 1 {
		return h
	}

	wrapped := h

	// loop in reverse to preserve middleware order
	for i := len(universalPlugins) - 1; i >= 0; i-- {
		wrapped = universalPlugins[i].UniversalHandler(wrapped)
	}

	return wrapped

}

//RegisterPlugin registers plugins. Should be called using init function in plugin package
func RegisterPlugin(name string, p Plugin) {
	pluginMu.Lock()
	defer pluginMu.Unlock()
	if p == nil {
		log.Fatal("Plugin is nil")
	}
	if _, duplicate := plugins[name]; duplicate {
		log.Fatal("Plugin name already taken")
	}
	plugins[name] = p
}

//RegisterEndpoint points an endpoint to a specific plugin while keeping
func RegisterEndpoint(request REQUEST, pattern string, handler func(echo.Context) error) {
	if handler == nil {
		log.Fatalf("Handler for endpoint %s is nil", pattern)
	}
	if _, duplicate := endpointMappings[request][pattern]; duplicate {
		log.Fatalf("Endpoint %s has already been allocated to a different plugin", pattern)
	}
	endpointMappings[request][pattern] = handler
}
