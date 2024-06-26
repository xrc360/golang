package goai

import (
	"fmt"
	"net/http"

	"github.com/xrcn/cg/container/gset"
	"github.com/xrcn/cg/errors/gcode"
	"github.com/xrcn/cg/errors/gerror"
	"github.com/xrcn/cg/internal/json"
	"github.com/xrcn/cg/os/gstructs"
	"github.com/xrcn/cg/text/gstr"
	"github.com/xrcn/cg/util/gconv"
)

// Parameters is specified by OpenAPI/Swagger 3.0 standard.
type Parameters []ParameterRef

type ParameterRef struct {
	Ref   string
	Value *Parameter
}

func (oai *OpenApiV3) newParameterRefWithStructMethod(field gstructs.Field, path, method string) (*ParameterRef, error) {
	var (
		tagMap    = field.TagMap()
		fieldName = field.Name()
	)
	for _, tagName := range gconv.StructTagPriority {
		if tagValue := field.Tag(tagName); tagValue != "" {
			fieldName = tagValue
			break
		}
	}
	fieldName = gstr.Split(gstr.Trim(fieldName), ",")[0]
	if fieldName == "" {
		fieldName = field.Name()
	}
	var parameter = &Parameter{
		Name:        fieldName,
		XExtensions: make(XExtensions),
	}
	if len(tagMap) > 0 {
		if err := oai.tagMapToParameter(tagMap, parameter); err != nil {
			return nil, err
		}
	}
	if parameter.In == "" {
		// Automatically detect its "in" attribute.
		if gstr.ContainsI(path, fmt.Sprintf(`{%s}`, parameter.Name)) {
			parameter.In = ParameterInPath
		} else {
			// Default the parameter input to "query" if method is "GET/DELETE".
			switch gstr.ToUpper(method) {
			case http.MethodGet, http.MethodDelete:
				parameter.In = ParameterInQuery

			default:
				return nil, nil
			}
		}
	}

	switch parameter.In {
	case ParameterInPath:
		// Required for path parameter.
		parameter.Required = true

	case ParameterInCookie, ParameterInHeader, ParameterInQuery:

	default:
		return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `invalid tag value "%s" for In`, parameter.In)
	}
	// Necessary schema or content.
	schemaRef, err := oai.newSchemaRefWithGolangType(field.Type().Type, tagMap)
	if err != nil {
		return nil, err
	}
	parameter.Schema = schemaRef

	// Ignore parameter.
	if !isValidParameterName(parameter.Name) {
		return nil, nil
	}

	// Required check.
	if parameter.Schema.Value != nil && parameter.Schema.Value.ValidationRules != "" {
		validationRuleArray := gstr.Split(parameter.Schema.Value.ValidationRules, "|")
		if gset.NewStrSetFrom(validationRuleArray).Contains(validationRuleKeyForRequired) {
			parameter.Required = true
		}
	}

	return &ParameterRef{
		Ref:   "",
		Value: parameter,
	}, nil
}

func (r ParameterRef) MarshalJSON() ([]byte, error) {
	if r.Ref != "" {
		return formatRefToBytes(r.Ref), nil
	}
	return json.Marshal(r.Value)
}
