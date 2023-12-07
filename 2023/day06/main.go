package main

import (
	"bufio"
	"fmt"
	"math"
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

func joinNumber(numbers []int) int {
	joined := 0
	pow := 0
	for i := len(numbers) - 1; i >= 0; i-- {
		if i == len(numbers)-1 {
			joined += numbers[i]
		} else {
			pow += int(math.Floor(math.Log10(float64(numbers[i+1])))) + 1
			joined += numbers[i] * int(math.Pow(float64(10), float64(pow)))
		}
	}
	return joined
}

func part1(distances, times []int) int {
	var numberOfWays []int
	for i := 0; i < len(distances); i++ {
		cnt := 0
		for t := 0; t < times[i]; t++ {
			new_distance := t * (times[i] - t)
			if new_distance > distances[i] {
				cnt++
			}
		}
		numberOfWays = append(numberOfWays, cnt)
	}

	prod := 1
	for _, v := range numberOfWays {
		prod *= v
	}
	return prod

}

func part2(distances, times []int) int {
	joinedTime := joinNumber(times)
	joinedDistance := joinNumber(distances)

	cnt := 0
	for t := 0; t < joinedTime; t++ {
		new_distance := t * (joinedTime - t)
		if new_distance > joinedDistance {
			cnt++
		}
	}
	return cnt
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("error opening the file", err)
		return
	}

	var times []int
	var distances []int

	areTimes := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = stripBeginningOfLine(line)
		for {
			line = skipSpace(line)
			number := parseNumber(&line)
			if areTimes {
				times = append(times, number)
			} else {
				distances = append(distances, number)
			}
			if len(line) == 0 {
				areTimes = !areTimes
				break
			}
		}
	}

	fmt.Println("Part #1:", part1(distances, times))
	fmt.Println("Part #2:", part2(distances, times))
}
