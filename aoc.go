package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func part1() {
	f, err := os.Open("./input/day01.txt")
	if err != nil {
		fmt.Println("Bye")
		return
	}

	scanner := bufio.NewScanner(f)

	keepdigits := func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		} else {
			return -1
		}
	}

	var acc int
	for scanner.Scan() {
		line := scanner.Text()

		digits := strings.Map(keepdigits, line)

		fmt.Println(">>" + line)

		first, _ := strconv.Atoi(digits[:1])
		last, _ := strconv.Atoi(digits[len(digits)-1:])

		calibration := first*10 + last
		fmt.Println(calibration)
		acc += calibration
	}

	fmt.Println(acc)
}

func main() {
	f, err := os.Open("./input/day01.txt")
	if err != nil {
		fmt.Println("Bye")
		return
	}

	scanner := bufio.NewScanner(f)

	keepdigits := func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		} else {
			return -1
		}
	}

	var acc int
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(">>" + line)
		//replacer := strings.NewReplacer(
		//	"one", "1", "two", "2", "three", "3", "four", "4",
		//	"five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9",
		//)
		//line = replacer.Replace(line)
		line = strings.ReplaceAll(line, "one", "o1e")
		line = strings.ReplaceAll(line, "two", "t2o")
		line = strings.ReplaceAll(line, "three", "t3e")
		line = strings.ReplaceAll(line, "four", "f4r")
		line = strings.ReplaceAll(line, "five", "f5e")
		line = strings.ReplaceAll(line, "six", "s6x")
		line = strings.ReplaceAll(line, "seven", "s7n")
		line = strings.ReplaceAll(line, "eight", "e8t")
		line = strings.ReplaceAll(line, "nine", "n9e")

		digits := strings.Map(keepdigits, line)

		fmt.Println(">>" + line)

		first, _ := strconv.Atoi(digits[:1])
		last, _ := strconv.Atoi(digits[len(digits)-1:])

		calibration := first*10 + last
		fmt.Println(calibration)
		acc += calibration
	}

	fmt.Println(acc)
}
