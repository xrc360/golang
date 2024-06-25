package goai

import (
	"github.com/xrc360/golang/errors/gerror"
	"github.com/xrc360/golang/internal/json"
	"github.com/xrc360/golang/util/gconv"
)

// Parameter is specified by OpenAPI/Swagger 3.0 standard.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#parameterObject
type Parameter struct {
	Name            string      `json:"name,omitempty"`
	In              string      `json:"in,omitempty"`
	Description     string      `json:"description,omitempty"`
	Style           string      `json:"style,omitempty"`
	Explode         *bool       `json:"explode,omitempty"`
	AllowEmptyValue bool        `json:"allowEmptyValue,omitempty"`
	AllowReserved   bool        `json:"allowReserved,omitempty"`
	Deprecated      bool        `json:"deprecated,omitempty"`
	Required        bool        `json:"required,omitempty"`
	Schema          *SchemaRef  `json:"schema,omitempty"`
	Example         interface{} `json:"example,omitempty"`
	Examples        *Examples   `json:"examples,omitempty"`
	Content         *Content    `json:"content,omitempty"`
	XExtensions     XExtensions `json:"-"`
}

func (oai *OpenApiV3) tagMapToParameter(tagMap map[string]string, parameter *Parameter) error {
	var mergedTagMap = oai.fillMapWithShortTags(tagMap)
	if err := gconv.Struct(mergedTagMap, parameter); err != nil {
		return gerror.Wrap(err, `mapping struct tags to Parameter failed`)
	}
	oai.tagMapToXExtensions(mergedTagMap, parameter.XExtensions)
	return nil
}

func (p Parameter) MarshalJSON() ([]byte, error) {
	var (
		b   []byte
		m   map[string]json.RawMessage
		err error
	)
	type tempParameter Parameter // To prevent JSON marshal recursion error.
	if b, err = json.Marshal(tempParameter(p)); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	for k, v := range p.XExtensions {
		if b, err = json.Marshal(v); err != nil {
			return nil, err
		}
		m[k] = b
	}
	return json.Marshal(m)
}
