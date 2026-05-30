package zeropad

import "testing"

func TestPadNumbers(t *testing.T) {
	tests := []struct {
		input    string
		x        int
		expected string
	}{
		// provided examples
		{"James Bond 7",  3, "James Bond 007"},
		{"PI=3.14",        2, "PI=03.14"},
		{"It's 3:13pm",   2, "It's 03:13pm"},
		{"It's 12:13pm",  2, "It's 12:13pm"},
		{"99UR1337",       6, "000099UR001337"},

		// edge cases
		{"",               3, ""},
		{"Hello",          3, "Hello"},
		{"007",            3, "007"},      // already wide enough
		{"12345",          3, "12345"},    // wider than x, no change
		{"0",              4, "0000"},     // zero pads too
		{"42",             0, "42"},       // x=0, no change
		{"3.14",           3, "003.14"},   // integer padded, fraction untouched
		{"1.5 and 2.75",   3, "001.5 and 002.75"},
		{".99",            3, ".99"},      // pure fraction, nothing to pad
		{"-5",             3, "-005"},     // '-' is not part of the number
		{"$3.99",          2, "$03.99"},
		{"a1b2",           3, "a001b002"},
	}

	for _, tt := range tests {
		got := PadNumbers(tt.input, tt.x)
		if got != tt.expected {
			t.Errorf("PadNumbers(%q, %d) = %q; want %q", tt.input, tt.x, got, tt.expected)
		}
	}
}
