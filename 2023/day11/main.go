package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a

}

func sum2d(arr [][]int) int {
	acc := 0
	for _, row := range arr {
		for _, v := range row {
			acc += v
		}
	}
	return acc
}

type Point struct {
	y, x int
}

func collectGalaxies(grid [][]byte) []Point {
	galaxies := make([]Point, 0)

	for y, row := range grid {
		for x, v := range row {
			if v == '#' {
				galaxies = append(galaxies, Point{y: y, x: x})
			}
		}
	}

	return galaxies
}

func calculateDistances(grid [][]byte, galaxies []Point) [][]int {
	dist := make([][]int, len(galaxies))

	for y := range galaxies {
		dist[y] = make([]int, len(galaxies))
	}

	for i, g1 := range galaxies {
		for j, g2 := range galaxies {
			dist[i][j] = abs(g1.x-g2.x) + abs(g1.y-g2.y)
		}
	}
	return dist
}

func adjustGalaxiesCords(grid [][]byte, galaxies []Point, part int) []Point {
	emptyRows := make([]int, 0)
	emptyCols := make([]int, 0)

	for y, row := range grid {
		isEmpty := true
		for _, v := range row {
			if v == '#' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyRows = append(emptyRows, y)
		}
	}

	for x := 0; x < len(grid[0]); x++ {
		isEmpty := true
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == '#' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyCols = append(emptyCols, x)
		}
	}

	var expansionRate int
	if part == 1 {
		expansionRate = 1
	} else {
		expansionRate = 1_000_000 - 1
	}
	newGalaxies := make([]Point, len(galaxies))
	copy(newGalaxies, galaxies)
	for _, emptyRow := range emptyRows {
		for i, galaxy := range galaxies {
			if galaxy.y > emptyRow {
				newGalaxies[i].y = newGalaxies[i].y + expansionRate
			}
		}
	}

	for _, emptyCol := range emptyCols {
		for i, galaxy := range galaxies {
			if galaxy.x > emptyCol {
				newGalaxies[i].x = newGalaxies[i].x + expansionRate
			}
		}
	}

	return newGalaxies
}

func part1(grid [][]byte, galaxies []Point) int {
	galaxies = adjustGalaxiesCords(grid, galaxies, 1)
	dist := calculateDistances(grid, galaxies)
	return sum2d(dist) / 2
}

func part2(grid [][]byte, galaxies []Point) int {
	galaxies = adjustGalaxiesCords(grid, galaxies, 2)
	dist := calculateDistances(grid, galaxies)
	return sum2d(dist) / 2
}

func main() {
	file, err := os.Open("input")
	if err != nil {
        fmt.Printf("error opening file: %v\n", err)
		return
	}
    defer file.Close()

	grid := make([][]byte, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}

	galaxies := collectGalaxies(grid)
	fmt.Println("Part #1:", part1(grid, galaxies))
	fmt.Println("Part #2:", part2(grid, galaxies))
}
