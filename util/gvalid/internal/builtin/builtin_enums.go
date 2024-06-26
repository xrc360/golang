package builtin

import (
	"errors"
	"fmt"

	"github.com/xrcn/cg/errors/gcode"
	"github.com/xrcn/cg/errors/gerror"
	"github.com/xrcn/cg/internal/json"
	"github.com/xrcn/cg/text/gstr"
	"github.com/xrcn/cg/util/gconv"
	"github.com/xrcn/cg/util/gtag"
)

// RuleEnums implements `enums` rule:
// Value should be in enums of its constant type.
//
// Format: enums
type RuleEnums struct{}

func init() {
	Register(RuleEnums{})
}

func (r RuleEnums) Name() string {
	return "enums"
}

func (r RuleEnums) Message() string {
	return "The {field} value `{value}` should be in enums of: {enums}"
}

func (r RuleEnums) Run(in RunInput) error {
	if in.ValueType == nil {
		return gerror.NewCode(
			gcode.CodeInvalidParameter,
			`value type cannot be empty to use validation rule "enums"`,
		)
	}
	var (
		pkgPath  = in.ValueType.PkgPath()
		typeName = in.ValueType.Name()
	)
	if pkgPath == "" {
		return gerror.NewCodef(
			gcode.CodeInvalidOperation,
			`no pkg path found for type "%s"`,
			in.ValueType.String(),
		)
	}
	var (
		typeId   = fmt.Sprintf(`%s.%s`, pkgPath, typeName)
		tagEnums = gtag.GetEnumsByType(typeId)
	)
	if tagEnums == "" {
		return gerror.NewCodef(
			gcode.CodeInvalidOperation,
			`no enums found for type "%s", missing using command "cg gen enums"?`,
			typeId,
		)
	}
	var enumsValues = make([]interface{}, 0)
	if err := json.Unmarshal([]byte(tagEnums), &enumsValues); err != nil {
		return err
	}
	if !gstr.InArray(gconv.Strings(enumsValues), in.Value.String()) {
		return errors.New(gstr.Replace(
			in.Message, `{enums}`, tagEnums,
		))
	}
	return nil
}
