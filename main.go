package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	os.Exit(run())
}

func wc(io io.Reader) (int, int, int) {
	var lines, words, bufsize int
	s := bufio.NewScanner(io)
	for s.Scan() {
		// TODO: FIXME
		words += len(strings.Split(strings.TrimSpace(s.Text()), " "))
		bufsize += len(s.Text())
		lines++
	}
	return lines, words, bufsize + lines
}

func run() int {
	if len(os.Args) == 1 {
		lines, words, bufsize := wc(os.Stdin)
		fmt.Printf("%8d%8d%8d\n", lines, words, bufsize)
		return 0
	}

	for _, f := range os.Args[1:] {
		fp, err := os.Open(f)
		if err != nil {
			fmt.Printf("wc: %s: open: No such file or directory\n", f)
			continue
		}
		defer fp.Close()
		lines, words, bufsize := wc(fp)
		fmt.Printf("%8d%8d%8d %s\n", lines, words, bufsize, f)
	}
	return 0
}
