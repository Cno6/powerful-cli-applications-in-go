package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, isCountLines bool, isCountBytes bool) int {
	scanner := bufio.NewScanner(r)

	switch {
	case isCountBytes:
		scanner.Split(bufio.ScanBytes)
	case isCountLines:
		scanner.Split(bufio.ScanLines)
	default:
		scanner.Split(bufio.ScanWords)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *bytes))
}
