package builtin

import (
	"errors"
	"strings"

	"github.com/xrcn/cg/internal/empty"
	"github.com/xrcn/cg/util/gutil"
)

// RuleRequiredWith implements `required-with` rule:
// Required if any of given fields are not empty.
//
// Format:  required-with:field1,field2,...
// Example: required-with:id,name
type RuleRequiredWith struct{}

func init() {
	Register(RuleRequiredWith{})
}

func (r RuleRequiredWith) Name() string {
	return "required-with"
}

func (r RuleRequiredWith) Message() string {
	return "The {field} field is required"
}

func (r RuleRequiredWith) Run(in RunInput) error {
	var (
		required   = false
		array      = strings.Split(in.RulePattern, ",")
		foundValue interface{}
	)
	for i := 0; i < len(array); i++ {
		_, foundValue = gutil.MapPossibleItemByKey(in.Data.Map(), array[i])
		if !empty.IsEmpty(foundValue) {
			required = true
			break
		}
	}

	if required && isRequiredEmpty(in.Value.Val()) {
		return errors.New(in.Message)
	}
	return nil
}
