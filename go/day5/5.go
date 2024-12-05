package day5

import (
	"bufio"
	"fmt"
	"io"
	"slices"
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

	return fmt.Sprintf("part 2 is: %v", p2(mustComeBefore, mustComeAfter, updates))
}

func p2(mustComeBefore, mustComeAfter map[int][]int, updates [][]int) [][]int {
	var count int

	notValids := getUpdatesNotValid(mustComeAfter, updates)
	sorted := make([][]int, 0)

	for _, update := range notValids {
		fmt.Printf("before sorting not valid %v\n", update)
		res := sortByRestriction(mustComeBefore, update)
		sorted = append(sorted, res)
		count += res[len(res)/2]
		fmt.Printf("after sorting not valid %v\n", sorted)
	}

	fmt.Println("count", count)
	return sorted
}

func sortByRestriction(mustComeBefore map[int][]int, update []int) []int {
	slices.SortFunc(update, func(a, b int) int {
		if slices.Contains(mustComeBefore[a], b) {
			return -1
		} else {
			return 1
		}
	})

	return update
}

func getUpdatesNotValid(mustComeAfter map[int][]int, updates [][]int) [][]int {
	notValid := make([][]int, 0, len(updates))

	for _, update := range updates {
		for i, n := range update {
			after := mustComeAfter[n] // this numbers cannot be at the rest of the slice
			if doesViolate(update[i:], after) {
				fmt.Printf("update %v violates --- the problem is the %d\n", update, n)
				notValid = append(notValid, update)
				break
			}
		}
	}

	return notValid
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
