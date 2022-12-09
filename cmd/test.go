package main

import (
	"fmt"
	"os"

	"github.com/rmasci/verbose"
	"github.com/spf13/pflag"
)

func main() {
	i := 0
	verb := verbose.New("%A %B %Y, %I:%M:%S %P %Z")
	flagset := pflag.FlagSet
	flagset.BoolVarP(&verb.V, "verbose", "v", false, "Verbose Mode")
	flagset.MarkHidden(verb)
	flagset.Parse(os.Args[1:])
	i++
	verb.Printf("Date Noline number. %d, %s\n", i, "Printf")
	i++
	verb.Println("Date No line number", i, "Println")
	verb = verbose.New()
	verb.PrintLine = true
	i++
	verb.Printf("line number. %d, %s\n", i, "Printf")
	i++
	verb.Println("line number", i, "Println")
	i++
	verb.Print("line number. ", i, " Print\n")
	i++

	verb = verbose.New("date")
	verb.V = true
	verb.PrintLine = true
	verb.Delimeter = "|"
	verb.Printf("Date keyworkd no line number. %d, %s\n", i, "Printf")
	i++
	verb.Println("date keyword no line number", i, "Println")
	i++
	verb.Print("line number. ", i, " Print\n")
	i++

	verb = verbose.New()
	verb.V = true
	verb.PrintLine = true
	verb.Printf("No Date  line number. %d, %s\n", i, "Printf")
	i++
	verb.Println("No Date line number", i, "Println")
	i++
	verb.Print("line number. ", i, " Print\n")
	i++
	fmt.Println("Done")
	var z verbose.Verb
	if z.Out == nil {
		fmt.Println("It's Nil!")
		z.V = true
		z.Dformat = verbose.TimeFormatStr("%Y-%m-%d %I:%M:%S %P %Z")
		z.PrintDate = true
		z.PrintLine = true
		z.Delimeter = "-"
		z.Printf("Verb printf %T\n", z.Out)
		out, _ := os.OpenFile("zout.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
		z.Out = out
		z.Printf("THis is to a file %T\n", z.Out)
		z.Println("THis is to a file using println", i)
		z.Print("This is to a file using print ", i, "\n")
	}
}
