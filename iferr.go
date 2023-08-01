package verbose

import (
	"os"
	"runtime"
)

// err out will always print if an error is present. to only print when verb is true use Err. Prints to whatever verbose.Out is set to.
func (v *Verb) ErrOut(err error, str string, e ...bool) bool {
	var exit bool
	// We're going to use verbose.Printf
	vOrigState := v.V
	if len(e) > 0 {
		exit = e[0]
	}
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		_, cfile, cline, _ := runtime.Caller(2)
		v.V = true
		v.Printf("error: %v -- %v\n\tfile: %v line: %v\n\tfile: %v line: %v\n", str, err, file, line, cfile, cline)
		if exit {
			os.Exit(1)
		}
		v.V = vOrigState
		return true
	}
	return false
}

// Only prints an error when it's verb.V is set to true
// returns true if err is set. Allows you to specify other actions for err while printing to log or stdout, stderr
// if verb.Err(err,str,false) {
//   return nil,"",blah blah
// }

func (v *Verb) Err(err error, str string, e ...bool) bool {
	var exit bool
	// We're going to use verbose.Printf
	if len(e) > 0 {
		exit = e[0]
	}
	if v.V {
		return v.ErrOut(err, str, exit)
	}
	if err != nil {
		return true
	}
	return false
}
