package main

import (
    "fmt"
    "bufio"
    "os"
)

func getStartIndices(grid [][]byte) (int, int) {
    for i, row := range grid {
        for j, c := range row {
            if c == 'S' {
                return i, j
            }
        }
    }
    return -1, -1
}

type Point struct {
    row, col int
}

func getDeltaArr(c byte) [][]int {
    switch c {
    case '|':
        return [][]int{{-1, 0}, {1, 0}}
    case '-':
        return [][]int{{0, 1}, {0, -1}}
    case 'L':
        return [][]int{{0, 1}, {-1, 0}}
    case 'J':
        return [][]int{{-1, 0}, {0, -1}}
    case '7':
        return [][]int{{1, 0}, {0, -1}}
    case 'F':
        return [][]int{{1, 0}, {0, 1}}
    case 'S':
        return [][]int{{0, 1}, {-1, 0}}
        // return [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
    default:
        return [][]int{}
    }

}

func printGrid(grid [][]int) {
    for _, row := range grid {
        for _, v := range row {
            fmt.Printf("%4d", v)
        }
        fmt.Println()
    }
}

func solve(grid [][]byte) int {
    i, j := getStartIndices(grid)

    dist := make([][]int, len(grid))
    visited := make([][]bool, len(grid))
    for y := range grid {
        dist[y] = make([]int, len(grid[y]))
        visited[y] = make([]bool, len(grid[y]))
        for x := range grid[y] {
            dist[y][x] = -1
        }
    }

    dist[i][j] = 0
    visited[i][j] = true
    maxDist := 0

    bfs := func(start Point) {
        queue := []Point{start}
        for len(queue) > 0 {
            node := queue[0]
            queue = queue[1:]

            deltas := getDeltaArr(grid[node.row][node.col])
            for _, delta := range deltas {
                dy, dx := delta[0], delta[1]
                newY, newX := node.row + dy, node.col + dx
                if newY >= 0 && newY < len(grid) && newX >= 0 && newX < len(grid[0]) && !visited[newY][newX] {
                    currDist := dist[node.row][node.col] + 1
                    dist[newY][newX] = currDist
                    if currDist > maxDist {
                        maxDist = currDist
                    }
                    visited[newY][newX] = true
                    queue = append(queue, Point{row: newY, col: newX})
                }
            }
        }
    }

    bfs(Point{row: i, col: j})
    return maxDist
}

func main() {
    file, err := os.Open("input")
    if err != nil {
        fmt.Println("error opening the file", err)
        return
    }

    grid := make([][]byte, 0)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        inner := make([]byte, len(line))
        for i := 0; i < len(line); i++ {
            inner[i] = line[i]
        }
        grid = append(grid, inner)
    }
    fmt.Println(solve(grid))
}
