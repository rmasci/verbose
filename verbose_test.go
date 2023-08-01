package verbose

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func Test_verboseNew(t *testing.T) {
	tests := []struct {
		name         string
		dformat      string
		dateExpected string
	}{
		{"nodate", "", "2006-01-02 15:04:05"},
		{"blank", "", "2006-01-02 15:04:05"},
		{"default", "default", "2006-01-02 15:04:05"},
		{"monthDayYear", "%m.%d.%Y %H.%M.%S", "01.02.2006 15.04.05"},
	}
	for _, test := range tests {
		var verb Verb
		if test.name == "blank" {
			verb = New(os.Stdout)
		} else {
			verb = New(os.Stdout, test.dformat)
		}
		if strings.TrimSpace(test.dateExpected) != strings.TrimSpace(verb.Dformat) {
			t.Errorf("%s: expected: \n%s received:\n%s", test.name, test.dateExpected, verb.Dformat)
			continue
		}
		verb.Delimeter = "-"
		if verb.Delimeter != "-" {
			t.Errorf("%s: wrong delimeter. expected ' ' got:'%s'", test.name, verb.Delimeter)
		}
	}
}

func Test_timeFormatStr(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		expected string
	}{
		{"year", "%Y", "2006"},
		{"month", "%m", "01"},
		{"date", "%D", "01/02/06"},
		{"monthName", "%B", "January"},
		{"monthShort", "%b", "Jan"},
		{"day", "%d", "02"},
		{"zzDay", "%j", "002"},
		{"dayName", "%A", "Monday"},
		{"dayShort", "%a", "Mon"},
		{"hour24", "%H", "15"},
		{"hour12", "%I", "03"},
		{"Minute", "%M", "04"},
		{"Second", "%S", "05"},
		{"Nano", "%N", ".000"},
		{"dstamp", "%F", "2006-01-02"},
		{"timestamp", "%T", "15:04:05"},
		{"zone", "%Z", "MST"},
		{"ampm", "%P", "PM"},
	}
	for _, test := range tests {
		x := TimeFormatStr(test.format)
		if x != test.expected {
			t.Errorf("%s: Wanted %s Got %s", test.name, test.expected, x)
		}
	}
}

func Test_verbPrint(t *testing.T) {
	var tn string
	tests := []struct {
		name      string
		action    string
		delim     string
		text      string
		printdate bool
		printline bool
	}{
		{"print", "print", "|", "Test print", true, false},
		{"print2", "print", "", "Test print", true, false},
		{"println", "println", "|", "Test println", true, false},
		{"println", "println", "|", "Test println", true, false},
		{"printf", "printf", "default", "Test printf", true, false},
		{"fprintln", "fprintln", "|", "Test println", true, false},
		{"fprintf", "fprintf", "default", "Test printf", true, false},
		{"DateWithLine", "number", "|", "Test print number", true, true},
		{"nodate", "nodate", "|", "No Date", false, false},
	}
	tf := fmt.Sprintf("%%Y%%m%%d")

	for _, test := range tests {
		tn = time.Now().Format(TimeFormatStr("%Y%m%d"))
		var expected string
		//stdout := os.Stdout
		r, w, _ := os.Pipe()
		//os.Stdout = w
		verb := New(w, tf)
		verb.V = true
		if test.delim == "default" {
			test.delim = " "
		} else {
			verb.Delimeter = test.delim
		}
		verb.PrintLine = test.printline
		verb.PrintDate = test.printdate
		switch test.action {
		case "print":
			verb.Print(test.text)
			expected = fmt.Sprintf("%s%s%s", tn, test.delim, test.text)
		case "println":
			verb.Println(test.text)
			expected = fmt.Sprintf("%s%s%s", tn, test.delim, test.text)
		case "printf":
			verb.Printf("%s", test.text)
			expected = fmt.Sprintf("%s%s%s", tn, test.delim, test.text)
		case "fprintln":
			verb.Fprintln(w, test.text)
			expected = fmt.Sprintf("%s%s%s", tn, test.delim, test.text)
		case "fprintf":
			verb.Fprintf(w, "%s", test.text)
			expected = fmt.Sprintf("%s%s%s", tn, test.delim, test.text)
		case "number":
			verb.Printf("%s", test.text)
			// 116 must equal the line number above
			expected = fmt.Sprintf("%s%sverbose_test.go:118%s%s", tn, test.delim, test.delim, test.text)
		case "nodate":
			verb.Printf("%s", test.text)
			expected = fmt.Sprintf("%s", test.text)
		}
		w.Close()
		//re := regexp.MustCompile(expected)
		bTxt, _ := io.ReadAll(r)
		txt := string(bTxt)
		//t.Logf("txt: %s", txt)
		//testStr := fmt.Sprintf("%s%s%s", tn, test.delim, test.text)
		txt = strings.TrimSpace(txt)
		if txt != expected {
			//if re.Match(bTxt) {
			t.Errorf("%s: Does not contain '%s' Has '%s'", test.name, expected, txt)
		}
		//os.Stdout = stdout
	}
}
