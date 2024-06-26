package gfile_test

import (
	"fmt"
	"time"

	"github.com/xrcn/cg/os/gfile"
)

func ExampleGetContentsWithCache() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_cache")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goxrc example content")

	// It reads the file content with cache duration of one minute,
	// which means it reads from cache after then without any IO operations within on minute.
	fmt.Println(gfile.GetContentsWithCache(tempFile, time.Minute))

	// write new contents will clear its cache
	gfile.PutContents(tempFile, "new goxrc example content")

	// There's some delay for cache clearing after file content change.
	time.Sleep(time.Second * 1)

	// read contents
	fmt.Println(gfile.GetContentsWithCache(tempFile))

	// May Output:
	// goxrc example content
	// new goxrc example content
}
