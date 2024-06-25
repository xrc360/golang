package gtag_test

import (
	"fmt"

	"github.com/xrc360/golang/frame/g"
	"github.com/xrc360/golang/util/gmeta"
	"github.com/xrc360/golang/util/gtag"
)

func ExampleSet() {
	type User struct {
		g.Meta `name:"User Struct" description:"{UserDescription}"`
	}
	gtag.Sets(g.MapStrStr{
		`UserDescription`: `This is a demo struct named "User Struct"`,
	})
	fmt.Println(gmeta.Get(User{}, `description`))

	// Output:
	// This is a demo struct named "User Struct"
}
