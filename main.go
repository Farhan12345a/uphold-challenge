package main

import (
	"fmt"
	"os"
	"strconv"
	"zeropad"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: zeropad <string> <width>")
		os.Exit(1)
	}

	x, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: width must be an integer, got %q\n", os.Args[2])
		os.Exit(1)
	}

	fmt.Println(zeropad.PadNumbers(os.Args[1], x))
}
