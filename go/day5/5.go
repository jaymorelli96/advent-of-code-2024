package day5

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	orderSeparator  = "|"
	updateSeparator = ","
)

func Run(input io.Reader) string {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	updates := make([][]int, 0)
	mustComeBefore := make(map[int][]int)
	mustComeAfter := make(map[int][]int)

	var shouldReadUpdates bool

	for scanner.Scan() {
		line := scanner.Text()

		if !strings.Contains(line, orderSeparator) {
			shouldReadUpdates = true
		}

		// read second part of input
		if shouldReadUpdates {
			updates = append(updates, readUpdates(line))
			continue
		}

		// read ordering rules
		// split input by |, the first part must come before the second part
		ordering := strings.Split(line, orderSeparator)
		x, err := strconv.Atoi(ordering[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(ordering[1])
		if err != nil {
			panic(err)
		}

		mustComeAfter[y] = append(mustComeAfter[y], x)
		mustComeBefore[x] = append(mustComeBefore[x], y)
	}

	for k, v := range mustComeBefore {
		fmt.Printf("%d must come before: %v\n", k, v)
	}
	for k, v := range mustComeAfter {
		fmt.Printf("%d must come after: %v\n", k, v)
	}

	return fmt.Sprintf("part 1 is: %d", p1(mustComeAfter, updates))
}

func p1(mustComeAfter map[int][]int, updates [][]int) int {
	var anwser int

	for _, update := range updates {
		var isViolating bool
		for i, n := range update {
			after := mustComeAfter[n] // this numbers cannot be at the rest of the slice
			if doesViolate(update[i:], after) {
				isViolating = true
				fmt.Printf("update %v violates --- the problem is the %d\n", update, n)
			}
		}

		if !isViolating {
			anwser += update[len(update)/2]
			fmt.Printf("update %v DOES NOT violates\n", update)
		}
	}

	return anwser
}

func doesViolate(sliceToCheck []int, numsToBeExlucded []int) bool {
	for _, check := range sliceToCheck {
		for _, excluded := range numsToBeExlucded {
			if check == excluded {
				return true
			}
		}
	}

	return false
}

func readUpdates(line string) []int {
	splitted := strings.Split(line, updateSeparator)
	updates := make([]int, 0, len(splitted))
	for _, s := range splitted {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		updates = append(updates, n)
	}

	return updates
}
