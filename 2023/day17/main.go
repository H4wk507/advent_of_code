package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Point struct {
	y, x int
}

type Move struct {
	position        Point
	movingDirection Direction
	straightMoves   int
}

func (q PQ[_]) Len() int           { return len(q) }
func (q PQ[_]) Less(i, j int) bool { return q[i].p < q[j].p }
func (q PQ[_]) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *PQ[T]) Push(x any)        { *q = append(*q, x.(pqi[T])) }
func (q *PQ[_]) Pop() (x any)      { x, *q = (*q)[len(*q)-1], (*q)[:len(*q)-1]; return x }
func (q *PQ[T]) GPush(v T, p int)  { heap.Push(q, pqi[T]{v, p}) }
func (q *PQ[T]) GPop() (T, int)    { x := heap.Pop(q).(pqi[T]); return x.v, x.p }

func print2d(grid [][]int) {
	for _, row := range grid {
		for _, val := range row {
			fmt.Print(val)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

type Direction string

const (
	North Direction = "North"
	East            = "East"
	South           = "South"
	West            = "West"
)

func getDirection(currPos, prevPos Point) Direction {
	dy, dx := currPos.y-prevPos.y, currPos.x-prevPos.x

	//fmt.Println(dy, dx)

	switch {
	case dy == 0 && dx == 0:
		return East
	case dy == 0 && dx == 1:
		return East
	case dy == 0 && dx == -1:
		return West
	case dy == 1:
		return South
	default:
		return North
	}
}

func getPositionFromDirection(y, x int, direction Direction) (int, int) {
	switch direction {
	case North:
		return y - 1, x
	case South:
		return y + 1, x
	case East:
		return y, x + 1
	default:
		return y, x - 1
	}
}

func isOppositeDirection(d1, d2 Direction) bool {
	switch d1 {
	case North:
		return d2 == South
	case South:
		return d2 == North
	case East:
		return d2 == West
	default:
		return d2 == East
	}
}

type pqi[T any] struct {
	v T
	p int
}

type PQ[T any] []pqi[T]

func printMap(m map[Point]int) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func main() {
	file, err := os.Open("./example")
	if err != nil {
		fmt.Printf("Error opening file %v\n", err)
	}
	defer file.Close()

	grid := make(map[Point]int)

	scanner := bufio.NewScanner(file)
	end := Point{y: 0, x: 0}
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		for x, c := range line {
			grid[Point{y: y, x: x}] = int(c - '0')
			end.x = len(line) - 1
		}
		end.y++
	}
	end.y--

	distances := make([][]int, end.y+1)
	for i := range distances {
		distances[i] = make([]int, end.x+1)
		for j := range distances[i] {
			distances[i][j] = 1 << 32
		}
	}

	// print2d(distances)
	distances[0][0] = 0
	distances[0][1] = grid[Point{y: 0, x: 1}]
	distances[1][0] = grid[Point{y: 1, x: 0}]

	visited := make(map[Point]int)
	queue := PQ[Move]{}
	queue.GPush(Move{position: Point{y: 0, x: 0}, movingDirection: East, straightMoves: 0}, distances[0][0])
	queue.GPush(Move{position: Point{y: 0, x: 0}, movingDirection: South, straightMoves: 0}, distances[0][0])
	directions := []Direction{North, West, South, East}
	for len(queue) > 0 {
		move, heat := queue.GPop()
		y, x := move.position.y, move.position.x

		// heat += grid[move.position]
		if y == end.y && x == end.x {
			fmt.Println(heat)
			print2d(distances)
			return
		}
		if _, exists := visited[move.position]; exists {
			// if v <= heat {
			continue
			// }
		}
		visited[move.position] = heat
		for _, direction := range directions {
			newY, newX := getPositionFromDirection(y, x, direction)
			newPosition := Point{y: newY, x: newX}

			if isOppositeDirection(direction, move.movingDirection) {
				continue
			}

			if _, exists := grid[newPosition]; !exists {
				continue
			}

			if direction == move.movingDirection {
				if move.straightMoves < 3 {
					fmt.Println(heat + grid[newPosition])
					if distances[newY][newX] > heat+grid[newPosition] {
						distances[newY][newX] = heat + grid[newPosition]
					}
					queue.GPush(Move{position: newPosition, movingDirection: direction, straightMoves: move.straightMoves + 1}, heat+grid[newPosition])
				}
			} else {
				if distances[newY][newX] > heat+grid[newPosition] {
					distances[newY][newX] = heat + grid[newPosition]
				}
				queue.GPush(Move{position: newPosition, movingDirection: direction, straightMoves: 1}, heat+grid[newPosition])
			}
		}
	}
}
