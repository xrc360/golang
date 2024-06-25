package gins

import (
	"fmt"

	"github.com/xrc360/golang/internal/instance"
	"github.com/xrc360/golang/net/gclient"
)

// HttpClient returns an instance of http client with specified name.
func HttpClient(name ...interface{}) *gclient.Client {
	var instanceKey = fmt.Sprintf("%s.%v", frameCoreComponentNameHttpClient, name)
	return instance.GetOrSetFuncLock(instanceKey, func() interface{} {
		return gclient.New()
	}).(*gclient.Client)
}
