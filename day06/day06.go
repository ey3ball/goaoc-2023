package day06

import (
	"bufio"
	"fmt"
)

var TimesInput = [4]int{ 40, 70, 98, 79 }
var TimesSample = [3]int{ 7, 15, 30 }
var DistanceInput = [4]int{ 215, 1051, 2147, 1005 }
var DistanceSample = [3]int{ 9, 40, 200 }

func parse(scanner *bufio.Scanner) ([]int, []int) {
	return TimesInput[:], DistanceInput[:]
	//return TimesSample[:], DistanceSample[:]
}

func Part1(scanner *bufio.Scanner) {
	times, best := parse(scanner)

	acc := 1
	for i, time := range(times) {
		beat := best[i]

		beaten_cnt := 0
		for j := 0; j < time; j++ {
			dst := j*(time - j)

			fmt.Println("Beat? ", time, j, " - (dst,beat) ", dst, beat)
			if dst > beat {
				beaten_cnt++
			}
		}
		fmt.Println("race: ", i, beaten_cnt)
		acc *= beaten_cnt
	}

	fmt.Println(acc)
}

func Part2(scanner *bufio.Scanner) {
	parse(scanner)

	fmt.Println(0)
}
