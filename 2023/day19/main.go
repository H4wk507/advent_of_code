package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Workflows map[string][]Expression
type Part map[string]int

type Expression struct {
	lhs, op, destination string
	rhs                  int
}

func parseExpression(exp string) Expression {
	expressionRegex := regexp.MustCompile(`([xmas])([<>])(\d+)`)
	splitted := strings.Split(exp, ":")
	if len(splitted) == 1 {
		return Expression{destination: splitted[0]}
	}
	expression, destination := splitted[0], splitted[1]
	match := expressionRegex.FindStringSubmatch(expression)
	lhs, op, rhsStr := match[1], match[2], match[3]
	rhs, _ := strconv.Atoi(rhsStr)
	return Expression{lhs, op, destination, rhs}
}

func getWorkflows(lines *[]string) Workflows {
	workflows := make(Workflows)

	workflowRegex := regexp.MustCompile(`([a-z]+)\{(.+,?)+\}`)
	i := 0
	for _, line := range *lines {
		i++
		if len(line) == 0 {
			break
		}
		match := workflowRegex.FindStringSubmatch(line)
		name, expression := match[1], match[2]
		expressions := strings.Split(expression, ",")
		for _, exp := range expressions {
			workflows[name] = append(workflows[name], parseExpression(exp))
		}
	}
	*lines = (*lines)[i:]
	return workflows
}

func getParts(lines []string) []Part {
	parts := make([]Part, 0)
	for _, line := range lines {
		part := make(map[string]int)
		groups := strings.Split(line[1:len(line)-1], ",")
		for _, group := range groups {
			equalSplit := strings.Split(group, "=")
			category, valStr := equalSplit[0], equalSplit[1]
			val, _ := strconv.Atoi(valStr)
			part[category] = val
		}
		parts = append(parts, part)
	}
	return parts
}

func sumPart(part Part) int {
	total := 0
	for _, v := range part {
		total += v
	}
	return total
}

func part1(workflows Workflows, parts []Part) int {
	total := 0
	for _, part := range parts {
		workflowName := "in"
		i := 0
		for i < len(workflows[workflowName]) {
			workflow := workflows[workflowName][i]
			isExpTrue := false
			switch workflow.op {
			case "<":
				isExpTrue = part[workflow.lhs] < workflow.rhs
			case ">":
				isExpTrue = part[workflow.lhs] > workflow.rhs
			case "":
				isExpTrue = true
			}
			if isExpTrue {
				if workflow.destination == "R" {
					break
				} else if workflow.destination == "A" {
					total += sumPart(part)
					break
				}
				workflowName = workflow.destination
				i = 0
			} else {
				i++
			}
		}
	}
	return total
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	workflows := getWorkflows(&lines)
	parts := getParts(lines)
	fmt.Println("Part #1:", part1(workflows, parts))
}
