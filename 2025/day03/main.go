package main

import (
	"bufio"
	"fmt"
	"os"
)

func p1(bank string) int64 {
	first := int64(0)
	second := int64(0)
	for digitIdx, digitChar := range bank {
		digit := int64(digitChar - '0')
		if digit > first && digitIdx < len(bank)-1 {
			first = digit
			second = 0
			continue
		}
		if first != 0 && digit > second {
			second = digit
		}
	}

	return first*10 + second
}

func p2(bank string) int64 {
	maxJoltage := int64(0)
	digitIndices := make([]int, 12)

	for i := range 12 {
		best := int64(0)
		for digitIdx, digitChar := range bank {
			digit := int64(digitChar - '0')
			lastDigitIdx := -1
			if i != 0 {
				lastDigitIdx = digitIndices[i-1]
			}
			if digitIdx > lastDigitIdx && digit > best && digitIdx < len(bank)-11+i {
				best = digit
				digitIndices[i] = digitIdx
			}
		}
		maxJoltage = maxJoltage*10 + best

	}

	return maxJoltage
}

func main() {
	f, _ := os.Open("inp")
	defer f.Close()

	js1 := int64(0)
	js2 := int64(0)

	s := bufio.NewScanner(f)
	for s.Scan() {
		bank := s.Text()
		js1 += p1(bank)
		js2 += p2(bank)
	}

	fmt.Printf("Part #1: %d\n", js1)
	fmt.Printf("Part #2: %d\n", js2)
}
