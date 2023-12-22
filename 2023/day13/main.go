package main

import (
	"bufio"
	"fmt"
	"os"
)

type Grid [][]byte

func getColsDistance(grid Grid, c1, c2 int) int {
	dist := 0
	for _, row := range grid {
		if row[c1] != row[c2] {
			dist++
		}
	}
	return dist
}

func getRowsDistance(grid Grid, r1, r2 int) int {
	m := len(grid[0])
	dist := 0
	for c := 0; c < m; c++ {
		if grid[r1][c] != grid[r2][c] {
			dist++
		}
	}
	return dist
}

func findReflectionCol(grid Grid, part int) int {
	m := len(grid[0])
	for col := 0; col < m-1; col++ {
		distSum := 0
		for l, r := col, col+1; l >= 0 && r < m && distSum <= part-1; l, r = l-1, r+1 {
			distSum += getColsDistance(grid, l, r)
		}
		if distSum == part-1 {
			return col + 1
		}
	}
	return -1
}

func findReflectionRow(grid Grid, part int) int {
	n := len(grid)
	for row := 0; row < n-1; row++ {
		distSum := 0
		for u, l := row, row+1; u >= 0 && l < n && distSum <= part-1; u, l = u-1, l+1 {
			distSum += getRowsDistance(grid, u, l)
		}
		if distSum == part-1 {
			return row + 1
		}
	}
	return -1
}

func solve(grids []Grid, part int) int {
	cnt := 0
	for _, grid := range grids {
		if col := findReflectionCol(grid, part); col != -1 {
			cnt += col
		} else if row := findReflectionRow(grid, part); row != -1 {
			cnt += row * 100
		}
	}
	return cnt
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		fmt.Printf("Error opnening file %v\n", err)
		return
	}
	defer file.Close()

	grids := make([]Grid, 0)
	grid := make(Grid, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			grids = append(grids, grid)
			grid = make(Grid, 0)
		} else {
			grid = append(grid, []byte(line))
		}
	}
	grids = append(grids, grid)

	fmt.Println("Part #1:", solve(grids, 1))
	fmt.Println("Part #2:", solve(grids, 2))
}
