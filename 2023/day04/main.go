package main

import (
	"bufio"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"os"
)

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func stripBeginningOfLine(line string) string {
	i := 0
	for line[i] != ':' {
		i++
	}
	return line[i+1:]
}

func skipSpace(line string) string {
	i := 0
	for line[i] == ' ' {
		i++
	}
	return line[i:]
}

func parseNumber(line *string) int {
	i, number := 0, 0
	for i < len(*line) && isDigit((*line)[i]) {
		number = 10*number + int((*line)[i]-'0')
		i++
	}
	*line = (*line)[i:]
	return number
}

// TODO: change to dp
func countCardsRec(arr []int, cardToMatches map[int]int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	res := arr
	for i := 0; i < len(arr); i++ {
		var children []int
		for j := arr[i] + 1; j <= cardToMatches[arr[i]]+arr[i]; j++ {
			children = append(children, j)
		}
		res = append(res, countCardsRec(children, cardToMatches)...)
	}
	return res
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		fmt.Println("Error opening the file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineno, totalScore := 1, 0
	var cardToMatches = make(map[int]int)
	var cards []int
	for scanner.Scan() {
		areWinningCards := true
		userCards, winningCards := mapset.NewSet[int](), mapset.NewSet[int]()
		line := scanner.Text()
		line = stripBeginningOfLine(line)
		line = skipSpace(line)
		for {
			number := parseNumber(&line)
			if areWinningCards {
				winningCards.Add(number)
			} else {
				userCards.Add(number)
			}
			if len(line) == 0 {
				break
			}
			line = skipSpace(line)
			if line[0] == '|' {
				areWinningCards = false
				line = skipSpace(line[1:])
			}
		}
		nmatches := userCards.Intersect(winningCards).Cardinality()
		totalScore += (1 << nmatches) >> 1

		cardToMatches[lineno] = nmatches
		cards = append(cards, lineno)
		lineno++
	}
	fmt.Println("Part #1", totalScore)
	fmt.Println("Part #2", len(countCardsRec(cards, cardToMatches)))
}
