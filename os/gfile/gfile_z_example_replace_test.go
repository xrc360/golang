package gfile_test

import (
	"fmt"
	"regexp"

	"github.com/xrcn/cg/os/gfile"
)

func ExampleReplaceFile() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_replace")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goxrc example content")

	// read contents
	fmt.Println(gfile.GetContents(tempFile))

	// It replaces content directly by file path.
	gfile.ReplaceFile("content", "replace word", tempFile)

	fmt.Println(gfile.GetContents(tempFile))

	// Output:
	// goxrc example content
	// goxrc example replace word
}

func ExampleReplaceFileFunc() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_replace")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goxrc example 123")

	// read contents
	fmt.Println(gfile.GetContents(tempFile))

	// It replaces content directly by file path and callback function.
	gfile.ReplaceFileFunc(func(path, content string) string {
		// Replace with regular match
		reg, _ := regexp.Compile(`\d{3}`)
		return reg.ReplaceAllString(content, "[num]")
	}, tempFile)

	fmt.Println(gfile.GetContents(tempFile))

	// Output:
	// goxrc example 123
	// goxrc example [num]
}

func ExampleReplaceDir() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_replace")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goxrc example content")

	// read contents
	fmt.Println(gfile.GetContents(tempFile))

	// It replaces content of all files under specified directory recursively.
	gfile.ReplaceDir("content", "replace word", tempDir, "gfile_example.txt", true)

	// read contents
	fmt.Println(gfile.GetContents(tempFile))

	// Output:
	// goxrc example content
	// goxrc example replace word
}

func ExampleReplaceDirFunc() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_replace")
		tempFile = gfile.Join(tempDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goxrc example 123")

	// read contents
	fmt.Println(gfile.GetContents(tempFile))

	// It replaces content of all files under specified directory with custom callback function recursively.
	gfile.ReplaceDirFunc(func(path, content string) string {
		// Replace with regular match
		reg, _ := regexp.Compile(`\d{3}`)
		return reg.ReplaceAllString(content, "[num]")
	}, tempDir, "gfile_example.txt", true)

	fmt.Println(gfile.GetContents(tempFile))

	// Output:
	// goxrc example 123
	// goxrc example [num]

}
