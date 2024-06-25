package gins

import (
	"github.com/xrc360/golang/os/gres"
)

// Resource returns an instance of Resource.
// The parameter `name` is the name for the instance.
func Resource(name ...string) *gres.Resource {
	return gres.Instance(name...)
}
