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
	pflag.BoolVarP(&verb.V, "verbose", "v", false, "Verbose Mode")
	pflag.Parse()
	i++
	verb.Fprintf(os.Stdout, "Date Noline number. %d, %s\n", i, "Printf")
	i++
	verb.Fprintln(os.Stdout, "Date No line number", i, "Println")
	verb = verbose.New()
	verb.PrintLine = true
	i++
	verb.Fprintf(os.Stdout, "line number. %d, %s\n", i, "Printf")
	i++
	verb.Fprintln(os.Stdout, "line number", i, "Println")
	i++
	verb.Fprint(os.Stderr, "line number. ", i, " Print\n")
	i++

	verb = verbose.New("date")
	verb.V = true
	verb.PrintLine = true
	verb.Delimeter = "|"
	verb.Fprintf(os.Stdout, "Date keyworkd no line number. %d, %s\n", i, "Printf")
	i++
	verb.Fprintln(os.Stdout, "date keyword no line number", i, "Println")
	i++
	verb.Fprint(os.Stderr, "line number. ", i, " Print\n")
	i++

	verb = verbose.New()
	verb.V = true
	verb.PrintLine = true
	verb.Fprintf(os.Stdout, "No Date  line number. %d, %s\n", i, "Printf")
	i++
	verb.Fprintln(os.Stdout, "No Date line number", i, "Println")
	i++
	verb.Fprint(os.Stderr, "line number. ", i, " Print\n")
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
