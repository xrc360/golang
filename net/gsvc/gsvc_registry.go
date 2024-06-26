package gsvc

import (
	"context"

	"github.com/xrcn/cg/errors/gcode"
	"github.com/xrcn/cg/errors/gerror"
)

// Register registers `service` to default registry..
func Register(ctx context.Context, service Service) (Service, error) {
	if defaultRegistry == nil {
		return nil, gerror.NewCodef(gcode.CodeNotImplemented, `no Registry is registered`)
	}
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	return defaultRegistry.Register(ctx, service)
}

// Deregister removes `service` from default registry.
func Deregister(ctx context.Context, service Service) error {
	if defaultRegistry == nil {
		return gerror.NewCodef(gcode.CodeNotImplemented, `no Registry is registered`)
	}
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	return defaultRegistry.Deregister(ctx, service)
}
