// go test *.go -bench=".*"

package gregex_test

import (
	"regexp"
	"testing"

	"github.com/xrcn/cg/text/gregex"
)

var pattern = `(\w+).+\-\-\s*(.+)`

var src = `CG is best! -- John`

func Benchmark_CG_IsMatchString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gregex.IsMatchString(pattern, src)
	}
}

func Benchmark_CG_MatchString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gregex.MatchString(pattern, src)
	}
}

func Benchmark_Compile(b *testing.B) {
	var wcdRegexp = regexp.MustCompile(pattern)
	for i := 0; i < b.N; i++ {
		wcdRegexp.MatchString(src)
	}
}

func Benchmark_Compile_Actual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wcdRegexp := regexp.MustCompile(pattern)
		wcdRegexp.MatchString(src)
	}
}
