package gutil_test

import (
	"fmt"
	"github.com/xrcn/cg/frame/g"
	"github.com/xrcn/cg/util/gutil"
)

func ExampleSliceInsertBefore() {
	s1 := g.Slice{
		0, 1, 2, 3, 4,
	}
	s2 := gutil.SliceInsertBefore(s1, 1, 8, 9)
	fmt.Println(s1)
	fmt.Println(s2)

	// Output:
	// [0 1 2 3 4]
	// [0 8 9 1 2 3 4]
}

func ExampleSliceInsertAfter() {
	s1 := g.Slice{
		0, 1, 2, 3, 4,
	}
	s2 := gutil.SliceInsertAfter(s1, 1, 8, 9)
	fmt.Println(s1)
	fmt.Println(s2)

	// Output:
	// [0 1 2 3 4]
	// [0 1 8 9 2 3 4]
}
