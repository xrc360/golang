package builtin

import (
	"errors"

	"github.com/xrc360/golang/text/gstr"
	"github.com/xrc360/golang/util/gconv"
	"github.com/xrc360/golang/util/gutil"
)

// RuleBeforeEqual implements `before-equal` rule:
// The datetime value should be after or equal to the value of field `field`.
//
// Format: before-equal:field
type RuleBeforeEqual struct{}

func init() {
	Register(RuleBeforeEqual{})
}

func (r RuleBeforeEqual) Name() string {
	return "before-equal"
}

func (r RuleBeforeEqual) Message() string {
	return "The {field} value `{value}` must be before or equal to field {pattern}"
}

func (r RuleBeforeEqual) Run(in RunInput) error {
	var (
		fieldName, fieldValue = gutil.MapPossibleItemByKey(in.Data.Map(), in.RulePattern)
		valueDatetime         = in.Value.Time()
		fieldDatetime         = gconv.Time(fieldValue)
	)
	if valueDatetime.Before(fieldDatetime) || valueDatetime.Equal(fieldDatetime) {
		return nil
	}
	return errors.New(gstr.ReplaceByMap(in.Message, map[string]string{
		"{field1}": fieldName,
		"{value1}": gconv.String(fieldValue),
	}))
}
