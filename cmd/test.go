package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rmasci/verbose"
	"github.com/spf13/pflag"
)

func main() {
	i := 0
	var testme string
	var help bool
	// Notice the formatting is the same as if you used the Linux date command
	verb := verbose.New("%A %B %Y, %I:%M:%S %P %Z")
	var flagset pflag.FlagSet
	flagset.BoolVarP(&verb.V, "verbose", "v", false, "Verbose Mode")
	flagset.BoolVarP(&help, "help", "h", false, "Help")
	flagset.StringVarP(&testme, "string", "s", "", "String to print out.")
	flagset.MarkHidden("verbose")
	flagset.Parse(os.Args[1:])
	if help {
		fmt.Printf("%s Usage:\n", filepath.Base(os.Args[0]))
		flagset.PrintDefaults()
		os.Exit(0)
	}
	// verb.V gets reset to 'false' every time a verbose.New() is called. Set the first verb.V to variable v.
	v := verb.V
	i++
	verb.Printf("Date with no line number. %d, Printf %s\n", i, testme)
	i++
	verb.Println("Date with no line number", i, "Println", testme)
	i++
	verb.Print("Date with no line number. ", i, " Print ", testme, "\n")

	// Use a default date. If nothing is specified, it won't print a date.
	verb = verbose.New("default")
	verb.V = v
	verb.PrintLine = true
	i++
	verb.Printf("Default Date with line number. %d, Printf %s\n", i, testme)
	i++
	verb.Println("Default Date with line number", i, "Println", testme)
	i++
	verb.Print("Default date with line number. ", i, " Print ", testme, "\n")
	i++
	// fprint
	verb.Fprintf(os.Stderr, "Default Date with line number. %d, Fprintf %s\n", i, testme)
	i++
	verb.Fprintln(os.Stderr, "Default Date with line number", i, "Fprintln", testme)
	i++
	verb.Fprint(os.Stderr, "Default date with line number. ", i, " Fprint ", testme, "\n")
	i++
	// no date specified, no date is printed
	verb = verbose.New()
	verb.V = v
	verb.PrintLine = true
	verb.Delimeter = "|"
	verb.Printf("No date, line number. %d, Printf %s\n", i, testme)
	i++
	verb.Println("No date, line number", i, "Println", testme)
	i++
	verb.Print("No date, line number. ", i, " Print ", testme, "\n")
	i++
	// write to a file.
	z := verbose.New()
	z.V = v
	z.PrintDate = true
	z.PrintLine = true
	// Change delimiter from space to " - "
	// you have to add spaces to delimeter
	z.Delimeter = "|"
	out, _ := os.OpenFile("zout.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	z.Out = out
	z.Printf("This is to a file %T\n", z.Out)
	z.Println("This is to a file using println", i)
	z.Print("This is to a file using print ", i, "\n")

	fmt.Println("Done. If this is the only line, try", filepath.Base(os.Args[0]), "-v -s \"Some Text\"")
}
