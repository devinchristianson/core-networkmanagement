package plugins

import (
	"net/http"
	"net"
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
//HostHandlerPlugin is an interface that extends the Plugin interface, in order to modify Hosts
type HostHandlerPlugin interface {
	Plugin
	HostHandler(CoreHost, CoreOption) (CoreHost, CoreOption)
}
//DomainHandlerPlugin is an interface that extends the Plugin interface, in order to modify Domains
type DomainHandlerPlugin interface {
	Plugin
	DomainHandler(CoreDomain, CoreOption) (CoreDomain, CoreOption)
}
//NetworkHandlerPlugin is an interface that extends the Plugin interface, in order to modify Networks
type NetworkHandlerPlugin interface {
	Plugin
	NetworkHandler(CoreNetwork, CoreOption) (CoreNetwork, CoreOption)
}
//CoreHost is the struct used to pass host data between plugins and the plugin package
type CoreHost struct {
	ips [] net.IP
	name string
	domain CoreDomain
	options map[string]CoreOption
}
//CoreDomain is the struct used to pass Domain data between plugins and the plugin package
type CoreDomain struct {
	domain string
	tld string
	options map[string]CoreOption
}
//CoreNetwork is the struct used to pass Network data between plugins and the plugin package
type CoreNetwork struct {
	network net.IPNet
	options map[string]CoreOption
}
//CoreOption is the struct used to  pass Network data between plugins and the plugin package
type CoreOption struct {
	s string
	i int 
}