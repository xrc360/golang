package builtin

import (
	"errors"

	"github.com/xrcn/cg/net/gipv4"
)

// RuleIpv4 implements `ipv4` rule:
// IPv4.
//
// Format: ipv4
type RuleIpv4 struct{}

func init() {
	Register(RuleIpv4{})
}

func (r RuleIpv4) Name() string {
	return "ipv4"
}

func (r RuleIpv4) Message() string {
	return "The {field} value `{value}` is not a valid IPv4 address"
}

func (r RuleIpv4) Run(in RunInput) error {
	var (
		ok    bool
		value = in.Value.String()
	)
	if ok = gipv4.Validate(value); !ok {
		return errors.New(in.Message)
	}
	return nil
}
