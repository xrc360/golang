package gudp_test

import (
	"fmt"

	"github.com/xrc360/golang/net/gudp"
)

func ExampleGetFreePort() {
	fmt.Println(gudp.GetFreePort())

	// May Output:
	// 57429 <nil>
}

func ExampleGetFreePorts() {
	fmt.Println(gudp.GetFreePorts(2))

	// May Output:
	// [57743 57744] <nil>
}
