package day13

import (
	"bufio"
	"fmt"
	"slices"
)

func transpose(pattern [][]rune) [][]rune {
	transposed := make([][]rune, 0)

	for i := 0; i < len(pattern[0]); i++ {
		col := make([]rune, 0)

		for _, l := range(pattern) {
			col = append(col, l[i])
		}

		transposed = append(transposed, col)
	}
	return transposed
}

func parse(scanner *bufio.Scanner) [][][]rune {
	h_patterns := make([][][]rune, 0)

	pattern := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			h_patterns = append(h_patterns, pattern)
			pattern = make([][]rune, 0)
			continue
		}

		pattern = append(pattern, []rune(line))
	}
	h_patterns = append(h_patterns, pattern)

	return h_patterns
}

func Show(ps [][][]rune) {
	for i, p := range(ps) {
		fmt.Println(">> ", i)
		fmt.Println("")

		for _, l := range(p) {
			fmt.Println(string(l))
		}
		fmt.Println("")
	}
}

func Show2(ps [][][]rune) {
	for i, p := range(ps) {
		p = transpose(p)
		fmt.Println(">> ", i)
		fmt.Println("")

		for _, l := range(p) {
			fmt.Println(string(l))
		}
		fmt.Println("")
	}
}


func Reflections(ps [][][]rune, mirror bool) int {
	acc := 0
	r := 0
	for _, p := range(ps) {
		if mirror {
			p = transpose(p)
		}

		width := len(p) 
		//fmt.Println("go w:", width)

		for line := 1; line < width; line++ {
			size := min(width - line, line)
			all_ok := true

			for j := 0; j < size; j++ {
				if !slices.Equal(p[line - size + j], p[line + size - j - 1]) {
					all_ok = false
				}
			}

			if all_ok {
				acc += line
				r++
				fmt.Println("Reflected at ", line, size)
				//break
			}
		}
	}
	fmt.Println("#", r)
	return acc
}

func Part1(scanner *bufio.Scanner) {
	ps := parse(scanner)

	acc := Reflections(ps, true)
	acc += 100*Reflections(ps, false)
	
	fmt.Println(acc)
}

func Part2(scanner *bufio.Scanner) {
	parse(scanner)

	fmt.Println(0)
}
