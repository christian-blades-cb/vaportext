package main

import (
	"bufio"
	"golang.org/x/text/width"
	"os"
)

func main() {
	runescanner := bufio.NewScanner(os.Stdin)
	runescanner.Split(bufio.ScanRunes)
	for runescanner.Scan() {
		r := runescanner.Bytes()
		os.Stdout.Write(width.Widen.Bytes(r))
	}
}
