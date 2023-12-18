package day15

import (
	"bufio"
	"fmt"
	"strconv"
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

type Lens struct {
	label string
	value int
}

func Find(slice []Lens, label string) (int, bool) {
	for i, l := range(slice) {
		if l.label == label {
			return i, true
		}
	}
	return -1, false
}

func Part2(scanner *bufio.Scanner) {
	instrs := parse(scanner)
	state := make([][]Lens, 256)

	for _, i := range(instrs) {
		if i[len(i)-1:] == "-" {
			label := i[:len(i)-1]
			box := hash(label)

			i, found := Find(state[box], label)
			if found {
				state[box] = append(state[box][:i], state[box][i+1:]...)
			}
		} else {
			label, val_, ok := strings.Cut(i, "=")
			if !ok {
				panic("not ok")
			}
			val, _ := strconv.Atoi(val_)
			box := hash(label)
		
			i, found := Find(state[box], label)
			if found {
				state[box][i] = Lens{label, val}
			} else {
				state[box] = append(state[box], Lens{label, val})
			}
		}
	}

	power := 0
	for h, box := range(state) {
		for p, l := range(box) {
			val := (h+1) * (p+1) * l.value
			power += val
		}
	}
	fmt.Println(power)
}
