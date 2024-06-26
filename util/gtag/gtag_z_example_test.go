package gtag_test

import (
	"fmt"

	"github.com/xrcn/cg/frame/g"
	"github.com/xrcn/cg/util/gmeta"
	"github.com/xrcn/cg/util/gtag"
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
