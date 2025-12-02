package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func isValid(n string) bool {
	if len(n)%2 != 0 {
		return true
	}

	m := len(n) / 2
	for i := 0; i < m; i++ {
		if n[i] != n[m+i] {
			return true
		}
	}
	return false
}

func hasPattern(n string, size int) bool {
	// assume len(n) % size == 0
	// size = 5

	// 222222, 2
	// len(n) = 6
	// pattern = [2, 2]

	pattern := make([]byte, size)
	for i := 0; i < len(n); i++ {
		pos := i % size
		if pattern[pos] == 0 {
			pattern[pos] = n[i]
		} else {
			if pattern[pos] != n[i] {
				return false
			}
		}
	}
	return true
}

func isValid2(n string) bool {
	if len(n) <= 1 {
		return true
	}

	l := len(n)
	for i := 1; i <= l/2; i++ {
		if l%i == 0 {
			if hasPattern(n, i) {
				return false
			}
		}
	}
	return true
}

func main() {
	var sum int64
	var sum2 int64
	f, _ := os.Open("inp")
	defer f.Close()

	byteContent, _ := io.ReadAll(f)
	content := string(byteContent)
	rs := strings.Split(content, ",")
	for _, r := range rs {
		pos := strings.Split(r, "-")
		l, _ := strconv.ParseInt(pos[0], 10, 64)
		h, _ := strconv.ParseInt(strings.TrimSuffix(pos[1], "\n"), 10, 64)
		for i := l; i <= h; i++ {
			iStr := fmt.Sprintf("%d", i)
			if !isValid(iStr) {
				sum += i
			}
			if !isValid2(iStr) {
				sum2 += i
			}
		}
	}
	fmt.Printf("Part #1: %d\n", sum)
	fmt.Printf("Part #2: %d\n", sum2)
}
