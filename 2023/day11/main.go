package main

import (
    "fmt"
    "bufio"
    "os"
)

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
    
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func print2dArr(arr [][]byte) {
    for _, row := range arr {
        fmt.Println(string(row))
    }
}

func addRow(original [][]byte, newRow []byte, y int) [][]byte {
    newArray := make([][]byte, len(original)+1)
    newIdx := 0
    for i := range original {
        newArray[i] = original[i]
        if i == y {
            newArray[i+1] = newRow
            newIdx = i+1
            break
        } 
    }
    for i := newIdx; i < len(original); i++ {
        newArray[i+1] = original[i]
    }
    return newArray
}

func addColumn(original [][]byte, newColumn []byte, x int) [][]byte {
	newArray := make([][]byte, len(original))
    for i := range newArray {
        newRow := make([]byte, len(original[i])+1)
        for j := 0; j < x; j++ {
            newRow[j] = original[i][j]
        }
        newRow[x] = newColumn[i]
        for j := x+1; j < len(newRow); j++ {
            newRow[j] = original[i][j-1]
        }
        newArray[i] = newRow
	}
	return newArray
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

func expandGrid(grid [][]byte) [][]byte {
    expandedGrid := grid

    /* row check */
    for y, row := range grid {
        expand := true
        for _, v := range row {
            if v != '.' {
                expand = false
                break
            }
        }
        if expand {
            expandedGrid = addRow(expandedGrid, row, y)
        }
    }

    /* column check */
    cnt := 0 
    expandedGrid2 := expandedGrid
    for x := 0; x < len(expandedGrid[0]); x++ {
        col := make([]byte, 0)
        expand := true
        for y := 0; y < len(expandedGrid); y++ {
            col = append(col, expandedGrid[y][x])
            if expandedGrid[y][x] != '.' {
                expand = false
            }
        }
        if expand {
            expandedGrid2 = addColumn(expandedGrid2, col, x+cnt)
            cnt++;
        }
    }

    return expandedGrid2
}

func calculateDistances(grid [][]byte, galaxies []Point) [][]int {
    dist := make([][]int, len(galaxies))

    for y := range galaxies {
        dist[y] = make([]int, len(galaxies))
    }

    for i, g1 := range galaxies {
        for j, g2 := range galaxies {
            dist[i][j] = abs(g1.x - g2.x) + abs(g1.y - g2.y)
        }
    }
    return dist
}

func main() {
    file, err := os.Open("input");
    if err != nil {
        return
    }

    grid := make([][]byte, 0)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        grid = append(grid, []byte(scanner.Text()))
    }
    
    expandedGrid := expandGrid(grid)
    galaxies := collectGalaxies(expandedGrid)
    dist := calculateDistances(expandedGrid, galaxies)
    acc := 0
    for _, row := range dist {
        for _, v := range row {
            acc += v
        }
    }
    print2dArr(expandedGrid)
    //fmt.Println(acc / 2)
    
}
