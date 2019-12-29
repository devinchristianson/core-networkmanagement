package plugins

import (
	"log"
	"sync"
	"net/http"
	"database/sql"
)

//Global variables
var (
	pluginMu sync.RWMutex
	plugins   = make(map[string]Plugin)
	activePlugins = make(map[string]Plugin)
	universalPlugins [] UniversalHandlerPlugin
	endpointMappings = make(map[string]func(http.ResponseWriter, *http.Request))
	mux *http.ServeMux
	db *sql.DB
)

//SetupPlugins activates plugins and loads endpoints
func SetupPlugins(m *http.ServeMux, d *sql.DB, names[] string) {
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
	for key, element := range endpointMappings {
        mux.HandleFunc(key, chainUniversalHandlers(element))
    }
}
func chainUniversalHandlers(h http.HandlerFunc) http.HandlerFunc {

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

//RegisterPlugin registers plugins. Should be called using init in plugin package
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
func RegisterEndpoint(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	if handler == nil {
		log.Fatalf("Handler for endpoint %s is nil", pattern)
	}
	if _, duplicate := endpointMappings[pattern]; duplicate {
		log.Fatalf("Endpoint %s has already been allocated to a different plugin", pattern)
	}
	endpointMappings[pattern] = handler
}