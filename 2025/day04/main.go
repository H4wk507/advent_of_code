package main

import (
	"bufio"
	"fmt"
	"os"
)

func getNeighborsCount(grid [][]byte, posY, posX int) int {
	var delta = [...]int{-1, 0, 1}
	cnt := 0

	for _, dy := range delta {
		for _, dx := range delta {
			if dy == 0 && dx == 0 {
				continue
			}

			y := posY + dy
			x := posX + dx

			if y < 0 || y > len(grid)-1 || x < 0 || x > len(grid[0])-1 {
				continue
			}

			if grid[y][x] == '@' {
				cnt++
			}
		}
	}

	return cnt
}

func main() {
	f, _ := os.Open("inp")
	defer f.Close()

	grid := make([][]byte, 0)

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		grid = append(grid, []byte(line))
	}

	p1 := 0
	p2 := 0
	toRemove := make([][2]int, 0)
	for i := 0; ; i++ {
		toRemove = toRemove[:0]

		for y := range grid {
			for x := range grid[y] {
				if grid[y][x] == '@' && getNeighborsCount(grid, y, x) < 4 {
					toRemove = append(toRemove, [2]int{y, x})
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}

		if i == 0 {
			p1 = len(toRemove)
		}
		p2 += len(toRemove)

		for _, pos := range toRemove {
			grid[pos[0]][pos[1]] = '.'
		}
	}

	fmt.Printf("Part #1: %d\n", p1)
	fmt.Printf("Part #2: %d\n", p2)
}
