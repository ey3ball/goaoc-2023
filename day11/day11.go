package day11

import (
	"bufio"
	"fmt"
)

type Pos struct {
	X int
	Y int
}

func parse(scanner *bufio.Scanner) []Pos {
	stars := make([]Pos, 0)

	// Keep track of lines/columns with objects
	xs := make(map[int]bool)
	ys := make(map[int]bool)
	
	y := 0
	for scanner.Scan() {
		line := scanner.Text()

		for x, c := range(line) {
			if c == '#' {
				stars = append(stars, Pos{x,y})
				xs[x] = true
				ys[y] = true
			}
		}
		y++
	}

	expanded_stars := make([]Pos, 0)
	for _, star := range(stars) {
		dx := star.X
		dy := star.Y

		for x, ok := range(xs) {
			if x < star.X && ok {
				dx--
			}
		}

		for y, ok := range(ys) {
			if y < star.Y && ok{
				dy--
			}
		}

		//expanded_stars = append(expanded_stars, Pos{star.X + (1000000-1)*dx, star.Y + (1000000-1)*dy})
		expanded_stars = append(expanded_stars, Pos{star.X + (1000000-1)*dx, star.Y + (1000000-1)*dy})
	}

	return expanded_stars
}

func abs(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}

func distance(p1 Pos, p2 Pos) int {
	dx := abs(p1.X - p2.X)
	dy := abs(p2.Y - p1.Y)

	return dx + dy
}

func Part1(scanner *bufio.Scanner) {
	stars := parse(scanner)

	pairs := 0
	sum := 0
	for i, p := range(stars) {
		for _, p2 := range(stars[i+1:]) {
			//fmt.Printf(
			//	"D %d -> %d : %d\n",
			//	(i+1), (j+i+2),
			//	distance(p, p2),
			//)
			sum += distance(p,p2)
			pairs++
		}
	}

	fmt.Println(sum)
	fmt.Println(pairs)
}

func Part2(scanner *bufio.Scanner) {
	parse(scanner)

	fmt.Println(0)
}
