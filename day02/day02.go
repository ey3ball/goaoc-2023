package day02

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)


type Draw struct {
	red	int
	blue	int
	green	int
}

func scan(line string) Draw {
	regex, err := regexp.Compile("([0-9]+) (red|green|blue)")
	if err != nil {
		panic("Bad regexp")
	}

	matches := regex.FindAllStringSubmatch(line, -1)
	draw := Draw{}
	for _, m := range(matches) {
		n, err := strconv.Atoi(m[1])
		if err != nil {
			panic("Fuuu")
		}

		switch m[2] {
		case "red":
			draw.red = n
		case "green":
			draw.green = n
		case "blue":
			draw.blue = n
		}
	}
	return draw
}

func parse(scanner *bufio.Scanner) [][]Draw {
	input := make([][]Draw, 0)
	for scanner.Scan() {
		line := scanner.Text()

		_, line, ok := strings.Cut(line, ":")
		if !ok {
			panic("Token not found")
		}

		draw := make([]Draw, 0)
		for _, split := range(strings.Split(line, ";")) {
			draw = append(draw, scan(split))
		}
		input = append(input, draw)
		//fmt.Println(line)
		//fmt.Println(draw)
	}
	return input
}

func Part1(scanner *bufio.Scanner) {
	games := parse(scanner)
	loaded := Draw{
		red:   12,
		blue:  14,
		green: 13,
	}

	acc := 0
	for i, game := range(games) {
		possible := true
		for _, draw := range(game) {
			if draw.blue > loaded.blue ||
			   draw.red > loaded.red ||
			   draw.green > loaded.green {
				possible = false
			}
		}

		if possible {
			acc += (i + 1)
		}
	}

	fmt.Println(acc)
}

func Part2(scanner *bufio.Scanner) {
	games := parse(scanner)

	acc := int64(0)
	for _, game := range(games) {
		min := Draw{
			red:   0,
			blue:  0,
			green: 0,
		}


		for _, draw := range(game) {
			if draw.blue > min.blue {
				min.blue = draw.blue
			}
			if draw.red > min.red {
				min.red = draw.red
			}
			if draw.green > min.green {
				min.green = draw.green
			}
		}
		power := min.red * min.blue * min.green
		fmt.Println(power)
		acc += int64(power)
	}
	fmt.Println(acc)
}
