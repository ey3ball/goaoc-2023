package day09

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func parse(scanner *bufio.Scanner) [][]int {
	out := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()

		values := strings.Fields(line)
		l := make([]int, 0)
		for _, v := range(values) {
			i, _ := strconv.Atoi(v)
			l = append(l, i)
		}
		out = append(out, l)
	}

	return out
}

func extrapolate(values []int) int {
	deltas := make([]int, 0)
	zeros := true
	for i := range values[1:] {
		d := values[i+1] - values[i]
		if d != 0 {
			zeros = false
		}
		deltas = append(deltas, d)
	}

	if zeros {
		return 0
	} else {
		return deltas[len(deltas) - 1] + extrapolate(deltas)
	}
}

func extrapolate2(values []int) int {
	deltas := make([]int, 0)
	zeros := true
	for i := range values[1:] {
		d := values[i+1] - values[i]
		if d != 0 {
			zeros = false
		}
		deltas = append(deltas, d)
	}

	if zeros {
		return 0
	} else {
		return deltas[0] - extrapolate2(deltas)
	}
}


func Part1(scanner *bufio.Scanner) {
	values := parse(scanner)

	acc := 0
	for _, v := range(values) {
		acc += v[len(v) - 1] + extrapolate(v)
	}
	fmt.Println(acc)
}

func Part2(scanner *bufio.Scanner) {
	values := parse(scanner)

	acc := 0
	for _, v := range(values) {
		acc += v[0] - extrapolate2(v)
	}
	fmt.Println(acc)
}
