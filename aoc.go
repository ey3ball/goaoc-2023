package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/ey3ball/goaoc-2023/day18b"
)

func main() {
	f, err := os.Open("./input/day18.txt")
	if err != nil {
		fmt.Println("Bye")
		return
	}

	scanner := bufio.NewScanner(f)

	fmt.Println("Part 1")
	day18b.Part1(scanner)

	f.Seek(0, 0)
	scanner = bufio.NewScanner(f)

	fmt.Println("Part 2")
	day18b.Part2(scanner)
}
