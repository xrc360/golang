// go test *.go -bench=".*" -benchmem

package gcmd_test

import (
	"os"
	"testing"

	"github.com/xrcn/cg/os/gcmd"
	"github.com/xrcn/cg/test/gtest"
)

func Test_Parse(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		os.Args = []string{"cg", "--force", "remove", "-fq", "-p=www", "path", "-n", "root"}
		p, err := gcmd.Parse(map[string]bool{
			"n, name":   true,
			"p, prefix": true,
			"f,force":   false,
			"q,quiet":   false,
		})
		t.AssertNil(err)
		t.Assert(len(p.GetArgAll()), 3)
		t.Assert(p.GetArg(0), "cg")
		t.Assert(p.GetArg(1), "remove")
		t.Assert(p.GetArg(2), "path")
		t.Assert(p.GetArg(2).String(), "path")

		t.Assert(len(p.GetOptAll()), 8)
		t.Assert(p.GetOpt("n"), "root")
		t.Assert(p.GetOpt("name"), "root")
		t.Assert(p.GetOpt("p"), "www")
		t.Assert(p.GetOpt("prefix"), "www")
		t.Assert(p.GetOpt("prefix").String(), "www")

		t.Assert(p.GetOpt("n") != nil, true)
		t.Assert(p.GetOpt("name") != nil, true)
		t.Assert(p.GetOpt("p") != nil, true)
		t.Assert(p.GetOpt("prefix") != nil, true)
		t.Assert(p.GetOpt("f") != nil, true)
		t.Assert(p.GetOpt("force") != nil, true)
		t.Assert(p.GetOpt("q") != nil, true)
		t.Assert(p.GetOpt("quiet") != nil, true)
		t.Assert(p.GetOpt("none") != nil, false)

		_, err = p.MarshalJSON()
		t.AssertNil(err)
	})
}

func Test_ParseArgs(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		p, err := gcmd.ParseArgs(
			[]string{"cg", "--force", "remove", "-fq", "-p=www", "path", "-n", "root"},
			map[string]bool{
				"n, name":   true,
				"p, prefix": true,
				"f,force":   false,
				"q,quiet":   false,
			})
		t.AssertNil(err)
		t.Assert(len(p.GetArgAll()), 3)
		t.Assert(p.GetArg(0), "cg")
		t.Assert(p.GetArg(1), "remove")
		t.Assert(p.GetArg(2), "path")
		t.Assert(p.GetArg(2).String(), "path")

		t.Assert(len(p.GetOptAll()), 8)
		t.Assert(p.GetOpt("n"), "root")
		t.Assert(p.GetOpt("name"), "root")
		t.Assert(p.GetOpt("p"), "www")
		t.Assert(p.GetOpt("prefix"), "www")
		t.Assert(p.GetOpt("prefix").String(), "www")

		t.Assert(p.GetOpt("n") != nil, true)
		t.Assert(p.GetOpt("name") != nil, true)
		t.Assert(p.GetOpt("p") != nil, true)
		t.Assert(p.GetOpt("prefix") != nil, true)
		t.Assert(p.GetOpt("f") != nil, true)
		t.Assert(p.GetOpt("force") != nil, true)
		t.Assert(p.GetOpt("q") != nil, true)
		t.Assert(p.GetOpt("quiet") != nil, true)
		t.Assert(p.GetOpt("none") != nil, false)
	})
}
