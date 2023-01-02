package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	var stdin *bufio.Reader
	var line []rune
	var c rune
	var err error

	stdin = bufio.NewReader(os.Stdin)

	fmt.Printf("Type something: ")

	for {
		c, _, err = stdin.ReadRune()
		if err == io.EOF || c == '\n' {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading standard input\n")
			os.Exit(1)
		}
		line = append(line, c)
	}

	fmt.Printf("\nYou entered: %s\n", string(line[:len(line)]))
}
