package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("inp")
	defer f.Close()
	s := bufio.NewScanner(f)
	c := 50
	ans1 := 0
	ans2 := 0
	for s.Scan() {
		l := s.Text()
		d := l[0]
		n, _ := strconv.Atoi(l[1:])
		switch d {
		case 'L':
			if c == 0 {
				ans2 += n / 100
			} else if n >= c {
				ans2 += 1 + (n-c)/100
			}
			c = (100 + ((c - n) % 100)) % 100
		case 'R':
			ans2 += (c + n) / 100
			c = (c + n) % 100
		}
		if c == 0 {
			ans1++
		}
	}
	fmt.Printf("Part #1: %d\n", ans1)
	fmt.Printf("Part #2: %d\n", ans2)
}
