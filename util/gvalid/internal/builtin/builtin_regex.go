package builtin

import (
	"errors"

	"github.com/xrcn/cg/text/gregex"
)

// RuleRegex implements `regex` rule:
// Value should match custom regular expression pattern.
//
// Format: regex:pattern
type RuleRegex struct{}

func init() {
	Register(RuleRegex{})
}

func (r RuleRegex) Name() string {
	return "regex"
}

func (r RuleRegex) Message() string {
	return "The {field} value `{value}` must be in regex of: {pattern}"
}

func (r RuleRegex) Run(in RunInput) error {
	if !gregex.IsMatchString(in.RulePattern, in.Value.String()) {
		return errors.New(in.Message)
	}
	return nil
}
