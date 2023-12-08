package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var cardToValue = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

const N_CARDS_IN_HAND = 5

func getStrength(hand string) HandType {
	counter := make(map[byte]int)
	for i := 0; i < N_CARDS_IN_HAND; i++ {
		counter[hand[i]]++
	}

	switch len(counter) {
	case 1:
		return FiveOfKind
	case 2:
		for _, count := range counter {
			if count == 4 {
				return FourOfKind
			}
		}
		return FullHouse
	case 3:
		for _, count := range counter {
			if count == 3 {
				return ThreeOfKind
			}
		}
		return TwoPair
	case 4:
		return OnePair
	default:
		return HighCard
	}
}

func getStrengthRec(hand string) int {
	newHand := strings.Split(hand, "")
	bestScore := 0
	for i := 0; i < N_CARDS_IN_HAND; i++ {
		if hand[i] == 'J' {
			for k := range cardToValue {
				if k != 'J' {
					newHand[i] = string(k)
					score := getStrengthRec(strings.Join(newHand, ""))
					bestScore = max(bestScore, score)
				}
			}
		}
	}
	return max(bestScore, int(getStrength(hand)))
}

func compareHands(hand1, hand2 string, part int) bool {
	var strength1, strength2 int
	if part == 1 {
		strength1, strength2 = int(getStrength(hand1)), int(getStrength(hand2))
	} else {
		strength1, strength2 = getStrengthRec(hand1), getStrengthRec(hand2)
		cardToValue['J'] = 0
	}

	if strength1 == strength2 {
		for i := 0; i < N_CARDS_IN_HAND; i++ {
			v1, v2 := cardToValue[hand1[i]], cardToValue[hand2[i]]
			if v1 != v2 {
				return v1 < v2
			}
		}
	}
	return strength1 < strength2
}

type handAndBid struct {
	hand string
	bid  int
}

func getHandValue(handAndBids []handAndBid) int {
	score := 0
	for i := 0; i < len(handAndBids); i++ {
		score += handAndBids[i].bid * (i + 1)
	}
	return score
}

func part1(handAndBids []handAndBid) int {
	sort.Slice(handAndBids, func(i, j int) bool {
		return compareHands(handAndBids[i].hand, handAndBids[j].hand, 1)
	})
	return getHandValue(handAndBids)
}

func part2(handAndBids []handAndBid) int {
	sort.Slice(handAndBids, func(i, j int) bool {
		return compareHands(handAndBids[i].hand, handAndBids[j].hand, 2)
	})
	return getHandValue(handAndBids)
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("error opening file", err)
		return
	}

	var handAndBids []handAndBid

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		hand := arr[0]
		bid, _ := strconv.Atoi(arr[1])
		handAndBids = append(handAndBids, handAndBid{hand: hand, bid: bid})
	}

	fmt.Println("Part #1:", part1(handAndBids))
	fmt.Println("Part #2:", part2(handAndBids))
}
