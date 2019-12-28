package plugins

import (
	"log"
	"sync"
	"net/http"
)

var (
	pluginMu sync.RWMutex
	plugins   = make(map[string]Plugin)
	activePlugins = make(map[string]Plugin)
	universalPlugins = make(map[string]UniversalHandlerPlugin)
	endpointMappings = make(map[string]func(http.ResponseWriter, *http.Request))
	mux *http.ServeMux
)

//SetupPlugins activates plugins and loads endpoints
func SetupPlugins(m *http.ServeMux, names[] string) {
	for _, n := range names {
		if _, exists := plugins[n]; !exists {
			log.Fatalf("Plugin %s does not exist", n)
		}
		plugins[n].Activate()
		if _, universal := plugins[n].(UniversalHandlerPlugin); universal {
			universalPlugins[n] = plugins[n].(UniversalHandlerPlugin)
		}
	}
	mux = m
	for key, element := range endpointMappings {
        mux.HandleFunc(key, element)
    }
}

//Plugin is an interface defining all plugin exported functions
type Plugin interface {
	Activate() //should register endpoints and add database tables etc
}

//UniversalHandlerPlugin is an interface that extends the Plugin interface, for if a plugin needs to be called on all endpoints
type UniversalHandlerPlugin interface {
	Plugin
	UniversalHandler(http.HandlerFunc) http.HandlerFunc
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