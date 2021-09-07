package manipulate

import (
	"fmt"
	"os"
	"flag"
	"strings"
)

// pass argument on cmd 
// run main.go a b c   << mean args = a b c 3 args
func ExCmd(){
	argsWithProg := os.Args  // full program and prog
    argsWithoutProg := os.Args[1:]  // show only args

    arg := os.Args[3]   // args index 3

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}


// can pass flag to program
// if not pass, flag will use default value
// if pass incrrect flag, then error
// run main.go -word=aaaa   << mean pass aaaa to word
// run main.go a b c	<< a b c will keep in tail
// use -h , --help for help
func ExCmdFlag(){
	wordPtr := flag.String("word", "foo", "a string")

    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    flag.Parse()

    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())  // tail for another argument but must not use flag
}

func ExCmdSubCmd(){
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
    fooEnable := fooCmd.Bool("enable", false, "enable")
    fooName := fooCmd.String("name", "", "name")

    barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
    barLevel := barCmd.Int("level", 0, "level")

    if len(os.Args) < 2 {
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }

    switch os.Args[1] {

    case "foo":
        fooCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'foo'")
        fmt.Println("  enable:", *fooEnable)
        fmt.Println("  name:", *fooName)
        fmt.Println("  tail:", fooCmd.Args())
    case "bar":
        barCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'bar'")
        fmt.Println("  level:", *barLevel)
        fmt.Println("  tail:", barCmd.Args())
    default:
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }
}

// for windows $env:BAR=3; go run .\cmd\app\main.go
// for linux BAR=3 go run .\cmd\app\main.go
func ExEnv(){
	os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR"))

    fmt.Println()
    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0])
    }
}