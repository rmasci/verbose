package verbose

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// Just like fmt.Print -- only prints when verbose.V is true. Only prints the date and line number if PrintDate and PrintLine are true
func (v *Verb) Print(a ...any) {
	if v.V {
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

// Just like fmt.Println -- only prints when verbose.V is true,  Only prints the date and line number if PrintDate and PrintLine are true
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

// Just like fmt.Printf, but only prints if verb.V is true  Only prints the date and line number if PrintDate and PrintLine are true
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

// Prints a interface (struct) in indented JSON. Only prints if verb.V is true  Line numbers are not printed.
func (v *Verb) Printj(data interface{}) {
	if v.PrintDate == true {
		fmt.Println(time.Now().Format(v.Dformat))
	}
	if v.V {
		jsonData, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Fprintf(v.Out, "Error marshaling data: %v\n", err)
			return
		}
		fmt.Fprintln(v.Out, string(jsonData))
	}
}
