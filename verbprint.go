// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//Package verbose provides a simple package to allow for printing out messages that help you show the flow of how your program works.
// And you don't have to remove them. use verbose as a global variable in yoru program
// package main
// import(...)
// var verb verbose.Verb
//
// func main() {
//     verb:=verbose.New()
//     flag.BoolVar(&verb.V, "v", false, "Verbose mode"
//
// Then in your code you can put in your code:
// ...code...
// verb.Println("Database Query:",query)
// ... code ...
// verb.Println("Calculating world peace -- this could take some time.", worldPeace)
//
// You can leave this in your code, hide that there is a -v option if you want to. But the result is that when things go wrong
// and inevitably things go wrong, you can see the flow of your program.

package verbose

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Verb struct {
	V         bool
	Dformat   string
	Delimeter string
	PrintDate bool
	PrintLine bool
	Out       io.Writer `default0:os.Stdout`
}

func New(a ...any) (v Verb) {
	if len(a) <= 0 {
		v.Dformat = "2006-01-02 15:04:05 "
		// Print Date not specified.
	} else if a[0] == "date" {
		v.Dformat = "2006-01-02 15:04:05 "
		v.PrintDate = true
		return v
	} else {
		str := fmt.Sprintln(a...)
		v.PrintDate = true
		v.Dformat = TimeFormatStr(str)
	}
	if v.Delimeter == "" {
		v.Delimeter = " "
	}
	v.Out = os.Stdout
	return v
}

// Just like fmt.Print -- only prints when verbose.V is true, and prints out the line number.
func (v *Verb) Print(a ...any) {
	if v.V {
		if v.Out == nil {
			v.Out = os.Stdout
		}
		if v.Delimeter == "" {
			v.Delimeter = " "
		}
		if v.PrintDate == true {
			tn := time.Now().Format(v.Dformat)
			fmt.Fprintf(v.Out, "%s%s", tn, v.Delimeter)
		}
		if v.PrintLine {
			_, file, line, ok := runtime.Caller(1)
			if !ok {
				file = "???"
				line = 0
			} else {
				file = filepath.Base(file)
			}
			fmt.Fprintf(v.Out, "%s:%d%s", file, line, v.Delimeter)
		}
		fmt.Fprint(v.Out, a...)
	}
}

// Just like fmt.Println -- only prints when verbose.V is true, and prints out the line number.
func (v *Verb) Println(a ...any) {
	if v.V {
		if v.Out == nil {
			v.Out = os.Stdout
		}
		if v.Delimeter == "" {
			v.Delimeter = " "
		}
		if v.PrintDate == true {
			tn := time.Now().Format(v.Dformat)
			fmt.Fprintf(v.Out, "%s%s", tn, v.Delimeter)
		}
		if v.PrintLine {
			_, file, line, ok := runtime.Caller(1)
			if !ok {
				file = "???"
				line = 0
			} else {
				file = filepath.Base(file)
			}
			fmt.Fprintf(v.Out, "%s:%d%s", file, line, v.Delimeter)
		}
		fmt.Fprintln(v.Out, a...)
	}
}

// Just like fmt.Printf, but only prints if verb.V is true -- puts date at beginning of string and includes the line number.
func (v *Verb) Printf(format string, a ...any) {
	if v.V {
		if v.Out == nil {
			v.Out = os.Stdout
		}
		if v.Delimeter == "" {
			v.Delimeter = " "
		}
		if v.PrintDate == true {
			tn := time.Now().Format(v.Dformat)
			fmt.Fprintf(v.Out, "%s%s", tn, v.Delimeter)
		}

		if v.PrintLine {
			_, file, line, ok := runtime.Caller(1)
			if !ok {
				file = "???"
				line = 0
			} else {
				file = filepath.Base(file)
			}
			fmt.Fprintf(v.Out, "%s:%d%s", file, line, v.Delimeter)
		}
		fmt.Fprintf(v.Out, format, a...)
	}
}
