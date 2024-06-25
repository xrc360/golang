// Package gsession implements manager and storage features for sessions.
package gsession

import (
	"github.com/xrc360/golang/errors/gcode"
	"github.com/xrc360/golang/errors/gerror"
	"github.com/xrc360/golang/util/guid"
)

var (
	// ErrorDisabled is used for marking certain interface function not used.
	ErrorDisabled = gerror.NewWithOption(gerror.Option{
		Text: "this feature is disabled in this storage",
		Code: gcode.CodeNotSupported,
	})
)

// NewSessionId creates and returns a new and unique session id string,
// which is in 32 bytes.
func NewSessionId() string {
	return guid.S()
}
