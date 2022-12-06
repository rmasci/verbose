package verbose

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type Verb struct {
	V         bool
	Dformat   string
	Delimeter string
	PrintDate bool
	PrintLine bool
}

func New(a ...any) (verb Verb) {
	if len(a) <= 0 {
		verb.Dformat = "2006-01-02 15:04:05 "
		// Print Date not specified.
	} else if a[0] == "date" {
		verb.Dformat = "2006-01-02 15:04:05 "
		verb.PrintDate = true
		return verb
	} else {
		str := fmt.Sprintln(a...)
		verb.PrintDate = true
		verb.Dformat = TimeFormatStr(str)
	}
	if verb.Delimeter == "" {
		verb.Delimeter = " "
	}
	return verb
}

// Just like
func (verb Verb) Println(a ...any) {
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
			fmt.Printf("%s:%d%s", file, line, verb.Delimeter)
		}
		fmt.Println(a...)
	}
}

// TODO: Print Line Numbers -- maybe just switch to log?
// Just like printf, but only prints if verb.V is true -- puts date at beginning of string.
func (verb Verb) Printf(format string, a ...any) {
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
			fmt.Printf("%s:%d%s", file, line, verb.Delimeter)
		}
		fmt.Printf(format, a...)
	}
}

// Converts Unix/Linux date stamp to Go Date Format.
// date +"%A %d %B %Y, %I:%M:%S %P %Z"
// Returns: Thursday December 12 2022, 06:03:08 P EST
// verbose.TimeFormatStr("%A %B %d %Y, %I:%M:%S %P %Z")
// Returns: "Monday January 02 2006, 03:05:05 PM MST". When put into time.Now().Format("Monday January 02 2006, 03:05:05 PM MST")
// Would give you: "Thursday December 12 2022, 06:03:08 PM EST"
func TimeFormatStr(tformat string) (fmtStr string) {
	tFormat := strings.Split(tformat, "")
	for i := 0; i < len(tFormat); i++ {
		if tFormat[i] == "%" {
			i++
			switch tFormat[i] {
			case "Y":
				fmtStr += "2006"
			case "m":
				fmtStr += "01"
			case "D":
				fmtStr += "01/02/06"
			case "B":
				fmtStr += "January"
			case "b":
				fmtStr += "Jan"
			case "d":
				fmtStr += "02"
			case "j":
				fmtStr += "002"
			case "A":
				fmtStr += "Monday"
			case "a":
				fmtStr += "Mon"
			case "H":
				fmtStr += "15"
			case "I":
				fmtStr += "03"
			case "M":
				fmtStr += "04"
			case "S":
				fmtStr += "05"
			case "N":
				fmtStr += ".000"
			case "F":
				fmtStr += "2006-01-02"
			case "T":
				fmtStr += "15:04:05"
			case "Z":
				fmtStr += "MST"
			case "P":
				fmtStr += "PM"
			}
		} else {
			fmtStr += tFormat[i]
		}

	}
	fmtStr = strings.TrimSpace(fmtStr)
	//fmtStr += " "
	return fmtStr
}
