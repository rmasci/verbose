// spinner implements simple "progress spinner" for terminal output.
//
// Spinner expects to have exclusive access to underlying *os.File, that
// nothing else is writing there while spinner is in use, otherwise output
// would be broken.
//
// If provided *os.File (usually os.Stdout or os.Sterr) is not attached to a
// terminal, spinner outputs nothing, that makes it safe to redirect program
// output to files, pipes, etc.
//
// Spinner can either be used manually, by first creating it with New function,
// then periodically calling Spin() method on it to refresh output and finally
// finishing with Clear() method call to clean output; or package-level Spin
// shortcut function can be used to launch background goroutine that handles
// output refresh.
package verbose

import (
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/term"
)

type Spinner struct {
	F *os.File
	// Type. See Types.md. This allows you to set the look of the spinner.
	Type int
	// Speed.  1000ms / <speed> will be how long the spinner will display each segment. Default is 10.
	Speed int
	// Quit is the channel used when you're done with the spinner.  spinner.Quit <- true
	Quit  chan bool
	text  string
	n     int
	Chars []string
}

// make it simple to create a spinner.
//
//	go verb.Spin()
//	time.Sleep(5 * time.Second)
//	verb.Quit <- true
//	fmt.Println("done")
func (v *Verb) Spin(I ...int) {
	if len(I) == 0 {
		I = append(I, 0)
	}
	if I[0] < 1 {
		I[0] = 0
	}
	spnr := NewSpinner("Working:", "stderr", I[0])
	spnr.Speed = 7
	go spnr.Start()
	quit := <-v.Quit
	fmt.Println("Done")
	spnr.Quit <- quit
}

func (s *Spinner) Start() {
	HideCursor()
	for {
		select {
		case <-s.Quit:
			s.Clear()
			return
		default:
			s.Spin()
		}
		time.Sleep(time.Duration(1000/s.Speed) * time.Millisecond)
	}
}

// Spinner implements terminal spinner attached to *os.File which usually
// either stdout or stderr. Both zero and nil values are valid and are no-op.
// If spinner created on an *os.File that is not attached to the terminal,
// spinner's methods do nothing.
//
// Its methods are NOT thread safe, and it expects to have exclusive access to
// underlying *os.File — that nothing is writing to it while Spinner's methods
// are in use.

// // Spin redraws output if underlying *os.File is attached to a terminal.
func (s *Spinner) Spin() {
	s.n = (s.n + 1) % len(s.Chars)
	//t := fmt.Sprintf("%s %s ", s.text, s.Chars[s.n])
	fmt.Fprintf(s.F, "\r%s %s ", s.text, s.Chars[s.n])
	//s.F.WriteString(t)
}

// Clear redraws output with spaces, clearing previous output if underlying
// *os.File is attached to a terminal.
func (s *Spinner) Clear() {
	if s == nil || s.F == nil {
		return
	}
	ShowCursor()
	// this line erases everything in the line.
	// \r is a carriage return, which moves the cursor to the beginning of the line.
	//strings.Repeat(" ", len(s.text)+3) creates a string of spaces with the same length as s.text plus 3 additional
	// spaces. This effectively overwrites the current spinner text with spaces, clearing it from the terminal.
	//strings.Repeat("\b", len(s.text)+3) creates a string of backspace characters with the same length as s.text plus 3
	// additional spaces. This moves the cursor back to the beginning of the line after clearing the text.
	// \r is a carriage return, which moves the cursor to the beginning of the line.
	fmt.Fprintf(s.F, "\r%s%s\r", strings.Repeat(" ", len(s.text)+3), strings.Repeat("\b", len(s.text)+3))
}

// NewSpinner returns new Spinner attached to Out specify 'stdout' for os.Stdout, or 'stderr' normally os.Stderr. default is stderr.
// os.Stderr. If Out is attached to a terminal, retrurned spinner would output
// text followed by space and "spinning" character on each Spin call.
// NOTE: Keep the 'text' string short. This doesn't work well if 'text' wrapps in the user's terminal.
// NOTE:, most putty users can't see the graphical characters. 0, 9, 14, 19 seem to work best.
// *os.File provided must not be nil. t is the 'type' of spinner.
func NewSpinner(text string, out string, t int) *Spinner {
	var f *os.File
	if out == "stdout" {
		f = os.Stdout
	} else {
		f = os.Stderr
	}
	if !term.IsTerminal(int(f.Fd())) {
		return &Spinner{}
	}
	ch := make(chan bool)
	return &Spinner{F: f, text: text, Type: 1, Speed: 10, Quit: ch, Chars: getType(t)}
}

func getType(t int) []string {
	Type := []string{
		`|/-|\-`,
		`←↖↑↗→↘↓↙`,
		`▁▃▄▅▆▇█▇▆▅▄▃▁`,
		`▖▘▝▗`,
		`┤┘┴└├┌┬┐`,
		`◢◣◤◥`,
		`◰◳◲◱`,
		`◴◷◶◵`,
		`◐◓◑◒`,
		`.oO@*`,
		`◡◡⊙⊙◠◠`,
		`⣾⣽⣻⢿⡿⣟⣯⣷`,
		`⠁⠂⠄⡀⢀⠠⠐⠈`,
		`⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏`,
		`abcdefghijklmnopqrstuvwxyz`,
		`▉▊▋▌▍▎▏▎▍▌▋▊▉`,
		`←↑→↓`,
		//`ｦｧｨｩｪｫｬｭｮｯｱｲｳｴｵｶｷｸｹｺｻｼｽｾｿﾀﾁﾂﾃﾄﾅﾆﾇﾈﾉﾊﾋﾌﾍﾎﾏﾐﾑﾒﾓﾔﾕﾖﾗﾘﾙﾚﾛﾜﾝ`,
		`▁▂▃▄▅▆▇█▉▊▋▌▍▎▏▏▎▍▌▋▊▉█▇▆▅▄▃▂▁`,
		`.oO°Oo.`,
		`⬒⬔⬓⬕`,
	}
	if t < 0 || t >= len(Type) {
		t = 0
	}
	return strings.Split(Type[t], "")
}

// HideCursor hides the cursor using ANSI escape codes.
func HideCursor() {
	fmt.Print("\033[?25l")
}

// ShowCursor shows the cursor using ANSI escape codes.
func ShowCursor() {
	fmt.Print("\033[?25h")
}
