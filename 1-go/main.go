package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	l1, l2, err := getLists(os.Stdin)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("total distance: %d\n", computeTotalDistance(l1, l2))
	fmt.Printf("total similiarity score: %d\n", computeSimiliarityScore(l1, l2))
}

func computeSimiliarityScore(l1, l2 []int) int {
	similiarity := make(map[int]int)
	for _, n1 := range l1 {
		var count int
		for _, n2 := range l2 {
			if n1 == n2 {
				count++
			}
		}

		similiarity[n1] += count
	}

	var totalScore int
	for k, v := range similiarity {
		totalScore += k * v
	}

	return totalScore
}

func computeTotalDistance(l1, l2 []int) int {
	sort.Ints(l1)
	sort.Ints(l2)

	var totalDistance int
	for i := range len(l1) {
		diff := l1[i] - l2[i]
		if diff <= 0 {
			diff *= -1
		}

		totalDistance += diff
	}

	return totalDistance
}

func getLists(input io.Reader) ([]int, []int, error) {
	var listOne []int
	var listTwo []int

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)

		n1, err := strconv.Atoi(numbers[0])
		if err != nil {
			return nil, nil, err
		}
		listOne = append(listOne, n1)

		n2, err := strconv.Atoi(numbers[1])
		if err != nil {
			return nil, nil, err
		}

		listTwo = append(listTwo, n2)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return listOne, listTwo, nil
}
