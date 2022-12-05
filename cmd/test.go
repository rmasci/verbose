package main

import (
	"github.com/rmasci/verbose"
	"github.com/spf13/pflag"
)

func main() {
	i := 0
	verb := verbose.New("%A %B %Y, %I:%M:%S %P %Z")
	pflag.BoolVarP(&verb.V, "verbose", "v", false, "Verbose Mode")
	pflag.Parse()
	i++
	verb.Printf("Printf %d, %s\n", i, "Hello There")
	i++
	verb.Println("Println", i, "Hello There")
	verb.Dformat = verbose.TimeFormatStr("%F %T")
	i++
	verb.Printf("Printf %d, %s\n", i, "Hello There")
	i++
	verb.Println("Println", i, "Hello There")
	verb.Delimeter = "|"
	i++
	verb.Printf("Printf %d, %s\n", i, "Hello There")
	i++
	verb.Println("Println", i, "Hello There")
	verb = verbose.New("no")
	verb.V = true
	i++
	verb.Printf("Printf %d, %s\n", i, "Hello There")
	i++
	verb.Println("Println", i, "Hello There")

}
