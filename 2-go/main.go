package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var safeCount int

reports:
	for scanner.Scan() {
		line := scanner.Text()
		report := strings.Fields(line)

		firstLevel, err := strconv.Atoi(report[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading input:", err)
			return
		}

		secondLevel, err := strconv.Atoi(report[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading input:", err)
			return
		}

		var isIncreasing bool

		if firstLevel-secondLevel < 0 {
			isIncreasing = true
		}

		for i := 1; i < len(report); i++ {
			currentLevel, err := strconv.Atoi(report[i-1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "error reading input:", err)
				return
			}

			nextLevel, err := strconv.Atoi(report[i])
			if err != nil {
				fmt.Fprintln(os.Stderr, "error reading input:", err)
				return
			}

			if currentLevel-nextLevel == 0 {
				fmt.Printf("line: %s, not safe is neither an increase or decrease\n", line)
				continue reports
			}

			if isIncreasing && currentLevel-nextLevel > 0 {
				fmt.Printf("line: %s, not safe should be increasing\n", line)
				continue reports
			}

			if !isIncreasing && currentLevel-nextLevel < 0 {
				fmt.Printf("line: %s, not safe should be decreasing\n", line)
				continue reports
			}

			if absInt(currentLevel-nextLevel) > 3 {
				fmt.Printf("line: %s, not safe\n", line)
				continue reports
			}

		}

		safeCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}

	fmt.Printf("safe count: %d\n", safeCount)
}

func absInt(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}
