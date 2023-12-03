package day03

import (
	"bufio"
	"fmt"
	"strconv"
	"unicode"
)

type Pos struct {x,y int}

func newserial(serials *map[int][]Pos, num string, pos *[]Pos) (string, []Pos) {
	serial, _ := strconv.Atoi(num)
	(*serials)[serial] = *pos

	return "", make([]Pos, 0)
}

type Input struct {
	symbols map[Pos]rune
	serials map[int][]Pos
}

func parse(scanner *bufio.Scanner) Input {
	symbols := make(map[Pos]rune)
	serials := make(map[int][]Pos)

	l := 0
	for scanner.Scan() {
		line := scanner.Text()

		num := ""
		pos := make([]Pos, 0)
		for n, c := range(line) {
			if unicode.IsDigit(c) {
				num += string(c)
				pos = append(pos, Pos{l, n})
			} else if c != '.' {
				symbols[Pos{l, n}] = c
			}

			if !unicode.IsDigit(c) {
				num, pos = newserial(&serials, num, &pos)
			}
		}

		if num != "" {
			num, pos = newserial(&serials, num, &pos)
		}

		l += 1
	}
	return Input{symbols, serials}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func neigh(pos []Pos, pos2 Pos) bool {
	for _, pos1 := range(pos) {
		if abs(pos1.x - pos2.x) <= 1 && abs(pos1.y - pos2.y) <= 1 {
			return true
		}
	}
	return false
}

func Part1(scanner *bufio.Scanner) {
	in := parse(scanner)

	acc := 0
	for serial, pos := range(in.serials) {
		for symbol_pos, _ := range(in.symbols) {
			if neigh(pos, symbol_pos) {
				acc += serial
				break
			}
		}
	}
	fmt.Println(acc)
}

func Part2(scanner *bufio.Scanner) {

}
