package day07

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	Cards []string
	Bid int
}

//func (h Hand) String() string {
//	str := ">"
//	for _, r := range(h.Cards) {
//		str += string(r)
//	}
//
//	str += "_" + strconv.Itoa(h.Bid)
//	return str
//}


var labels = [13]string{
	"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2",
}

var heads = make(map[string]int)

func init() {
	for i, l := range(labels) {
		heads[l] = len(labels) - i - 1
	}
}

func stats(hand []string) map[int][]string {
	counter := make(map[string]int)

	for _, r := range(hand) {
		counter[r] += 1
	}

	freqs := make(map[int][]string)
	for card, freq := range(counter) {
		freqs[freq] = append(freqs[freq], card)
	}

	return freqs
}

func strength(hand []string) int {
	freqs := stats(hand)

	if len(freqs[5]) == 1 {
		return 6
	} else if len(freqs[4]) == 1 {
		return 5
	} else if len(freqs[3]) == 1 && len(freqs[2]) == 1 {
		return 4
	} else if len(freqs[3]) == 1 {
		return 3
	} else if len(freqs[2]) == 2 {
		return 2
	} else if len(freqs[2]) == 1 {
		return 1
	} else {
		return 0
	}
}

func less(hand1 []string, hand2 []string) bool {
	for i := range(hand1) {
		if heads[hand1[i]] < heads[hand2[i]] {
			return true
		}
		if heads[hand1[i]] > heads[hand2[i]] {
			return false
		}
	}
	return true
}

type HandSlice []Hand

func (p HandSlice) Len() int {
	return len(p)
}

func (p HandSlice) Less(i, j int) bool {
	hand1 := p[i]
	hand2 := p[j]

	if strength(hand1.Cards) == strength(hand2.Cards) {
		return less(hand1.Cards, hand2.Cards)
	} else {
		return strength(hand1.Cards) <= strength(hand2.Cards)
	}
}

func (p HandSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]	
}


func parse(scanner *bufio.Scanner) []Hand {
	hands := make([]Hand, 0)
	
	fmt.Println(heads)

	for scanner.Scan() {
		line := scanner.Text()
		
		field := strings.Fields(line)
		bid, _ := strconv.Atoi(field[1])
		
		hand := make([]string, 0)
		for i := range(field[0]) {
			hand = append(hand, field[0][i:i+1])
		}


		hands = append(hands, Hand{hand, bid})
	}


	return hands
	//return TimesSample[:], DistanceSample[:]
}


func Part1(scanner *bufio.Scanner) {
	hands := parse(scanner)

	sort.Sort(HandSlice(hands))

	acc := 0
	for i, hand := range(hands) {
		acc += (i + 1) * hand.Bid
	}

	fmt.Println(acc)
}

func Part2(scanner *bufio.Scanner) {
	parse(scanner)

	fmt.Println(0)
}
