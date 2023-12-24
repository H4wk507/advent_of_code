package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

type Grid [][]byte
type Direction int

const (
	RoundedRock = 'O'
	CubeRock    = '#'
	EmptySpace  = '.'
)

const (
	North Direction = iota
	East
	South
	West
)

func slideRockNorth(grid Grid, y, x int) {
	newY := y
	for newY > 0 && grid[newY-1][x] == EmptySpace {
		newY--
	}
	grid[y][x] = EmptySpace
	grid[newY][x] = RoundedRock
}

func slideRockWest(grid Grid, y, x int) {
	newX := x
	for newX > 0 && grid[y][newX-1] == EmptySpace {
		newX--
	}
	grid[y][x] = EmptySpace
	grid[y][newX] = RoundedRock
}

func slideRockSouth(grid Grid, y, x int) {
	newY := y
	for newY < len(grid)-1 && grid[newY+1][x] == EmptySpace {
		newY++
	}
	grid[y][x] = EmptySpace
	grid[newY][x] = RoundedRock
}

func slideRockEast(grid Grid, y, x int) {
	newX := x
	for newX < len(grid[0])-1 && grid[y][newX+1] == EmptySpace {
		newX++
	}
	grid[y][x] = EmptySpace
	grid[y][newX] = RoundedRock
}

func slideRocks(grid Grid, direction Direction) {
	if direction == North || direction == West {
		for y, row := range grid {
			for x, val := range row {
				if val == RoundedRock {
					switch direction {
					case North:
						slideRockNorth(grid, y, x)
					case West:
						slideRockWest(grid, y, x)
					}
				}
			}
		}
	} else {
		for y := len(grid) - 1; y >= 0; y-- {
			for x := len(grid[y]) - 1; x >= 0; x-- {
				if grid[y][x] == RoundedRock {
					switch direction {
					case South:
						slideRockSouth(grid, y, x)
					case East:
						slideRockEast(grid, y, x)
					}
				}
			}
		}
	}
}

func doCycle(grid Grid) {
	slideRocks(grid, North)
	slideRocks(grid, West)
	slideRocks(grid, South)
	slideRocks(grid, East)
}

func getTotalLoad(grid Grid) int {
	total := 0
	n := len(grid)
	for y, row := range grid {
		for _, val := range row {
			if val == RoundedRock {
				total += n - y
			}
		}
	}
	return total
}

func getCyclicTransformationOfN(mod, startCycleIdx, n int) int {
	for n >= mod {
		n -= (mod - startCycleIdx)
	}
	return n
}

func to2D(arr []byte, rowLength int) Grid {
	grid := make(Grid, 0)
	for i := 0; i < len(arr); i += rowLength {
		grid = append(grid, arr[i:i+rowLength])
	}
	return grid
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		fmt.Printf("Error opening file %v\n", err)
		return
	}
	defer file.Close()

	grid := make([][]byte, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}

	slideRocks(grid, North)
	fmt.Println("Part #1:", getTotalLoad(grid))

	pattern := make([]string, 0)
	n := 1_000_000_000
	var startCycleIdx int
	for i := 0; i <= n; i++ {
		collapsed := string(bytes.Join(grid, nil))
		if idx := slices.Index(pattern, collapsed); idx != -1 {
			startCycleIdx = idx
			break
		} else {
			pattern = append(pattern, collapsed)
		}
		doCycle(grid)
	}
	newN := getCyclicTransformationOfN(len(pattern), startCycleIdx, n)
	grid = to2D([]byte(pattern[newN]), len(grid[0]))
	fmt.Println("Part #2:", getTotalLoad(grid))
}
