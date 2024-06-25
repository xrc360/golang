package goai

import (
	"github.com/xrc360/golang/internal/json"
)

// Callback is specified by OpenAPI/Swagger standard version 3.0.
type Callback map[string]*Path

type Callbacks map[string]*CallbackRef

type CallbackRef struct {
	Ref   string
	Value *Callback
}

func (r CallbackRef) MarshalJSON() ([]byte, error) {
	if r.Ref != "" {
		return formatRefToBytes(r.Ref), nil
	}
	return json.Marshal(r.Value)
}
