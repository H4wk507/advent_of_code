package main

import (
	"bufio"
	"fmt"
	"github.com/gammazero/deque"
	"os"
	"strconv"
	"strings"
)

func allZeros(deq *deque.Deque[int]) bool {
	for i := 0; i < deq.Len(); i++ {
		if deq.At(i) != 0 {
			return false
		}
	}
	return true
}

func createTriangle(allNumbers []*deque.Deque[int], part int) []*deque.Deque[int] {
	for {
		triangleHeight := len(allNumbers)
		upperArr := allNumbers[triangleHeight-1]
		lowerArr := deque.New[int](0, upperArr.Len()-1)
		for i := 0; i < upperArr.Len()-1; i++ {
			lowerArr.PushBack(upperArr.At(i+1) - upperArr.At(i))
		}
		if allZeros(lowerArr) {
			if part == 1 {
				lowerArr.PushBack(0)
			} else {
				lowerArr.PushFront(0)
			}
			allNumbers = append(allNumbers, lowerArr)
			break
		}
		allNumbers = append(allNumbers, lowerArr)
	}
	return extrapolateTriangle(allNumbers, part)
}

func extrapolateTriangle(allNumbers []*deque.Deque[int], part int) []*deque.Deque[int] {
	triangleHeight := len(allNumbers)
	for i := triangleHeight - 1; i > 0; i-- {
		lowerArr := allNumbers[i]
		upperArr := allNumbers[i-1]
		if part == 1 {
			upperArr.PushBack(upperArr.Back() + lowerArr.Back())
		} else {
			upperArr.PushFront(upperArr.Front() - lowerArr.Front())
		}
		allNumbers[i-1] = upperArr
	}
	return allNumbers
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("error while opening file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	score1, score2 := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Split(line, " ")
		allNumbers := make([]*deque.Deque[int], 0)
		var numbers deque.Deque[int]
		for _, el := range lineArr {
			number, _ := strconv.Atoi(el)
			numbers.PushBack(number)
		}
		allNumbers = append(allNumbers, &numbers)
		part1Triangle := createTriangle(allNumbers, 1)
		part2Triangle := createTriangle(allNumbers, 2)
		score1 += part1Triangle[0].Back()
		score2 += part2Triangle[0].Front()
	}
	fmt.Println("Part #1:", score1)
	fmt.Println("Part #2:", score2)
}
