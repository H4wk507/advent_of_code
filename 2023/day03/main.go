package main

import (
	"bufio"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"os"
)

type PartNumber struct {
	value, y, startX, endX int
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isSymbol(c byte) bool {
	return !isDigit(c) && c != '.'
}

func isPartNumber(board [][]byte, number, y, startPosX, endPosX int) bool {
	n, m := len(board), len(board[0])

	checkSymbol := func(row, col int) bool {
		return row >= 0 && row < n-1 && col >= 0 && col < m-1 && isSymbol(board[row][col])
	}

	checkChorizontal := func(row int) bool {
		if row >= 0 && row < n-1 {
			for i := startPosX; i <= endPosX; i++ {
				if isSymbol(board[row][i]) {
					return true
				}
			}
		}
		return false
	}

	return checkSymbol(y, startPosX-1) || checkSymbol(y-1, startPosX-1) || checkSymbol(y+1, startPosX-1) ||
		checkSymbol(y, endPosX+1) || checkSymbol(y-1, endPosX+1) || checkSymbol(y+1, endPosX+1) ||
		checkChorizontal(y-1) ||
		checkChorizontal(y+1)

}

func getGearProduct(y, x int, board [][]byte, partNumbers []PartNumber) int {
	n, m := len(board), len(board[0])

	linked := mapset.NewSet[int]()
	for _, dy := range [...]int{-1, 0, 1} {
		for _, dx := range [...]int{-1, 0, 1} {
			if dy == 0 && dx == 0 {
				continue
			}
			if dy+y >= 0 && dy+y < n && dx+x >= 0 && dx+x < m {
				if isDigit(board[dy+y][dx+x]) {
					for _, partNumber := range partNumbers {
						if dy+y == partNumber.y && x+dx <= partNumber.endX && x+dx >= partNumber.startX {
							linked.Add(partNumber.value)
						}
					}
				}
			}
		}
	}
	prod := 1
	if linked.Cardinality() == 2 {
		for v := range linked.Iter() {
			prod *= v
		}
		return prod
	} else {
		return 0
	}
}

func fileToBoard() [][]byte {
	file, err := os.Open("./input")
	if err != nil {
		fmt.Println("Failed opening the file", err)
		os.Exit(1)
	}
	defer file.Close()

	var board [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineLength := len(line)
		innerArray := make([]byte, lineLength)
		for i := 0; i < lineLength; i++ {
			innerArray[i] = line[i]
		}
		board = append(board, innerArray)
	}
	return board
}

func main() {
	sum := 0
	board := fileToBoard()
	var partNumbers []PartNumber
	for y, line := range board {
		startPos, endPos := 0, 0
		for i := 0; i < len(line); i++ {
			number := 0
			startPos = i
			for j := i; i < len(line) && isDigit(line[i]); j++ {
				number = 10*number + int(line[j]-'0')
				i++
			}
			endPos = i - 1
			if isPartNumber(board, number, y, startPos, endPos) {
				partNumbers = append(partNumbers, PartNumber{value: number, y: y, startX: startPos, endX: endPos})
				sum += number
			}
		}
	}
	gearRatioSum := 0
	for y, line := range board {
		for x, c := range line {
			if c == '*' {
				gearRatioSum += getGearProduct(y, x, board, partNumbers)
			}
		}
	}
	fmt.Println("Part #1:", sum)
	fmt.Println("Part #2:", gearRatioSum)
}
