package goai

import (
	"github.com/xrcn/cg/internal/json"
	"github.com/xrcn/cg/text/gstr"
	"github.com/xrcn/cg/util/gconv"
)

// ExternalDocs is specified by OpenAPI/Swagger standard version 3.0.
type ExternalDocs struct {
	URL         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
}

func (ed *ExternalDocs) UnmarshalValue(value interface{}) error {
	var valueBytes = gconv.Bytes(value)
	if json.Valid(valueBytes) {
		return json.UnmarshalUseNumber(valueBytes, ed)
	}
	var (
		valueString = string(valueBytes)
		valueArray  = gstr.Split(valueString, "|")
	)
	ed.URL = valueArray[0]
	if len(valueArray) > 1 {
		ed.Description = valueArray[1]
	}
	return nil
}
