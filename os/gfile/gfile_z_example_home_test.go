package gfile_test

import (
	"fmt"

	"github.com/xrc360/golang/os/gfile"
)

func ExampleHome() {
	// user's home directory
	homePath, _ := gfile.Home()
	fmt.Println(homePath)

	// May Output:
	// C:\Users\hailaz
}
