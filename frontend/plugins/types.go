package plugins

import (
	"net/http"
)

//Plugin is an interface defining all plugin exported functions
type Plugin interface {
	Activate() //should register endpoints and add database tables etc
	Active() bool //should check if plugin has already been activated
}

//UniversalHandlerPlugin is an interface that extends the Plugin interface, for if a plugin needs to be called on all endpoints
type UniversalHandlerPlugin interface {
	Plugin
	UniversalHandler(http.HandlerFunc) http.HandlerFunc
}