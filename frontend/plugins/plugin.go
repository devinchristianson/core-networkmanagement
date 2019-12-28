package plugins

import (
	"log"
	"sync"
	"net/http"
)

var (
	pluginMu sync.RWMutex
	Plugins   = make(map[string]Plugin)
	universalPlugins = make(map[string]UniversalHandlerPlugin)
	mux http.ServeMux
)

func Setup (mux *http.ServeMux) {
	mux = mux
}

//Plugin is an interface defining all plugin exported functions
type Plugin interface {
	Setup(*http.ServeMux) //should register endpoints and add database tables etc
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
	if _, duplicate := Plugins[name]; duplicate {
		log.Fatal("Plugin name already taken")
	}
	if _, universal := Plugins[name].(UniversalHandlerPlugin); universal {
		universalPlugins[name] = p.(UniversalHandlerPlugin)
	}
	Plugins[name] = p 
}

func RegisterEndpoint (pattern string, handler func(http.ResponseWriter, *http.Request)) {

}