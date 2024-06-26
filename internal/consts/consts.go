// Package consts defines constants that are shared all among packages of framework.
package consts

const (
	ConfigNodeNameDatabase        = "database"
	ConfigNodeNameLogger          = "logger"
	ConfigNodeNameRedis           = "redis"
	ConfigNodeNameViewer          = "viewer"
	ConfigNodeNameServer          = "server"     // General version configuration item name.
	ConfigNodeNameServerSecondary = "httpserver" // New version configuration item name support from v2.

	// StackFilterKeyForGoXrc is the stack filtering key for all GoXrc module paths.
	// Eg: .../pkg/mod/github.com/xrcn/cg/v2@v2.0.0-20211011134327-54dd11f51122/debug/gdebug/gdebug_caller.go
	StackFilterKeyForGoXrc = "github.com/xrcn/cg/"
)
