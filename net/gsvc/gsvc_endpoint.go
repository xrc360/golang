// Package gsvc provides service registry and discovery definition.
package gsvc

import (
	"fmt"

	"github.com/xrcn/cg/errors/gcode"
	"github.com/xrcn/cg/errors/gerror"
	"github.com/xrcn/cg/text/gstr"
	"github.com/xrcn/cg/util/gconv"
)

// LocalEndpoint implements interface Endpoint.
type LocalEndpoint struct {
	host string // host can be either IPv4 or IPv6 address.
	port int    // port is port as commonly known.
}

// NewEndpoint creates and returns an Endpoint from address string of pattern "host:port",
// eg: "192.168.1.100:80".
func NewEndpoint(address string) Endpoint {
	array := gstr.SplitAndTrim(address, EndpointHostPortDelimiter)
	if len(array) != 2 {
		panic(gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid address "%s" for creating endpoint, endpoint address is like "ip:port"`,
			address,
		))
	}
	return &LocalEndpoint{
		host: array[0],
		port: gconv.Int(array[1]),
	}
}

// Host returns the IPv4/IPv6 address of a service.
func (e *LocalEndpoint) Host() string {
	return e.host
}

// Port returns the port of a service.
func (e *LocalEndpoint) Port() int {
	return e.port
}

// String formats and returns the Endpoint as a string, like: 192.168.1.100:80.
func (e *LocalEndpoint) String() string {
	return fmt.Sprintf(`%s:%d`, e.host, e.port)
}
