package zeropad

import (
	"fmt"
	"strings"
	"unicode"
)


func PadNumbers(s string, x int) string {
	var out strings.Builder
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {

		if !unicode.IsDigit(runes[i]) {
			out.WriteRune(runes[i])
			continue
		}

		start := i
		for i < len(runes) && unicode.IsDigit(runes[i]) {
			i++
		}
		number := string(runes[start:i])
		i-- 


		if start > 0 && runes[start-1] == '.' {
			out.WriteString(number)
			continue
		}

		out.WriteString(fmt.Sprintf("%0*s", x, number))
	}

	return out.String()
}
