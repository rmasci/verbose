package verbose

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestVerb_Printf(t *testing.T) {
	// Create a new Verb instance
	v := New()

	// Set up test cases
	tests := []struct {
		format   string
		a        []interface{}
		expected string
	}{
		{
			format:   "Hello, %s!",
			a:        []interface{}{"World"},
			expected: "Hello, World!",
		},
		{
			format:   "The answer is %d",
			a:        []interface{}{42},
			expected: "The answer is 42",
		},
		// Add more test cases here
	}

	// Set up a buffer to capture the output
	var buf bytes.Buffer
	v.Out = &buf

	// Run the tests
	for _, test := range tests {
		// Reset the buffer
		buf.Reset()

		// Call the Printf method
		v.Printf(test.format, test.a...)

		// Check if the output matches the expected value
		if got := buf.String(); got != test.expected {
			t.Errorf("Printf(%q, %v) = %q, want %q", test.format, test.a, got, test.expected)
		}
	}
}
func TestNew(t *testing.T) {
	// Set up test cases
	tests := []struct {
		name     string
		w        io.Writer
		a        []interface{}
		expected Verb
	}{
		{
			name:     "No arguments",
			w:        os.Stdout,
			a:        []interface{}{},
			expected: Verb{Dformat: "2006-01-02 15:04:05 ", Delimeter: " ", Out: os.Stdout},
		},
		{
			name:     "Default argument",
			w:        os.Stdout,
			a:        []interface{}{"default"},
			expected: Verb{Dformat: "2006-01-02 15:04:05 ", Delimeter: " ", Out: os.Stdout, PrintDate: true},
		},
		{
			name:     "Custom argument",
			w:        os.Stdout,
			a:        []interface{}{"Custom format"},
			expected: Verb{Dformat: "Custom format", Delimeter: " ", Out: os.Stdout, PrintDate: true},
		},
	}

	// Run the tests
	for _, test := range tests {
		// Call the New function
		got := New(test.w, test.a...)

		// Check if the output matches the expected value
		if got != test.expected {
			t.Errorf("New(%v, %v) = %v, want %v", test.w, test.a, got, test.expected)
		}
	}
}

func TestVerb_Print(t *testing.T) {
	// Create a new Verb instance
	v := New()

	// Set up test cases
	tests := []struct {
		V              bool
		Delimeter      string
		PrintDate      bool
		PrintLine      bool
		Dformat        string
		a              []interface{}
		expectedOutput string
	}{
		{
			V:              true,
			Delimeter:      " ",
			PrintDate:      true,
			PrintLine:      true,
			Dformat:        "2006-01-02 15:04:05",
			a:              []interface{}{"Hello", "World"},
			expectedOutput: "2006-01-02 15:04:05 verbprint_test.go:25 Hello World",
		},
		{
			V:              false,
			Delimeter:      " ",
			PrintDate:      true,
			PrintLine:      true,
			Dformat:        "2006-01-02 15:04:05",
			a:              []interface{}{"Hello", "World"},
			expectedOutput: "",
		},
		// Add more test cases here
	}

	// Set up a buffer to capture the output
	var buf bytes.Buffer
	v.Out = &buf

	// Run the tests
	for _, test := range tests {
		// Reset the buffer
		buf.Reset()

		// Set the Verb instance properties
		v.V = test.V
		v.Delimeter = test.Delimeter
		v.PrintDate = test.PrintDate
		v.PrintLine = test.PrintLine
		v.Dformat = test.Dformat

		// Call the Print method
		v.Print(test.a...)

		// Check if the output matches the expected value
		if got := buf.String(); got != test.expectedOutput {
			t.Errorf("Print(%v) = %q, want %q", test.a, got, test.expectedOutput)
		}
	}
}
func TestVerb_Printj(t *testing.T) {
	// Create a new Verb instance
	v := New()

	// Set up test cases
	tests := []struct {
		V        bool
		data     interface{}
		expected string
	}{
		{
			V:        true,
			data:     struct{ Name string }{"John"},
			expected: "{\n  \"Name\": \"John\"\n}\n",
		},
		// Add more test cases here
	}

	// Set up a buffer to capture the output
	var buf bytes.Buffer
	v.Out = &buf

	// Run the tests
	for _, test := range tests {
		// Reset the buffer
		buf.Reset()

		// Set the Verb instance properties
		v.V = test.V

		// Call the Printj method
		v.Printj(test.data)

		// Check if the output matches the expected value
		if got := buf.String(); got != test.expected {
			t.Errorf("Printj(%v) = %q, want %q", test.data, got, test.expected)
		}
	}
}
