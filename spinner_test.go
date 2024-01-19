package verbose

import (
	"os"
	"testing"
	"time"

	"github.com/briandowns/spinner"
)

func TestSpin(t *testing.T) {
	quit := make(chan bool)
	v := New()

	go func() {
		time.Sleep(1 * time.Second)
		quit <- true
	}()

	// Redirect stdout to a buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	v.Spin(quit)

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read from the buffer and check the output
	out := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		out <- buf.String()
	}()

	expected := "  Working: |  Working: /  Working: -  Working: \\"
	select {
	case got := <-out:
		if got != expected {
			t.Errorf("Spin() = %q, want %q", got, expected)
		}
	case <-time.After(2 * time.Second):
		t.Errorf("Spin() took too long")
	}
}
