package testdata

import "github.com/xrcn/cg/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
