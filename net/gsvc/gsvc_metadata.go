package gsvc

import (
	"github.com/xrcn/cg/container/gvar"
)

// Set sets key-value pair into metadata.
func (m Metadata) Set(key string, value interface{}) {
	m[key] = value
}

// Sets sets key-value pairs into metadata.
func (m Metadata) Sets(kvs map[string]interface{}) {
	for k, v := range kvs {
		m[k] = v
	}
}

// Get retrieves and returns value of specified key as gvar.
func (m Metadata) Get(key string) *gvar.Var {
	if v, ok := m[key]; ok {
		return gvar.New(v)
	}
	return nil
}

// IsEmpty checks and returns whether current Metadata is empty.
func (m Metadata) IsEmpty() bool {
	return len(m) == 0
}
