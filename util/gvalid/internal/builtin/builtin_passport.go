package builtin

import (
	"errors"

	"github.com/xrcn/cg/text/gregex"
)

// RulePassport implements `passport` rule:
// Universal passport format rule:
// Starting with letter, containing only numbers or underscores, length between 6 and 18
//
// Format: passport
type RulePassport struct{}

func init() {
	Register(RulePassport{})
}

func (r RulePassport) Name() string {
	return "passport"
}

func (r RulePassport) Message() string {
	return "The {field} value `{value}` is not a valid passport format"
}

func (r RulePassport) Run(in RunInput) error {
	ok := gregex.IsMatchString(
		`^[a-zA-Z]{1}\w{5,17}$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
