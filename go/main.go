package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/jaymorelli96/advent-of-code-2024/go/dayone"
	"github.com/jaymorelli96/advent-of-code-2024/go/daytwo"
)

func main() {
	path := flag.String("file", "", "file path to run use cases, if omitted it will read from stdin")
	day := flag.String("day", "two", "define which day (challenge) to run")
	flag.Parse()

	var r io.Reader
	if *path == "" {
		r = os.Stdin
	} else {
		f, err := os.Open(*path)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		r = f
	}

	adventOfCode := map[string]func(io.Reader) string{
		"one": func(r io.Reader) string { return dayone.Run(r) },
		"two": func(r io.Reader) string { return daytwo.Run(r) },
	}

	fmt.Println(adventOfCode[*day](r))
}
