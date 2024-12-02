package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func toIntArray(arr []string) []int {
    intArr := make([]int, len(arr))
    for i, num := range arr {
        numInt, _ := strconv.Atoi(num)
        intArr[i] = numInt
    }
    return intArr
}

const Broken = '#'
const Operational = '.'
const Unknown = '?'

func isValidComb(row []byte, groups []int) bool {
    brokenSubstrLength := 0
    groupIdx := 0

    for _, c := range row {
        if c == Unknown {
            return false
        } else if c == Broken {
            brokenSubstrLength++
        } else {
            if brokenSubstrLength > 0 {
                if groupIdx >= len(groups) || groups[groupIdx] != brokenSubstrLength {
                    return false
                }
                groupIdx++
            }
            brokenSubstrLength = 0
        }
    }
    if groupIdx == len(groups) - 1 {
        return brokenSubstrLength > 0  && brokenSubstrLength == groups[groupIdx]
    }
    return groupIdx == len(groups) && brokenSubstrLength == 0
}

type StackVal struct {
    row []byte
    start int
}

func solve(row []byte, groups []int) int {
    stack := make([]StackVal, 0)
    stack = append(stack, StackVal{row: row, start: 0})
    cnt := 0
    for len(stack) > 0 {
        r := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        if isValidComb(r.row, groups) {
            cnt++
        }

        for i := r.start; i < len(r.row); i++ {
            if r.row[i] == Unknown {
                newR := make([]byte, len(r.row))
                newR2 := make([]byte, len(r.row))
                copy(newR, r.row)
                copy(newR2, r.row)
                newR[i] = Broken
                stack = append(stack, StackVal{row: newR, start: i+1})
                newR2[i] = Operational
                stack = append(stack, StackVal{row: newR2, start: i+1})
            }
        }
    }
    return cnt
}

func isAllUnknown(arr []byte) bool {
    for _, v := range arr {
        if v != Unknown {
            return false
        }
    }
    return true
}

func isAllBroken(arr []byte) bool {
    for _, v := range arr {
        if v != Broken && v != Unknown {
            return false
        }
    }
    return true
}


func solve2(row []byte, groups []int, start, groupIdx int) int {
    score := 0
    i := start
    fmt.Println(string(row), i, groupIdx)
    for groupIdx < len(groups) && i <= len(row)-groups[groupIdx] {
        slice := row[i:i+groups[groupIdx]]
        if isAllBroken(slice) {
            groupIdx++ 
            i += len(slice) + 1
        } else if isAllUnknown(slice) {
            i += len(slice) + 1
            score += solve2(row, groups, i, groupIdx)
            groupIdx++ 
        } else {
            i++
        }
    }
    if groupIdx >= len(groups) && i >= len(row) {
        return score + 1
    }
    return score
}

func main() {
    file, err := os.Open("input")
    if err != nil {
        fmt.Printf("Error opening the file %v\n", err)
        return
    }
    defer file.Close()

    totalScore := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        splitted := strings.Split(scanner.Text(), " ")
        row, groups := []byte(splitted[0]), toIntArray(strings.Split(splitted[1], ","))
        score := solve2(row, groups, 0, 0)
        totalScore += score
        fmt.Println(score)
    }
    fmt.Println(totalScore)
}
