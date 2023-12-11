package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "github.com/gammazero/deque"
)

func allZeros(deq *deque.Deque[int]) bool {
    for i := 0; i < deq.Len(); i++ {
        if deq.At(i) != 0 {
            return false
        }
    }
    return true
}

func main() {
    file, err := os.Open("input")
    if err != nil {
        fmt.Println("error while opening file", err)
        return
    }

    scanner := bufio.NewScanner(file)
    score := 0
    for scanner.Scan() {
        line := scanner.Text()
        lineArr := strings.Split(line, " ")
        allNumbers := make([]deque.Deque[int], 0)
        var numbers deque.Deque[int]
        for _, el := range lineArr {
            number, _ := strconv.Atoi(el)
            numbers.PushBack(number)
        }
        allNumbers = append(allNumbers, numbers)
        for {
            upperArr := allNumbers[len(allNumbers)-1]
            lowerArr := deque.New[int](0, allNumbers[len(allNumbers)-1].Len()-1)
            for i := 0; i < upperArr.Len()-1; i++ {
                lowerArr.PushBack(upperArr.At(i+1) - upperArr.At(i))
            }
            if allZeros(lowerArr) {
                lowerArr.PushBack(0)
                allNumbers = append(allNumbers, *lowerArr)
                break
            }
            allNumbers = append(allNumbers, *lowerArr)
        }
        // extrapolate
        for i := len(allNumbers) - 1; i > 0; i-- {
            lowerArr := allNumbers[i] 
            upperArr := allNumbers[i-1] 
            lastVal := upperArr.Back()
            upperArr.PushBack(lastVal+lowerArr.Back())
            allNumbers[i-1]=upperArr
        }
        score += allNumbers[0].Back()
    }
    fmt.Println(score)
}
