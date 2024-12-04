package dayfour

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const word = "XMAS"

func Run(input io.Reader) string {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	matrix := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, strings.Split(line, ""))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}

	var totalCount int

	for i := range matrix {
		for j := range matrix[i] {
			// potential word matching
			if matrix[i][j] == string(word[0]) {
				totalCount += matrixRightDown(matrix, i, i+len(word), j, j+len(word))
				totalCount += matrixRightUp(matrix, i, i-len(word)+1, j, j+len(word))
				totalCount += matrixLeftDown(matrix, i, i+len(word), j, j-len(word)+1)
				totalCount += matrixLeftUp(matrix, i, i-len(word)+1, j, j-len(word)+1)
			}
		}
	}

	return fmt.Sprintf("total apperance of the word %s is = %d", word, totalCount)
}

func matrixRightDown(m [][]string, rowStart, rowFinish, colStart, colFinish int) int {
	var horizontal strings.Builder
	var vertical strings.Builder
	var diagonal strings.Builder

	var di int // counter for diagonal for rows
	for i := rowStart; i < len(m) && i < rowFinish; i++ {
		var dj int // counter for diagonal for columns
		for j := colStart; j < len(m[i]) && j < colFinish; j++ {
			if i == rowStart {
				fmt.Fprintf(&horizontal, "%s", m[i][j])
			}

			if j == colStart {
				fmt.Fprintf(&vertical, "%s", m[i][j])
			}

			if di == dj {
				fmt.Fprintf(&diagonal, "%s", m[i][j])
			}

			dj++
		}

		di++
	}

	return countWordsInAllDirections(word, horizontal.String(), vertical.String(), diagonal.String())
}

func matrixRightUp(m [][]string, rowStart, rowFinish, colStart, colFinish int) int {
	var horizontal strings.Builder
	var vertical strings.Builder
	var diagonal strings.Builder

	var di int // counter for diagonal for rows
	for i := rowStart; i >= 0 && i >= rowFinish; i-- {
		var dj int // counter for diagonal for columns
		for j := colStart; j < len(m[i]) && j < colFinish; j++ {
			if i == rowStart {
				fmt.Fprintf(&horizontal, "%s", m[i][j])
			}

			if j == colStart {
				fmt.Fprintf(&vertical, "%s", m[i][j])
			}

			if di == dj {
				fmt.Fprintf(&diagonal, "%s", m[i][j])
			}

			dj++
		}

		di++
	}

	return countWordsInAllDirections(word, "", vertical.String(), diagonal.String())
}

func matrixLeftDown(m [][]string, rowStart, rowFinish, colStart, colFinish int) int {
	var horizontal strings.Builder
	var vertical strings.Builder
	var diagonal strings.Builder

	var di int // counter for diagonal for rows
	for i := rowStart; i < len(m) && i < rowFinish; i++ {
		var dj int // counter for diagonal for columns
		for j := colStart; j >= 0 && j >= colFinish; j-- {
			if i == rowStart {
				fmt.Fprintf(&horizontal, "%s", m[i][j])
			}

			if j == colStart {
				fmt.Fprintf(&vertical, "%s", m[i][j])
			}

			if di == dj {
				fmt.Fprintf(&diagonal, "%s", m[i][j])
			}

			dj++
		}

		di++
	}

	return countWordsInAllDirections(word, horizontal.String(), "", diagonal.String())
}

func matrixLeftUp(m [][]string, rowStart, rowFinish, colStart, colFinish int) int {
	var horizontal strings.Builder
	var vertical strings.Builder
	var diagonal strings.Builder

	var di int // counter for diagonal for rows
	for i := rowStart; i >= 0 && i >= rowFinish; i-- {
		var dj int // counter for diagonal for columns
		for j := colStart; j >= 0 && j >= colFinish; j-- {
			if i == rowStart {
				fmt.Fprintf(&horizontal, "%s", m[i][j])
			}

			if j == colStart {
				fmt.Fprintf(&vertical, "%s", m[i][j])
			}

			if di == dj {
				fmt.Fprintf(&diagonal, "%s", m[i][j])
			}

			dj++
		}

		di++
	}

	return countWordsInAllDirections(word, "", "", diagonal.String())
}

func countWordsInAllDirections(word, horizontal, vertical, diagonal string) int {
	var count int

	if strings.Contains(horizontal, word) {
		count++
	}

	if strings.Contains(vertical, word) {
		count++
	}

	if strings.Contains(diagonal, word) {
		count++
	}

	if count > 0 {
		fmt.Printf("horizontal: %s\nvertical: %s\ndiagonal: %s\ncount: %d\n", horizontal, vertical, diagonal, count)
	}
	return count
}
