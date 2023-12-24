package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const NUMBER_OF_BOXES = 256

func getHashVal(seq string) int {
	ret := 0
	for _, c := range seq {
		ret = ((ret + int(c)) * 17) % NUMBER_OF_BOXES
	}
	return ret
}

func part1(sequences []string) int {
	total := 0
	for _, seq := range sequences {
		total += getHashVal(seq)
	}
	return total
}

type Lens struct {
	label       string
	focalLength int
}

func getLensLabelIdx(lenses []Lens, label string) int {
	for i, lens := range lenses {
		if lens.label == label {
			return i
		}
	}
	return -1
}

func remove(slice []Lens, idx int) []Lens {
	return append(slice[:idx], slice[idx+1:]...)
}

func part2(sequences []string) int {
	boxes := make([][]Lens, NUMBER_OF_BOXES)
	for _, seq := range sequences {
		if strings.Contains(seq, "=") {
			splitted := strings.Split(seq, "=")
			label, focalLengthStr := splitted[0], splitted[1]
			focalLength, _ := strconv.Atoi(focalLengthStr)
			boxNumber := getHashVal(label)
			newLens := Lens{label: label, focalLength: focalLength}
			if idx := getLensLabelIdx(boxes[boxNumber], label); idx != -1 {
				boxes[boxNumber][idx] = newLens
			} else {
				boxes[boxNumber] = append(boxes[boxNumber], newLens)
			}
		} else {
			label := strings.Split(seq, "-")[0]
			boxNumber := getHashVal(label)
			if idx := getLensLabelIdx(boxes[boxNumber], label); idx != -1 {
				boxes[boxNumber] = remove(boxes[boxNumber], idx)
			}
		}
	}

	total := 0
	for i, box := range boxes {
		for j, lens := range box {
			total += (i + 1) * (j + 1) * lens.focalLength
		}
	}
	return total
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		fmt.Printf("error while opening file %v\n", err)
		return
	}
	defer file.Close()

	var sequences []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sequences = strings.Split(scanner.Text(), ",")
	}

	fmt.Println("Part #1:", part1(sequences))
	fmt.Println("Part #2:", part2(sequences))
}
