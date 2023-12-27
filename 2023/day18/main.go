package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	y, x int
}

type Direction byte

const (
	Up    Direction = 'U'
	Right Direction = 'R'
	Down  Direction = 'D'
	Left  Direction = 'L'
)

func parseNumber(str string) int {
	ret := 0
	for i := 0; str[i] >= '0' && str[i] <= '9'; i++ {
		ret = ret*10 + int(str[i]-'0')

	}
	return ret
}

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

func main() {
	file, err := os.Open("./input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	curr, area := Point{0, 0}, 0
	for scanner.Scan() {
		line := scanner.Text()
		dir, dist := Direction(line[0]), parseNumber(line[2:])
		dy, dx := getDeltasFromDirection(dir)
		// Shoelace formula
		next := Point{curr.y + dy*dist, curr.x + dx*dist}
		area += ((curr.x*next.y - next.x*curr.y) + dist)
		curr = next
	}
	fmt.Println(area/2 + 1)
}
