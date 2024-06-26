package utils_test

import (
	"regexp"
	"testing"

	"github.com/xrcn/cg/internal/utils"
)

var (
	replaceCharReg, _ = regexp.Compile(`[\-\.\_\s]+`)
)

func Benchmark_RemoveSymbols(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.RemoveSymbols(`-a-b._a c1!@#$%^&*()_+:";'.,'01`)
	}
}

func Benchmark_RegularReplaceChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		replaceCharReg.ReplaceAllString(`-a-b._a c1!@#$%^&*()_+:";'.,'01`, "")
	}
}
