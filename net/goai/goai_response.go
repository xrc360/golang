package goai

import (
	"github.com/xrc360/golang/errors/gerror"
	"github.com/xrc360/golang/internal/json"
	"github.com/xrc360/golang/util/gconv"
)

// Response is specified by OpenAPI/Swagger 3.0 standard.
type Response struct {
	Description string      `json:"description"`
	Headers     Headers     `json:"headers,omitempty"`
	Content     Content     `json:"content,omitempty"`
	Links       Links       `json:"links,omitempty"`
	XExtensions XExtensions `json:"-"`
}

func (oai *OpenApiV3) tagMapToResponse(tagMap map[string]string, response *Response) error {
	var mergedTagMap = oai.fillMapWithShortTags(tagMap)
	if err := gconv.Struct(mergedTagMap, response); err != nil {
		return gerror.Wrap(err, `mapping struct tags to Response failed`)
	}
	oai.tagMapToXExtensions(mergedTagMap, response.XExtensions)
	return nil
}

func (r Response) MarshalJSON() ([]byte, error) {
	var (
		b   []byte
		m   map[string]json.RawMessage
		err error
	)
	type tempResponse Response // To prevent JSON marshal recursion error.
	if b, err = json.Marshal(tempResponse(r)); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	for k, v := range r.XExtensions {
		if b, err = json.Marshal(v); err != nil {
			return nil, err
		}
		m[k] = b
	}
	return json.Marshal(m)
}
