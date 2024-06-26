package goai

import (
	"github.com/xrcn/cg/text/gstr"
)

// XExtensions stores the `x-` custom extensions.
type XExtensions map[string]string

func (oai *OpenApiV3) tagMapToXExtensions(tagMap map[string]string, extensions XExtensions) {
	for k, v := range tagMap {
		if gstr.HasPrefix(k, "x-") || gstr.HasPrefix(k, "X-") {
			extensions[k] = v
		}
	}
}
