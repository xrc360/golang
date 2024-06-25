package builtin

import (
	"errors"

	"github.com/xrc360/golang/text/gregex"
)

// RuleQQ implements `qq` rule:
// Tencent QQ number.
//
// Format: qq
type RuleQQ struct{}

func init() {
	Register(RuleQQ{})
}

func (r RuleQQ) Name() string {
	return "qq"
}

func (r RuleQQ) Message() string {
	return "The {field} value `{value}` is not a valid QQ number"
}

func (r RuleQQ) Run(in RunInput) error {
	ok := gregex.IsMatchString(
		`^[1-9][0-9]{4,}$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
