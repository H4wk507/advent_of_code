package main

import (
	"bufio"
	"fmt"
	"os"
)

type Grid [][]Tile

type Direction int

const (
	North Direction = 1 << iota
	East            = 1 << iota
	South           = 1 << iota
	West            = 1 << iota
)

type Tile struct {
	val       byte
	visitedBy Direction
}

type Beam struct {
	y, x      int
	direction Direction
}

func getEnergizedTiles(grid Grid) int {
	total := 0
	for _, row := range grid {
		for _, gridVal := range row {
			if gridVal.visitedBy > 0 {
				total++
			}
		}
	}
	return total
}

func handleTile(grid *Grid, y, x int, beams *[]Beam, beam *Beam) bool {
	cell := &(*grid)[y][x]

	// This path has already been searched coming from that beam.direction
	// so there is no point in searching it again.
	if cell.visitedBy&beam.direction != 0 {
		return false
	}
	cell.visitedBy |= beam.direction

	switch cell.val {
	case '/':
		handleSlashMirror(beam)
	case '\\':
		handleBackslashMirror(beam)
	case '-':
		handleHorizontalSplitter(beams, beam, y, x)
	case '|':
		handleVerticalSplitter(beams, beam, y, x)
	default:
		cell.visitedBy |= beam.direction
	}
	return true
}

func handleSlashMirror(beam *Beam) {
	switch beam.direction {
	case East, West:
		beam.direction >>= 1
	default:
		beam.direction <<= 1
	}
}

func handleBackslashMirror(beam *Beam) {
	switch beam.direction {
	case East:
		beam.direction <<= 1
	case South:
		beam.direction >>= 1
	case West:
		beam.direction >>= 3
	case North:
		beam.direction <<= 3
	}
}

func handleHorizontalSplitter(beams *[]Beam, beam *Beam, y, x int) {
	if beam.direction == North || beam.direction == South {
		beam.direction = West
		*beams = append(*beams, Beam{y: y, x: x + 1, direction: East})
	}
}

func handleVerticalSplitter(beams *[]Beam, beam *Beam, y, x int) {
	if beam.direction == East || beam.direction == West {
		beam.direction = North
		*beams = append(*beams, Beam{y: y + 1, x: x, direction: South})
	}
}

func moveBeam(beam *Beam) {
	switch beam.direction {
	case North:
		beam.y--
	case East:
		beam.x++
	case South:
		beam.y++
	case West:
		beam.x--
	}
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		fmt.Printf("Error opening the file %v\n", err)
		return
	}
	defer file.Close()

	grid := make(Grid, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]Tile, len(line))
		for i, c := range line {
			row[i] = Tile{val: byte(c), visitedBy: 0}
		}
		grid = append(grid, row)
	}

	beams := []Beam{{y: 0, x: 0, direction: East}}
	for len(beams) > 0 {
		beam := beams[0]
		beams = beams[1:]

		for {
			y, x := beam.y, beam.x
			if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
				break
			}
			if !handleTile(&grid, y, x, &beams, &beam) {
				break
			}
			moveBeam(&beam)
		}
	}
	fmt.Println("Part #1:", getEnergizedTiles(grid))
}
