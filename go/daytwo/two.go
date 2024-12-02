package daytwo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Run(input io.Reader) string {
	scanner := bufio.NewScanner(input)

	var safeCount int

	dampenerPossibilities := []func([]string, int) []string{
		func(r []string, l int) []string { return slices.Delete(slices.Clone(r), l-1, l) }, // remove current level
		func(r []string, l int) []string { return slices.Delete(slices.Clone(r), l, l+1) }, // remove next level
		func(r []string, _ int) []string { return slices.Delete(slices.Clone(r), 0, 1) },   // remove first level
	}

	for scanner.Scan() {
		line := scanner.Text()
		report := strings.Fields(line)

		isSafe, level := processReport(report)
		for _, dampener := range dampenerPossibilities {
			if isSafe {
				break
			}

			isSafe, _ = processReport(dampener(report, level))
		}

		if isSafe {
			safeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}

	return fmt.Sprintf("safe count with dampener: %d\n", safeCount)
}

func processReport(report []string) (bool, int) {
	firstLevel, err := strconv.Atoi(report[0])
	if err != nil {
		panic(err)
	}

	secondLevel, err := strconv.Atoi(report[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
		panic(err)
	}
	var isIncreasing bool

	if firstLevel-secondLevel < 0 {
		isIncreasing = true
	}
	for i := 1; i < len(report); i++ {
		currentLevel, err := strconv.Atoi(report[i-1])
		if err != nil {
			panic(err)
		}

		nextLevel, err := strconv.Atoi(report[i])
		if err != nil {
			panic(err)
		}

		if currentLevel-nextLevel == 0 {
			return false, i
		}

		if isIncreasing && currentLevel-nextLevel > 0 {
			return false, i
		}

		if !isIncreasing && currentLevel-nextLevel < 0 {
			return false, i
		}

		if absInt(currentLevel-nextLevel) > 3 {
			return false, i
		}
	}

	return true, -1
}

func absInt(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}
