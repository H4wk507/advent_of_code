package main

import (
	"bufio"
	"fmt"
	"os"
)

var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func getNumberFromLine(line string) int {
	leftFound, rightFound := false, false
	leftDigit, rightDigit := 0, 0

	for l, r := 0, len(line)-1; !leftFound || !rightFound; l, r = l+1, r-1 {
		if !leftFound && isDigit(line[l]) {
			leftFound = true
			leftDigit = int(line[l] - '0')
		}
		if !rightFound && isDigit(line[r]) {
			rightFound = true
			rightDigit = int(line[r] - '0')
		}
	}
	return leftDigit*10 + rightDigit
}

func getNumberFromLine2(line string) int {
	leftFound, rightFound := false, false
	leftDigit, rightDigit := 0, 0

	for l, r := 0, len(line)-1; !leftFound || !rightFound; l, r = l+1, r-1 {
		if !leftFound {
			if isDigit(line[l]) {
				leftFound = true
				leftDigit = int(line[l] - '0')
			} else {
				for word, digit := range numbers {
					n := len(word)
					i := 0

					for i < n && line[l] == word[i] {
						l++
						i++
					}

					if i == n {
						leftFound = true
						leftDigit = digit
						break
					} else {
						l -= i
					}
				}
			}
		}
		if !rightFound {
			if isDigit(line[r]) {
				rightFound = true
				rightDigit = int(line[r] - '0')
			} else {
				for word, digit := range numbers {
					n := len(word)
					i := n - 1

					for i >= 0 && line[r] == word[i] {
						r--
						i--
					}

					if i == -1 {
						rightFound = true
						rightDigit = digit
						break
					} else {
						r += n - 1 - i
					}
				}
			}
		}
	}
	return leftDigit*10 + rightDigit
}

func solve(inputFile string, lineParser func(string) int) int {
	readFile, err := os.Open(inputFile)

	if err != nil {
		fmt.Println("unable to read file: %v", err)
		os.Exit(1)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		sum += lineParser(line)
	}
	readFile.Close()

	return sum
}

func main() {
	fmt.Println("Part #1: ", solve("./input", getNumberFromLine))
	fmt.Println("Part #2: ", solve("./input2", getNumberFromLine2))
}
