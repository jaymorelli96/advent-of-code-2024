package dayfour

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	word        = "MAS"
	focusLetter = "A"
)

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
			// if i != 2 || j != 6 {
			// 	continue
			// }

			if matrix[i][j] == focusLetter {
				var forwardDiag strings.Builder
				var backwardDiag strings.Builder

				_, _, diagBottomRight := matrixRightDown(matrix, i, i+len(word), j, j+len(word))
				_, _, diagTopRight := matrixRightUp(matrix, i, i-len(word)+1, j, j+len(word))
				_, _, diagBottomLeft := matrixLeftDown(matrix, i, i+len(word), j, j-len(word)+1)
				_, _, diagTopLeft := matrixLeftUp(matrix, i, i-len(word)+1, j, j-len(word)+1)
				fmt.Printf(
					"top right: %s\nbottom right: %s\nbottom left: %s\ntop left: %s\n",
					diagTopRight[:min(len(word)-1, len(diagTopRight))],
					diagBottomRight[:min(len(word)-1, len(diagBottomRight))],
					diagBottomLeft[:min(len(word)-1, len(diagBottomLeft))],
					diagTopLeft[:min(len(word)-1, len(diagTopLeft))],
				)

				diagTopRight = diagTopRight[:min(len(word)-1, len(diagTopRight))]
				diagBottomRight = diagBottomRight[:min(len(word)-1, len(diagBottomRight))]
				diagBottomLeft = diagBottomLeft[:min(len(word)-1, len(diagBottomLeft))]
				diagTopLeft = diagTopLeft[:min(len(word)-1, len(diagTopLeft))]

				forwardDiag.WriteString(diagTopLeft)
				forwardDiag.WriteString(diagBottomRight)

				// /
				backwardDiag.WriteString(diagTopRight)
				backwardDiag.WriteString(diagBottomLeft)

				fWord := forwardDiag.String()[1:]
				bWord := backwardDiag.String()[1:]

				if (fWord == word || fWord == reverseString(word)) && (bWord == word || bWord == reverseString(word)) {
					totalCount++
				} else {
					fmt.Printf("not here: %d %d\n", i, j)
				}

				fmt.Printf("forward: %s\nbackward: %s\n", fWord, bWord)
			}
		}
	}

	return fmt.Sprintf("total apperance of the word %s is = %d", word, totalCount)
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func matrixRightDown(m [][]string, rowStart, rowFinish, colStart, colFinish int) (string, string, string) {
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

	return horizontal.String(), vertical.String(), diagonal.String()
}

func matrixRightUp(m [][]string, rowStart, rowFinish, colStart, colFinish int) (string, string, string) {
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

	return horizontal.String(), vertical.String(), diagonal.String()
}

func matrixLeftDown(m [][]string, rowStart, rowFinish, colStart, colFinish int) (string, string, string) {
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

	return horizontal.String(), vertical.String(), diagonal.String()
}

func matrixLeftUp(m [][]string, rowStart, rowFinish, colStart, colFinish int) (string, string, string) {
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

	return horizontal.String(), vertical.String(), diagonal.String()
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

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
