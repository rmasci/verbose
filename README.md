# Verbose

Allows you to easily add in 'debugging' statements to your code. You don't have to remove them once you finalize your code.

```cgo
package main

import (
    "fmt"
    "flag"
    "github.com/rmasci/verbose"
 )
    // The date string is formatted the same as using the Linux date Command
    verb := verbose.New("%A %B %Y, %I:%M:%S %P %Z")
	flagset:=pflag.FlagSet
	flagset.BoolVarP(&verb.V, "verbose", "v", false, "Verbose Mode")
	// This hides the verb option if the user does --help. or flagset.PrintDefaults()
	flagset.MarkHidden(verb)
	flagset.Parse(os.Args[1:])
	verb.Printf("%s started\n",os.Args[0])
	verb.Printf("Remaining Args passed: %v\n",flagset.Args())
	... lots of code ...
	verb.Println("Query Database")
	... lots of code ...
	verb.Println("Rest Call to www.somewhere.com")
```
When the user runs a program like this if the -v or --verbose is not passed as an option, the verb.Print statements will not print.

The default output is to os.Stdout - but this can be changed. If in the above example we wanted to write the output to a file. or to stderr we can do that two ways.
1. set verb.Out:=*os.Writer
```cgo
// set message to file
w:=os.Create("verbose.txt")
verb.Out=w
verb.Println("This line goes to a file")
```
2. Use verb.Fprintf(os.Writer, "Your message goes here.")
```cgo
// Set message to stderr:
verb.Fprintf(os.Stderr, "My Message Goes Here")
```


