package goai

import (
	"github.com/xrcn/cg/internal/json"
)

// Example is specified by OpenAPI/Swagger 3.0 standard.
type Example struct {
	Summary       string      `json:"summary,omitempty"`
	Description   string      `json:"description,omitempty"`
	Value         interface{} `json:"value,omitempty"`
	ExternalValue string      `json:"externalValue,omitempty"`
}

type Examples map[string]*ExampleRef

type ExampleRef struct {
	Ref   string
	Value *Example
}

func (r ExampleRef) MarshalJSON() ([]byte, error) {
	if r.Ref != "" {
		return formatRefToBytes(r.Ref), nil
	}
	return json.Marshal(r.Value)
}
