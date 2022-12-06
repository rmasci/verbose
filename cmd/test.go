package main

import (
	"fmt"

	"github.com/rmasci/verbose"
	"github.com/spf13/pflag"
)

func main() {
	i := 0
	verb := verbose.New("%A %B %Y, %I:%M:%S %P %Z")
	pflag.BoolVarP(&verb.V, "verbose", "v", false, "Verbose Mode")
	pflag.Parse()
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

	verb = verbose.New("date")
	verb.V = true
	verb.PrintLine = true
	verb.Delimeter = "|"
	verb.Printf("Date keyworkd no line number. %d, %s\n", i, "Printf")
	i++
	verb.Println("date keyword no line number", i, "Println")
	i++
	verb = verbose.New()
	verb.V = true
	verb.PrintLine = true
	verb.Printf("No Date  line number. %d, %s\n", i, "Printf")
	i++
	verb.Println("No Date line number", i, "Println")
	i++
	fmt.Println("Done")
}
