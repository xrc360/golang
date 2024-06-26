package gfile_test

import (
	"fmt"

	"github.com/xrcn/cg/os/gfile"
)

func ExampleCopy() {
	// init
	var (
		srcFileName = "gfile_example.txt"
		srcTempDir  = gfile.Temp("gfile_example_copy_src")
		srcTempFile = gfile.Join(srcTempDir, srcFileName)

		// copy file
		dstFileName = "gfile_example_copy.txt"
		dstTempFile = gfile.Join(srcTempDir, dstFileName)

		// copy dir
		dstTempDir = gfile.Temp("gfile_example_copy_dst")
	)

	// write contents
	gfile.PutContents(srcTempFile, "goxrc example copy")

	// copy file
	gfile.Copy(srcTempFile, dstTempFile)

	// read contents after copy file
	fmt.Println(gfile.GetContents(dstTempFile))

	// copy dir
	gfile.Copy(srcTempDir, dstTempDir)

	// list copy dir file
	fList, _ := gfile.ScanDir(dstTempDir, "*", false)
	for _, v := range fList {
		fmt.Println(gfile.Basename(v))
	}

	// Output:
	// goxrc example copy
	// gfile_example.txt
	// gfile_example_copy.txt
}
