package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	y, x int
}

type Direction string

const (
	Up    Direction = "U"
	Right Direction = "R"
	Down  Direction = "D"
	Left  Direction = "L"
)

func getDeltasFromDirection(dir Direction) (int, int) {
	switch dir {
	case Up:
		return -1, 0
	case Down:
		return 1, 0
	case Right:
		return 0, 1
	case Left:
		return 0, -1
	default:
		fmt.Println("unreachable")
		return 0, 0
	}
}

var numberToDirection = map[byte]Direction{'0': "R", '1': "D", '2': "L", '3': "U"}

func parseHex(hex string) (Direction, int) {
	numberStr := hex[:5]
	number, _ := strconv.ParseInt(numberStr, 16, 0)
	direction := numberToDirection[hex[5]]
	return direction, int(number)
}

type State struct {
	dir  Direction
	dist int
	hex  string
}

func solve(states []State, part int) int {
	curr, area := Point{0, 0}, 0

	getDirectionAndDist := func(state State) (Direction, int) {
		if part == 1 {
			return state.dir, state.dist
		}
		return parseHex(state.hex)
	}

	for _, state := range states {
		dir, dist := getDirectionAndDist(state)
		dy, dx := getDeltasFromDirection(dir)
		next := Point{curr.y + dy*dist, curr.x + dx*dist}
		area += ((curr.x*next.y - next.x*curr.y) + dist)
		curr = next
	}
	return area/2 + 1
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	re := regexp.MustCompile(`([URDL]) (\d+) \(#([0-9a-f]+)\)`)
	states := make([]State, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		dir, distStr, hex := Direction(match[1]), match[2], match[3]
		dist, _ := strconv.Atoi(distStr)
		states = append(states, State{dir, dist, hex})
	}
	fmt.Println("Part #1:", solve(states, 1))
	fmt.Println("Part #2:", solve(states, 2))
}

// https://en.wikipedia.org/wiki/Shoelace_formula
