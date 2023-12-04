package day04

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Card []int

func numbers(list_str string) []int {
	regex, err := regexp.Compile("([0-9]+)")
	if err != nil {
		panic("Bad regexp")
	}

	matches := regex.FindAllStringSubmatch(list_str, -1)

	nums := make([]int, 0)
	for _, m := range(matches) {
		n, err := strconv.Atoi(m[1])
		if err != nil {
			panic("Not a number")
		}

		nums = append(nums, n)
	}

	return nums
}

func parse(scanner *bufio.Scanner) ([][]int, []map[int]bool) {
	cards := make([][]int, 0)
	winners := make([]map[int]bool, 0)

	for scanner.Scan() {
		line := scanner.Text()

		_, card_string, _ := strings.Cut(line, ":")
		got_, win_, _ := strings.Cut(card_string, "|")
		
		got := numbers(got_)
		win := numbers(win_)
		winset := make(map[int]bool)

		for _, w := range(win) {
			winset[w] = true
		}

		cards = append(cards, got)
		winners = append(winners, winset)
	}

	return cards, winners
}

func score(got []int, winners map[int]bool) int64 {
	acc := int64(0)

	for _, n := range(got) {
		if winners[n] {
			if acc != 0 {
				acc *= 2
			} else {
				acc =1
			}
		}
	}
	return acc
}

func Part1(scanner *bufio.Scanner) {
	cards, winners := parse(scanner)

	acc := int64(0)
	for i, card := range(cards) {
		acc += score(card, winners[i])
	}

	fmt.Println(acc)
}

func Part2(scanner *bufio.Scanner) {
	parse(scanner)
}
