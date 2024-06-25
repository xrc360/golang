package goai

import (
	"github.com/xrc360/golang/internal/json"
)

// Header is specified by OpenAPI/Swagger 3.0 standard.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#headerObject
type Header struct {
	Parameter
}

type Headers map[string]HeaderRef

type HeaderRef struct {
	Ref   string
	Value *Header
}

func (r HeaderRef) MarshalJSON() ([]byte, error) {
	if r.Ref != "" {
		return formatRefToBytes(r.Ref), nil
	}
	return json.Marshal(r.Value)
}
