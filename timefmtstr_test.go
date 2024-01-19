package verbose

import (
	"testing"
)

func TestTimeFormatStr(t *testing.T) {
	tests := []struct {
		tformat string
		want    string
	}{
		{
			tformat: "%A %B %d %Y, %I:%M:%S %P %Z",
			want:    "Monday January 02 2006, 03:05:05 PM MST",
		},
		// Add more test cases here
	}

	v := New()
	for _, test := range tests {
		got := v.TimeFormatStr(test.tformat)
		if got != test.want {
			t.Errorf("TimeFormatStr(%q) = %q, want %q", test.tformat, got, test.want)
		}
	}
}
