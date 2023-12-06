package day06

import (
	"bufio"
	"fmt"
)

var TimesInput = [4]int{ 40, 70, 98, 79 }
var TimesSample = [3]int{ 7, 15, 30 }
var DistanceInput = [4]int{ 215, 1051, 2147, 1005 }
var DistanceSample = [3]int{ 9, 40, 200 }

var RealTimesInput = 40709879
var RealTimesSample = 71530
var RealDistanceInput = 215105121471005
var RealDistanceSample = 940200


func parse(scanner *bufio.Scanner) ([]int, []int) {
	return TimesInput[:], DistanceInput[:]
	//return TimesSample[:], DistanceSample[:]
}

func parse2(scanner *bufio.Scanner) (int, int) {
	return RealTimesInput, RealDistanceInput
	//return RealTimesSample, RealDistanceSample
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

func beats(try_time int, total_time int, beat int) bool {
	dst := try_time*(total_time - try_time)

	if dst > beat {
		return true
	} else {
		return false
	}
}

func Part2(scanner *bufio.Scanner) {
	time, beat := parse2(scanner)

	start := 0
	end := time / 2

	for (end - start) >= 2 {
		next_try := (start + end) / 2

		if beats(next_try, time, beat) {
			end = next_try
		} else {
			start = next_try
		}
		//fmt.Println(start, end)
	}
	fmt.Println("Lower:", end)
	lower := end

	start = time / 2
	end = time

	for (end - start) >= 2 {
		next_try := (start + end) / 2

		if !beats(next_try, time, beat) {
			end = next_try
		} else {
			start = next_try
		}
		fmt.Println(start, end)
	}
	upper := start
	fmt.Println((upper - lower) + 1)
}
