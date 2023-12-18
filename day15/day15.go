package day15

import (
	"bufio"
	"fmt"
	"strings"
)

func parse(scanner *bufio.Scanner) []string {
	scanner.Scan()
	line := scanner.Text()

	instr := strings.Split(line, ",")

	return instr
}

func hash(in string) int {
	hash := 0
	for _, c := range(in) {
		hash += int(c)
		hash *= 17
		hash %= 256
	}
	return hash
}

func Part1(scanner *bufio.Scanner) {
	instrs := parse(scanner)

	fmt.Println(hash("HASH"))

	acc := 0
	for _, i := range(instrs) {
		h := hash(i)
		acc += h
	}
	fmt.Println(acc)
}

func Part2(scanner *bufio.Scanner) {
	parse(scanner)

	fmt.Println(0)
}
