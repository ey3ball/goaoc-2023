package day08

import (
	"bufio"
	"fmt"
	"regexp"
	"sort"
)

type Node struct {
	Left string
	Right string
}


// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int64) int64 {
      for b != 0 {
              t := b
              b = a % b
              a = t
      }
      return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int64, integers ...int64) int64 {
      result := a * b / GCD(a, b)

      for i := 0; i < len(integers); i++ {
              result = LCM(result, integers[i])
      }

      return result
}

func parse(scanner *bufio.Scanner) (string, map[string]Node) {
	scanner.Scan()
	directions := scanner.Text()
	scanner.Scan()

	parse, _ := regexp.Compile("[A-Z0-9][A-Z0-9][A-Z0-9]")

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

type State struct {
	Inst int
	Node string
}

func find_period(start string, dirs string, net map[string]Node) (int, int, int, string) {
	i := 0
	pos := start

	seen := make(map[State]int)
	cycle := 0
	init := 0
	for cycle == 0 || pos[2] != 'Z' {
		//fmt.Println(pos, i % len(dirs))
		when, ok := seen[State{i % len(dirs), pos}] 

		if ok && cycle == 0 {
			cycle = i - when
			init = when
		}
		seen[State{i % len(dirs), pos}] = i

		dir := dirs[i % len(dirs)]
		if dir == 'R' {
			pos = net[pos].Right
		} else {
			pos = net[pos].Left
		}
		i++
	}

	return init, cycle, i % cycle, pos
}


func Part2(scanner *bufio.Scanner) {
	dirs, net := parse(scanner)

	start := make([]string, 0)
	for pos := range(net) {
		if pos[2] == 'A' {
			start = append(start, pos)
		}
	}

	sort.Strings(start)

	fmt.Println(start)

	acc := int64(1)
	for _, s := range(start) {
		fmt.Println(">> ", s)
		init, period, offset, _ := find_period(s, dirs, net)

		acc = LCM(int64(acc), int64(period))
		fmt.Println(acc)
		fmt.Println(init, period, offset)
	}
}

//func Part2(scanner *bufio.Scanner) {
//	dirs, net := parse(scanner)
//
//	visited := make(map[string]bool)
//	for pos, _ := range(net) {
//		if pos[2] == 'A' {
//			visited[pos] = true
//		}
//	}
//
//	fmt.Println(visited)
//
//	i := 0
//	for {
//		//fmt.Println(visited)
//		dir := dirs[i % len(dirs)]
//
//		next_visited := make(map[string]bool)
//
//		for pos := range(visited) {
//			if dir == 'R' {
//				next_visited[net[pos].Right] = true
//			} else {
//				next_visited[net[pos].Left] = true
//			}
//		}
//
//		all_z := true
//		for pos := range(next_visited) {
//			if pos[2] != 'Z' {
//				all_z = false
//				break
//			}
//		}
//
//		i++
//
//		if all_z {
//			break
//		}
//		visited = next_visited
//	}
//
//	fmt.Println(i)
//}
