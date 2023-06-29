package main

import (
	"fmt"
	"github.com/rmasci/verbose"
	"github.com/spf13/pflag"
	"os"
)

func main() {
	var out string
	var x, y int
	var verb verbose.Verb
	var err error
	var help bool
	pflag.BoolVarP(&verb.V, "verbose", "v", false, "Verbose mode.")
	pflag.IntVarP(&x, "x-value", "x", 2, "Set value of x")
	pflag.IntVarP(&y, "y-value", "y", 2, "Set value of y")
	pflag.BoolVarP(&help, "help", "h", false, "Help Message")
	pflag.StringVarP(&out, "out", "o", "stdout", "Set stdout, stderr, or to a file")
	pflag.Parse()
	if help {
		pflag.PrintDefaults()
		fmt.Println("You can test this as follows")
		fmt.Println("test2 -x5 -y5")
		fmt.Printf("Output will be:\n\tx+y=10\n")
		fmt.Println("test2 -x5 -y5 -v -o stderr")
		fmt.Printf("Output will be:\n\tSetting output to stderr\n\tSetting x to 5\n\tSetting y to 2\n\tx+y= 7\n\tThis is fprintln\n\tThis is fprint\n\tThis is printf\n\tThis is println\n\tThis is print\n")
		fmt.Println("test2 -x5 -y5 -v -o stderr 2> /dev/null")
		fmt.Printf("Output will be:\n\tx+y=10\nThis redirects the verbose output to stderr which is directed with 2> to /dev/null.\nYou can try this with other such as -o /tmp/test2.log")
	}
	switch out {
	case "stdout":
		verb.Out = os.Stdout
		verb.Println("Setting output to stdout")
	case "stderr":
		verb.Out = os.Stderr
		verb.Println("Setting output to stderr")
	default:
		verb.Out, err = os.OpenFile(out, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		verb.Println("Setting output to", out)
	}
	verb.Println("Setting x to", x)
	verb.Println("Setting y to", y)
	fmt.Println("x+y=", x+y)
	verb.Fprintf(verb.Out, "This is fprintf %s", "\n")
	verb.Fprintln(verb.Out, "This is fprintln")
	verb.Fprint(verb.Out, "This is fprint", "\n")
	verb.Printf("This is printf %s", "\n")
	verb.Println("This is println")
	verb.Print("This is print", "\n")

}
