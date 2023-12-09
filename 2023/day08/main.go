package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(x, y int) int {
	return (x * y) / gcd(x, y)
}

func skipChar(line string, c byte) string {
	i := 0
	for i < len(line) && line[i] == c {
		i++
	}
	return line[i:]
}

func getNodeVal(line *string) string {
	lineVal := *line
	if len(lineVal) < 3 {
		return ""
	}

	nodeVal := string(lineVal[0]) + string(lineVal[1]) + string(lineVal[2])
	*line = lineVal[3:]
	return nodeVal
}

type Node struct {
	val         string
	left, right *Node
}

type Graph struct {
	nodes map[string]*Node
}

func part1(currPosition, dstPosition *Node, directions string, graph *Graph) int {
	nsteps := 0
	for currPosition.val != dstPosition.val {
		if directions[nsteps%len(directions)] == 'L' {
			currPosition = graph.nodes[currPosition.val].left
		} else {
			currPosition = graph.nodes[currPosition.val].right
		}
		nsteps++
	}
	return nsteps
}

func all(arr []bool) bool {
	for _, a := range arr {
		if !a {
			return false
		}
	}
	return true
}

/**
 * Ok, so the resulting number is too big to calculate using normal graph
 * traversal, so we have to be smarter. We can assume that all the ghosts
 * will eventually meet at the nodes ending with 'Z'. We can calculate
 * number of steps it takes a particular ghost to reach the ending node and
 * ASSUME that this forms a cycle of that length (it's stupid but it turns out
 * to be true and only of that assumption this solution works). So now it's
 * pretty easy to see that if we calculate Least Common Multiply of all that
 * steps we get the number of steps required for all the ghosts to meet at
 * the ending nodes.
 **
 */
func part2(currPositions []*Node, directions string, graph *Graph) int {
	var hasReachedCycle = make([]bool, len(currPositions))
	nsteps := 0
	result := 1
	for !all(hasReachedCycle) {
		direction := directions[nsteps%len(directions)]
		for idx, currPosition := range currPositions {
			if direction == 'L' {
				*currPositions[idx] = *graph.nodes[currPosition.val].left
			} else {
				*currPositions[idx] = *graph.nodes[currPosition.val].right
			}
			if currPositions[idx].val[2] == 'Z' && !hasReachedCycle[idx] {
				hasReachedCycle[idx] = true
				result = lcm(result, nsteps+1)
			}
		}
		nsteps++
	}
	return result
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("error while opening file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineno := 0
	var currPositions []*Node
	var currPosition *Node
	var dstPosition Node
	g := Graph{nodes: make(map[string]*Node, 0)}
	var directions string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			lineno++
			continue
		}
		if lineno == 0 {
			directions = line
		} else {
			val := getNodeVal(&line)
			line = skipChar(line, ' ')
			line = skipChar(line, '=')
			line = skipChar(line, ' ')
			line = skipChar(line, '(')
			left := getNodeVal(&line)
			line = skipChar(line, ',')
			line = skipChar(line, ' ')
			right := getNodeVal(&line)
			line = skipChar(line, ')')
			node := Node{val: val, left: &Node{val: left, left: nil, right: nil}, right: &Node{val: right, left: nil, right: nil}}
			g.nodes[val] = &node
			if val == "ZZZ" {
				dstPosition = node
			}
			if val[2] == 'A' {
				currPositions = append(currPositions, &node)
			}
			if val == "AAA" {
				currPosition = &node
			}
		}
		lineno++
	}
	fmt.Println("Part #1:", part1(currPosition, &dstPosition, directions, &g))
	fmt.Println("Part #2:", part2(currPositions, directions, &g))
}
