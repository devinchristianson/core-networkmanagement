package root
import (
	"sync"
	"core-networkmanager/frontend/plugins"
)

var (
	mux sync.Mutex
	active bool
)

//init registers the plugin as available with Plugin manager
func init() {
	plugins.RegisterPlugin("root", &plugin{})
}

//plugin struct that for the plugin to impliment required and optional interfaces as defined in Plugin package
type plugin struct {}

//Activate sets up endpoints and databases, and any other initialization 
func (p plugin ) Activate () {
	if(!active) {
		mux.Lock()
		active = setup()
		mux.Unlock()
	}
}

//Active returns true if the plugin has already been activated
func (p plugin ) Active () bool {
	return active
}