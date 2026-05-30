package zeropad

import (
	"fmt"
	"strings"
	"unicode"
)

// PadNumbers returns a copy of s where every whole number is left-padded
// with zeros to a minimum width of x digits.
// Numbers that follow a '.' are treated as decimal fractions and skipped.
//
// Time:  O(n) — single pass through the string
// Space: O(n) — output is at most input length + padding added
func PadNumbers(s string, x int) string {
	var out strings.Builder
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {

		// Not a digit — write it and move on.
		if !unicode.IsDigit(runes[i]) {
			out.WriteRune(runes[i])
			continue
		}

		// Collect all consecutive digits into a single number.
		start := i
		for i < len(runes) && unicode.IsDigit(runes[i]) {
			i++
		}
		number := string(runes[start:i])
		i-- // step back so the outer loop's i++ lands on the next non-digit

		// If preceded by '.', it's a decimal fraction — write as-is.
		if start > 0 && runes[start-1] == '.' {
			out.WriteString(number)
			continue
		}

		// Otherwise, pad to width x and write.
		out.WriteString(fmt.Sprintf("%0*s", x, number))
	}

	return out.String()
}
