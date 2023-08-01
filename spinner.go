package verbose

import (
	"fmt"
	"os"
	"time"

	"github.com/artyom/spinner"
)

func (v *Verb) Spin(quit chan bool) {
	fmt.Println("")
	spnr := spinner.New(os.Stdout, "  Working: ")
	for {
		select {
		case <-quit:
			spnr.Clear()
			return
		default:
			spnr.Spin()
		}
		time.Sleep(250 * time.Millisecond)
	}
}
