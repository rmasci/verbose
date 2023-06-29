// Package verbose provides a simple package to allow for printing out messages that help you show the flow of how your program works.
// And you don't have to remove them. Just implement it as a 'flag':
// verb:=verbose.New()
// flag.BoolVar(&verb.V, "v", false, "Verbose mode"
//
// Then in your code you can put in your code:
// func main() {
// ...code...
// verb.Println("Database Query:",query)
// ... code ...
// verb.Println("Calculating world peace -- this could take some time.", worldPeace)
// You can leave this in your code, hide that there is a -v option if you want to. But the result is that when things go wrong
// and inevitably things go wrong, you can see the flow of your program.
// For information about UTF-8 strings in Go, see https://blog.golang.org/strings.
package verbose

import (
	"fmt"
	"io"
	"path/filepath"
	"runtime"
	"time"
)

// Just like fmt.Fprint -- only prints when verbose.V is true.  Only prints the date and line number if PrintDate and PrintLine are true
func (verb *Verb) Fprint(w io.Writer, a ...any) {
	if verb.V {
		if verb.Delimeter == "" {
			verb.Delimeter = " "
		}
		if verb.PrintDate == true {
			tn := time.Now().Format(verb.Dformat)
			fmt.Printf("%s%s", tn, verb.Delimeter)
		}
		if verb.PrintLine {
			_, file, line, ok := runtime.Caller(1)
			if !ok {
				file = "???"
				line = 0
			} else {
				file = filepath.Base(file)
			}
			fmt.Fprintf(w, "%s:%d%s", file, line, verb.Delimeter)
		}
		fmt.Fprint(w, a...)
	}
}

func MyTest() string {
	return "Successfull"
}

// Just like fmt.Fprintln -- only prints when verbose.V is true. Only prints the date and line number if PrintDate and PrintLine are true
func (verb *Verb) Fprintln(w io.Writer, a ...any) {
	if verb.V {
		if verb.Delimeter == "" {
			verb.Delimeter = " "
		}
		if verb.PrintDate == true {
			tn := time.Now().Format(verb.Dformat)
			fmt.Printf("%s%s", tn, verb.Delimeter)
		}
		if verb.PrintLine {
			_, file, line, ok := runtime.Caller(1)
			if !ok {
				file = "???"
				line = 0
			} else {
				file = filepath.Base(file)
			}
			fmt.Fprintf(w, "%s:%d%s", file, line, verb.Delimeter)
		}
		fmt.Fprintln(w, a...)
	}
}

// Just like fmt.Fprintf, but only prints if verb.V is true. Only prints the date and line number if PrintDate and PrintLine are true
func (verb *Verb) Fprintf(w io.Writer, format string, a ...any) {
	if verb.V {
		if verb.Delimeter == "" {
			verb.Delimeter = " "
		}
		if verb.PrintDate == true {
			tn := time.Now().Format(verb.Dformat)
			fmt.Printf("%s%s", tn, verb.Delimeter)
		}

		if verb.PrintLine {
			_, file, line, ok := runtime.Caller(1)
			if !ok {
				file = "???"
				line = 0
			} else {
				file = filepath.Base(file)
			}
			fmt.Fprintf(w, "%s:%d%s", file, line, verb.Delimeter)
		}
		fmt.Fprintf(w, format, a...)
	}
}
