package main

import (
	"bufio"
	"fmt"
	"os"
)

func colsEquals(grid [][]byte, c1, c2 int) bool {
	n := len(grid)
	for r := 0; r < n; r++ {
		if grid[r][c1] != grid[r][c2] {
			return false
		}
	}
	return true
}

func rowsEquals(grid [][]byte, r1, r2 int) bool {
	m := len(grid[0])
	for c := 0; c < m; c++ {
		if grid[r1][c] != grid[r2][c] {
			return false
		}
	}
	return true
}

func findReflectionCol(grid [][]byte, prevRefCol int) int {
	m := len(grid[0])
	for col := 0; col < m-1; col++ {
		found := true
		for l, r := col, col+1; l >= 0 && r < m; l, r = l-1, r+1 {
			if !colsEquals(grid, l, r) {
				found = false
				break
			}
		}
		if found {
			if col+1 != prevRefCol {
				return col + 1
			}
		}
	}
	return -1
}

func findReflectionRows(grid [][]byte, prevRefRow int) int {
	n := len(grid)
	for row := 0; row < n-1; row++ {
		found := true
		for u, l := row, row+1; u >= 0 && l < n; u, l = u-1, l+1 {
			if !rowsEquals(grid, u, l) {
				found = false
				break
			}
		}
		if found {
			if row+1 != prevRefRow {
				return row + 1
			}
		}
	}
	return -1
}

type GridAndReflection struct {
	grid           [][]byte
	refCol, refRow int
}

func swapSymbol(grid [][]byte, row, col int) {
	if grid[row][col] == '#' {
		grid[row][col] = '.'
	} else {
		grid[row][col] = '#'
	}
}

func getNewReflectionScore(gridAndReflection GridAndReflection) int {
	grid := gridAndReflection.grid
	n := len(grid)
	for i := 0; i < n; i++ {
		m := len(grid[i])
		for j := 0; j < m; j++ {
			swapSymbol(grid, i, j)
			col := findReflectionCol(grid, gridAndReflection.refCol)
			if col != -1 {
				return col
			}
			row := findReflectionRows(grid, gridAndReflection.refRow)
			if row != -1 {
				return row * 100
			}
			swapSymbol(grid, i, j)
		}
	}
	return 0
}

func part2(gridAndReflections []GridAndReflection) int {
	cnt := 0
	for _, gridAndReflection := range gridAndReflections {
		cnt += getNewReflectionScore(gridAndReflection)
	}
	return cnt
}

func part1(gridAndReflections []GridAndReflection) int {
	cnt := 0
	for i, gridAndReflection := range gridAndReflections {
		grid := gridAndReflection.grid
		col := findReflectionCol(grid, -1)
		if col != -1 {
			cnt += col
			gridAndReflections[i].refCol = col
			continue
		}
		row := findReflectionRows(grid, -1)
		if row != -1 {
			cnt += row * 100
			gridAndReflections[i].refRow = row
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

	gridAndReflections := make([]GridAndReflection, 0)
	grid := make([][]byte, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			gridAndReflections = append(gridAndReflections, GridAndReflection{grid, -1, -1})
			grid = make([][]byte, 0)
		} else {
			grid = append(grid, []byte(line))
		}
	}
	gridAndReflections = append(gridAndReflections, GridAndReflection{grid, -1, -1})

	fmt.Println("Part #1:", part1(gridAndReflections))
	fmt.Println("Part #2:", part2(gridAndReflections))
}
