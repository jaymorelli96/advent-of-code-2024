package daythree

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

const uncorruptedPattern = `(mul\([0-9]{1,3},[0-9]{1,3}\))|(do\(\))|(don\'t\(\))`

func Run(r io.Reader) string {
	scanner := bufio.NewScanner(r)

	var line string
	for scanner.Scan() {
		line += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	uncorruptedRegex := regexp.MustCompile(uncorruptedPattern)
	data := uncorruptedRegex.FindAll([]byte(line), -1)

	var total int
	var doNotProcess bool

	for _, d := range data {
		s := string(d)
		fmt.Println(s)

		switch s {
		case "do()":
			doNotProcess = false
			continue
		case "don't()":
			doNotProcess = true
			continue
		}

		if doNotProcess {
			continue
		}

		s = strings.Replace(s, "mul(", "", 1)
		s = strings.Replace(s, ")", "", 1)

		values := strings.Split(s, ",")

		x, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}

		total += x * y
	}

	return fmt.Sprintf("total value: %d\n", total)
}
