package gins

import (
	"github.com/xrcn/cg/i18n/gi18n"
)

// I18n returns an instance of gi18n.Manager.
// The parameter `name` is the name for the instance.
func I18n(name ...string) *gi18n.Manager {
	return gi18n.Instance(name...)
}
