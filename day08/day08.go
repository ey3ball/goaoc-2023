package day08

import (
	"bufio"
	"fmt"
	"regexp"
)

type Node struct {
	Left string
	Right string
}

func parse(scanner *bufio.Scanner) (string, map[string]Node) {
	scanner.Scan()
	directions := scanner.Text()
	scanner.Scan()

	parse, _ := regexp.Compile("[A-Z][A-Z][A-Z]")

	network := make(map[string]Node)
	for scanner.Scan() {
		line := scanner.Text()

		matches := parse.FindAllString(line, 3)
		network[matches[0]] = Node{matches[1], matches[2]}
	}

	return directions, network
}


func Part1(scanner *bufio.Scanner) {
	dirs, net := parse(scanner)

	pos := "AAA"
	i := 0
	for ; pos != "ZZZ"; i++ {
		dir := dirs[i % len(dirs)]

		if dir == 'R' {
			pos = net[pos].Right
		} else {
			pos = net[pos].Left
		}
	}

	fmt.Println(i)
}

func Part2(scanner *bufio.Scanner) {
	parse(scanner)

	fmt.Println(0)
}
