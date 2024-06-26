package gcmd_test

import (
	"context"
	"fmt"
	"os"

	"github.com/xrcn/cg/frame/g"
	"github.com/xrcn/cg/os/gcmd"
	"github.com/xrcn/cg/os/gctx"
	"github.com/xrcn/cg/os/genv"
)

func ExampleInit() {
	gcmd.Init("cg", "build", "main.go", "-o=cg.exe", "-y")
	fmt.Printf(`%#v`, gcmd.GetArgAll())

	// Output:
	// []string{"cg", "build", "main.go"}
}

func ExampleGetArg() {
	gcmd.Init("cg", "build", "main.go", "-o=cg.exe", "-y")
	fmt.Printf(
		`Arg[0]: "%v", Arg[1]: "%v", Arg[2]: "%v", Arg[3]: "%v"`,
		gcmd.GetArg(0), gcmd.GetArg(1), gcmd.GetArg(2), gcmd.GetArg(3),
	)

	// Output:
	// Arg[0]: "cg", Arg[1]: "build", Arg[2]: "main.go", Arg[3]: ""
}

func ExampleGetArgAll() {
	gcmd.Init("cg", "build", "main.go", "-o=cg.exe", "-y")
	fmt.Printf(`%#v`, gcmd.GetArgAll())

	// Output:
	// []string{"cg", "build", "main.go"}
}

func ExampleGetOpt() {
	gcmd.Init("cg", "build", "main.go", "-o=cg.exe", "-y")
	fmt.Printf(
		`Opt["o"]: "%v", Opt["y"]: "%v", Opt["d"]: "%v"`,
		gcmd.GetOpt("o"), gcmd.GetOpt("y"), gcmd.GetOpt("d", "default value"),
	)

	// Output:
	// Opt["o"]: "cg.exe", Opt["y"]: "", Opt["d"]: "default value"
}

func ExampleGetOpt_Def() {
	gcmd.Init("cg", "build", "main.go", "-o=cg.exe", "-y")

	fmt.Println(gcmd.GetOpt("s", "Def").String())

	// Output:
	// Def
}

func ExampleGetOptAll() {
	gcmd.Init("cg", "build", "main.go", "-o=cg.exe", "-y")
	fmt.Printf(`%#v`, gcmd.GetOptAll())

	// May Output:
	// map[string]string{"o":"cg.exe", "y":""}
}

func ExampleGetOptWithEnv() {
	fmt.Printf("Opt[cg.test]:%s\n", gcmd.GetOptWithEnv("cg.test"))
	_ = genv.Set("CG_TEST", "YES")
	fmt.Printf("Opt[cg.test]:%s\n", gcmd.GetOptWithEnv("cg.test"))

	// Output:
	// Opt[cg.test]:
	// Opt[cg.test]:YES
}

func ExampleParse() {
	os.Args = []string{"cg", "build", "main.go", "-o=cg.exe", "-y"}
	p, err := gcmd.Parse(g.MapStrBool{
		"o,output": true,
		"y,yes":    false,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(p.GetOpt("o"))
	fmt.Println(p.GetOpt("output"))
	fmt.Println(p.GetOpt("y") != nil)
	fmt.Println(p.GetOpt("yes") != nil)
	fmt.Println(p.GetOpt("none") != nil)
	fmt.Println(p.GetOpt("none", "Def"))

	// Output:
	// cg.exe
	// cg.exe
	// true
	// true
	// false
	// Def
}

func ExampleCommandFromCtx() {
	var (
		command = gcmd.Command{
			Name: "start",
		}
	)

	ctx := context.WithValue(gctx.New(), gcmd.CtxKeyCommand, &command)
	unAddCtx := context.WithValue(gctx.New(), gcmd.CtxKeyCommand, &gcmd.Command{})
	nonKeyCtx := context.WithValue(gctx.New(), "Testkey", &gcmd.Command{})

	fmt.Println(gcmd.CommandFromCtx(ctx).Name)
	fmt.Println(gcmd.CommandFromCtx(unAddCtx).Name)
	fmt.Println(gcmd.CommandFromCtx(nonKeyCtx) == nil)

	// Output:
	// start
	//
	// true
}

func ExampleCommand_AddCommand() {
	commandRoot := &gcmd.Command{
		Name: "cg",
	}
	commandRoot.AddCommand(&gcmd.Command{
		Name: "start",
	}, &gcmd.Command{})

	commandRoot.Print()

	// Output:
	// USAGE
	//     cg COMMAND [OPTION]
	//
	// COMMAND
	//     start
}

func ExampleCommand_AddCommand_Repeat() {
	commandRoot := &gcmd.Command{
		Name: "cg",
	}
	err := commandRoot.AddCommand(&gcmd.Command{
		Name: "start",
	}, &gcmd.Command{
		Name: "stop",
	}, &gcmd.Command{
		Name: "start",
	})

	fmt.Println(err)

	// Output:
	// command "start" is already added to command "cg"
}

func ExampleCommand_AddObject() {
	var (
		command = gcmd.Command{
			Name: "start",
		}
	)

	command.AddObject(&TestCmdObject{})

	command.Print()

	// Output:
	// USAGE
	//     start COMMAND [OPTION]
	//
	// COMMAND
	//     root    root env command
}

func ExampleCommand_AddObject_Error() {
	var (
		command = gcmd.Command{
			Name: "start",
		}
	)

	err := command.AddObject(&[]string{"Test"})

	fmt.Println(err)

	// Output:
	// input object should be type of struct, but got "*[]string"
}

func ExampleCommand_Print() {
	commandRoot := &gcmd.Command{
		Name: "cg",
	}
	commandRoot.AddCommand(&gcmd.Command{
		Name: "start",
	}, &gcmd.Command{})

	commandRoot.Print()

	// Output:
	// USAGE
	//     cg COMMAND [OPTION]
	//
	// COMMAND
	//     start
}

func ExampleScan() {
	fmt.Println(gcmd.Scan("cg scan"))

	// Output:
	// cg scan
}

func ExampleScanf() {
	fmt.Println(gcmd.Scanf("cg %s", "scanf"))

	// Output:
	// cg scanf
}

func ExampleParserFromCtx() {
	parser, _ := gcmd.Parse(nil)

	ctx := context.WithValue(gctx.New(), gcmd.CtxKeyParser, parser)
	nilCtx := context.WithValue(gctx.New(), "NilCtxKeyParser", parser)

	fmt.Println(gcmd.ParserFromCtx(ctx).GetArgAll())
	fmt.Println(gcmd.ParserFromCtx(nilCtx) == nil)

	// Output:
	// [cg build main.go]
	// true
}

func ExampleParseArgs() {
	p, _ := gcmd.ParseArgs([]string{
		"cg", "--force", "remove", "-fq", "-p=www", "path", "-n", "root",
	}, nil)

	fmt.Println(p.GetArgAll())
	fmt.Println(p.GetOptAll())

	// Output:
	// [cg path]
	// map[force:remove fq: n:root p:www]
}

func ExampleParser_GetArg() {
	p, _ := gcmd.ParseArgs([]string{
		"cg", "--force", "remove", "-fq", "-p=www", "path", "-n", "root",
	}, nil)

	fmt.Println(p.GetArg(-1, "Def").String())
	fmt.Println(p.GetArg(-1) == nil)

	// Output:
	// Def
	// true
}
