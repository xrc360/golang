package gfile_test

import (
	"fmt"

	"github.com/xrcn/cg/os/gfile"
)

func ExampleScanDir() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_scan_dir")
		tempFile = gfile.Join(tempDir, fileName)

		tempSubDir  = gfile.Join(tempDir, "sub_dir")
		tempSubFile = gfile.Join(tempSubDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goxrc example content")
	gfile.PutContents(tempSubFile, "goxrc example content")

	// scans directory recursively
	list, _ := gfile.ScanDir(tempDir, "*", true)
	for _, v := range list {
		fmt.Println(gfile.Basename(v))
	}

	// Output:
	// gfile_example.txt
	// sub_dir
	// gfile_example.txt
}

func ExampleScanDirFile() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_scan_dir_file")
		tempFile = gfile.Join(tempDir, fileName)

		tempSubDir  = gfile.Join(tempDir, "sub_dir")
		tempSubFile = gfile.Join(tempSubDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goxrc example content")
	gfile.PutContents(tempSubFile, "goxrc example content")

	// scans directory recursively exclusive of directories
	list, _ := gfile.ScanDirFile(tempDir, "*.txt", true)
	for _, v := range list {
		fmt.Println(gfile.Basename(v))
	}

	// Output:
	// gfile_example.txt
	// gfile_example.txt
}

func ExampleScanDirFunc() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_scan_dir_func")
		tempFile = gfile.Join(tempDir, fileName)

		tempSubDir  = gfile.Join(tempDir, "sub_dir")
		tempSubFile = gfile.Join(tempSubDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goxrc example content")
	gfile.PutContents(tempSubFile, "goxrc example content")

	// scans directory recursively
	list, _ := gfile.ScanDirFunc(tempDir, "*", true, func(path string) string {
		// ignores some files
		if gfile.Basename(path) == "gfile_example.txt" {
			return ""
		}
		return path
	})
	for _, v := range list {
		fmt.Println(gfile.Basename(v))
	}

	// Output:
	// sub_dir
}

func ExampleScanDirFileFunc() {
	// init
	var (
		fileName = "gfile_example.txt"
		tempDir  = gfile.Temp("gfile_example_scan_dir_file_func")
		tempFile = gfile.Join(tempDir, fileName)

		fileName1 = "gfile_example_ignores.txt"
		tempFile1 = gfile.Join(tempDir, fileName1)

		tempSubDir  = gfile.Join(tempDir, "sub_dir")
		tempSubFile = gfile.Join(tempSubDir, fileName)
	)

	// write contents
	gfile.PutContents(tempFile, "goxrc example content")
	gfile.PutContents(tempFile1, "goxrc example content")
	gfile.PutContents(tempSubFile, "goxrc example content")

	// scans directory recursively exclusive of directories
	list, _ := gfile.ScanDirFileFunc(tempDir, "*.txt", true, func(path string) string {
		// ignores some files
		if gfile.Basename(path) == "gfile_example_ignores.txt" {
			return ""
		}
		return path
	})
	for _, v := range list {
		fmt.Println(gfile.Basename(v))
	}

	// Output:
	// gfile_example.txt
	// gfile_example.txt
}
