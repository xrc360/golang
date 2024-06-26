package gfile_test

import (
	"fmt"

	"github.com/xrcn/cg/os/gfile"
)

func ExampleSearch() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_search")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goxrc example content")

	// search file
	realPath, _ := gfile.Search(fileName, tempDir)
	fmt.Println(gfile.Basename(realPath))

	// Output:
	// gfile_example.txt
}
