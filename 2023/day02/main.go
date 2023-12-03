package main

import (
	"bufio"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"os"
)

const RED_LIMIT = 12
const GREEN_LIMIT = 13
const BLUE_LIMIT = 14

var colors = [...]string{"red", "blue", "green"}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func parseNumber(line *string) int {
	res := 0
	i := 0
	for isDigit((*line)[i]) {
		res = 10*res + int((*line)[i]-'0')
		i++
	}
	*line = (*line)[i:]
	return res
}

func parseColor(line *string) string {
	for _, color := range colors {
		j, k := 0, 0
		for j < len(*line) && k < len(color) && (*line)[j] == color[k] {
			j++
			k++
		}
		if k == len(color) {
			*line = (*line)[j:]
			return color
		}
	}
	return ""
}

func parseStartOfLine(line *string) string {
	i := 0
	for (*line)[i] != ':' {
		i++
	}
	i += 2 // skip ':' and ' '
	return (*line)[i:]
}

func skipChar(line string, char byte) string {
	i := 0
	for line[i] == char {
		i++
	}
	return line[i:]
}

type colorMaxes struct {
	red   int
	blue  int
	green int
}

func main() {
	readFile, err := os.Open("./input")

	if err != nil {
		fmt.Println("unable to read file: %v", err)
		os.Exit(1)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	invalidGames := mapset.NewSet[int]()
	lineno := 1
	powerSum := 0
	for fileScanner.Scan() {
		colorMaxes := colorMaxes{red: 0, blue: 0, green: 0}
		line := fileScanner.Text()
		line = parseStartOfLine(&line)
		for len(line) > 0 {
			number := parseNumber(&line)
			line = skipChar(line, ' ')
			color := parseColor(&line)
			if color == "red" {
				colorMaxes.red = max(colorMaxes.red, number)
				if number > RED_LIMIT {
					invalidGames.Add(lineno)
				}
			} else if color == "green" {
				colorMaxes.green = max(colorMaxes.green, number)
				if number > GREEN_LIMIT {
					invalidGames.Add(lineno)
				}
			} else if color == "blue" {
				colorMaxes.blue = max(colorMaxes.blue, number)
				if number > BLUE_LIMIT {
					invalidGames.Add(lineno)
				}
			}
			if len(line) != 0 {
				line = skipChar(line, ',')
				line = skipChar(line, ';')
				line = skipChar(line, ' ')
			}
		}
		powerSum += colorMaxes.red * colorMaxes.blue * colorMaxes.green
		lineno++
	}
	sum := ((lineno) * (lineno - 1)) / 2
	for v := range invalidGames.Iter() {
		sum -= v
	}
	fmt.Println("Part #1: ", sum)
	fmt.Println("Part #2: ", powerSum)

	readFile.Close()
}
