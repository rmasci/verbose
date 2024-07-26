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
)

type Verb struct {
	// V when set to true enables the verbose printing.
	V bool
	// set the date format using standard Go Formatting 2006/01/02 15:04:05
	Dformat string
	// Set the delimiter between date, line number and print string.
	Delimeter string
	// If set to false, date will not be printed
	PrintDate bool
	// If set to false, line number will not be printed
	PrintLine bool
	// Set where to write the print statements. By default it's stderr, but you can change it to stdout, or to a file.
	Out io.Writer `default0:os.Stderr`
	// Quit is a verbose channel
	Quit chan bool
}

// Returns a type Verb and sets some defaults.
// If nothing passed verbose.New(), no date is used
// if verbose.New("default") use a default date string.
// Date string can be customized by either setting Dformat using go date format string or by passing a linux date compatible string to verbose.New.
func New(w io.Writer, a ...any) (v Verb) {
	if len(a) <= 0 {
		v.Dformat = "2006-01-02 15:04:05 "
	} else if a[0] == "default" || a[0] == "" {
		v.Dformat = "2006-01-02 15:04:05 "
		v.PrintDate = true
		return v
	} else {
		str := fmt.Sprintln(a...)
		v.PrintDate = true
		v.Dformat = TimeFormatStr(str)
	}

	v.Delimeter = " "
	v.Out = w
	v.Quit = make(chan bool)
	return v
}
