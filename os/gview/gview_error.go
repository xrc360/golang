package gview

import (
	"github.com/xrc360/golang/os/gcmd"
)

const (
	// commandEnvKeyForErrorPrint is used to specify the key controlling error printing to stdout.
	// This error is designed not to be returned by functions.
	commandEnvKeyForErrorPrint = "gf.gview.errorprint"
)

// errorPrint checks whether printing error to stdout.
func errorPrint() bool {
	return gcmd.GetOptWithEnv(commandEnvKeyForErrorPrint, true).Bool()
}
