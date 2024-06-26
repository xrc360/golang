package builtin

import (
	"errors"

	"github.com/xrcn/cg/net/gipv4"
	"github.com/xrcn/cg/net/gipv6"
)

// RuleIp implements `ip` rule:
// IPv4/IPv6.
//
// Format: ip
type RuleIp struct{}

func init() {
	Register(RuleIp{})
}

func (r RuleIp) Name() string {
	return "ip"
}

func (r RuleIp) Message() string {
	return "The {field} value `{value}` is not a valid IP address"
}

func (r RuleIp) Run(in RunInput) error {
	var (
		ok    bool
		value = in.Value.String()
	)
	if ok = gipv4.Validate(value) || gipv6.Validate(value); !ok {
		return errors.New(in.Message)
	}
	return nil
}
