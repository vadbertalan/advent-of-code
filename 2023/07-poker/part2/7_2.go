// https://adventofcode.com/2023/day/7

package main

import (
	"aoc/utils"
	"cmp"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

const aocDay int = 7

const handLen = 5

type hand = string

// look into why not const
var cardPwr map[rune]int = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,

	'J': 0,
}

type handType = int

const (
	five handType = iota
	four
	fh
	three
	twoPair
	pair
	high
)

var typePwr map[handType]int = map[handType]int{
	five:    7,
	four:    6,
	fh:      5,
	three:   4,
	twoPair: 3,
	pair:    2,
	high:    1,
}

func CountEachRune(str string) (map[rune]int, []int) {
	runeMap := map[rune]int{}
	for _, c := range str {
		_, ok := runeMap[c]
		if ok {
			runeMap[c]++
		} else {
			runeMap[c] = 1
		}
	}
	occurrences := make([]int, 0, len(runeMap))
	for _, occ := range runeMap {
		occurrences = append(occurrences, occ)
	}
	return runeMap, occurrences
}

type cardCount = map[rune]int

var typeFuncs map[handType](func(hand hand) bool) = map[handType](func(hand hand) bool){
	five: func(hand hand) bool {
		cc, _ := CountEachRune(hand)
		return cc[rune(hand[0])] == 5
	},
	four: func(hand hand) bool {
		return strings.Count(string(hand), string(hand[0])) == 4 ||
			strings.Count(string(hand), string(hand[1])) == 4
	},
	fh: func(hand hand) bool {
		_, occurrences := CountEachRune(hand)
		sort.Ints(occurrences)
		return utils.EqualArr(occurrences, []int{2, 3})
	},
	three: func(hand hand) bool {
		_, occurrences := CountEachRune(hand)
		sort.Ints(occurrences)
		return utils.EqualArr(occurrences, []int{1, 1, 3})
	},
	twoPair: func(hand hand) bool {
		_, occurrences := CountEachRune(hand)
		sort.Ints(occurrences)
		return utils.EqualArr(occurrences, []int{1, 2, 2})
	},
	pair: func(hand hand) bool {
		_, occurrences := CountEachRune(hand)
		sort.Ints(occurrences)
		return utils.EqualArr(occurrences, []int{1, 1, 1, 2})
	},
	// could be omitted, as this should be the default
	high: func(hand hand) bool {
		cc, _ := CountEachRune(hand)
		return len(cc) == 5
	},
}

func evalJoker(hand hand) hand {
	if !strings.Contains(hand, "J") {
		return hand
	}

	if hand == "JJJJJ" {
		// could be "11111" too
		return "AAAAA"
	}

	maxHand := hand

	// In theory it would be enough to replace the Js with the first non 'J' char,
	// since the Js come into play only at the det. of the hand type
	// For example: JJ34J could be both 44344 or 33343.
	for _, card := range hand {
		if card != 'J' {
			tempHand := strings.ReplaceAll(hand, "J", string(card))
			if cmpHands(maxHand, tempHand) == -1 {
				maxHand = tempHand
			}
		}
	}

	return maxHand
}

func getHandType(hand hand) handType {
	for t, f := range typeFuncs {
		isOfType := f(hand)
		if isOfType {
			return t
		}
	}
	panic("should have found type, at least a `high`")
}

func cmpHandsByCards(hand1, hand2 hand) int {
	for i := 0; i < len(hand1); i++ {
		card1 := rune(hand1[i])
		card2 := rune(hand2[i])
		cmpResult := cmp.Compare(cardPwr[card1], cardPwr[card2])
		if cmpResult != 0 {
			return cmpResult
		}
	}
	// They are equal
	return 0
}

func cmpHands(hand1, hand2 string) int {
	handType1 := getHandType(hand1)
	handType2 := getHandType(hand2)

	cmpResult := cmp.Compare(typePwr[handType1], typePwr[handType2])
	if cmpResult != 0 {
		return cmpResult
	}

	return cmpHandsByCards(hand1, hand2)
}

func cmpHandsWithFixedHandTypes(hand1, hand2 string, handType1, handType2 handType) int {
	cmpResult := cmp.Compare(typePwr[handType1], typePwr[handType2])
	if cmpResult != 0 {
		return cmpResult
	}

	return cmpHandsByCards(hand1, hand2)
}

func cmpHandsWithJ(hand1, hand2 string) int {
	jokerEvaluatedHand1 := evalJoker(hand1)
	jokerEvaluatedHand2 := evalJoker(hand2)

	jokerEvaluatedHandType1 := getHandType(jokerEvaluatedHand1)
	jokerEvaluatedHandType2 := getHandType(jokerEvaluatedHand2)

	cmpResult := cmpHandsWithFixedHandTypes(hand1, hand2, jokerEvaluatedHandType1, jokerEvaluatedHandType2)
	if cmpResult != 0 {
		return cmpResult
	}

	return cmpHandsByCards(hand1, hand2)
}

type handScorePair struct {
	hand  hand
	score int
}

func second(lines []string) {
	fmt.Println("--- Second ---")

	result := 0

	handScorePairs := []handScorePair{}

	for _, line := range lines {
		fmt.Println(line)

		hand, scoreStr := utils.SplitIn2(line, " ")
		score, _ := strconv.Atoi(scoreStr)

		handScorePairs = append(handScorePairs, handScorePair{hand, score})
	}

	slices.SortFunc(handScorePairs, func(hsp1, hsp2 handScorePair) int {
		return cmpHandsWithJ(hsp1.hand, hsp2.hand)
	})

	for rank, hsp := range handScorePairs {
		result += (rank + 1) * hsp.score
	}

	fmt.Println("Result:", result)
}

// too high 252155519
// too high 252144968
// too high 252144282
// good 	252137472

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("../%d.%s", aocDay, inputFileExtension))

	second(lines)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
