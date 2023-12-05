package day05

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Range struct {
	dst int64
	src int64
	n int64
}

func map_seed(seed int64, ranges []Range) int64 {
	for _, r := range(ranges) {
		if seed >= r.src && seed < r.src + r.n {
			//fmt.Println("s:", seed)
			//fmt.Println(r.src)
			//fmt.Println(r.dst)
			return (seed - r.src + r.dst)
		}
	}
	return seed
}

func map_location(loc int64, ranges []Range) int64 {
	for _, r := range(ranges) {
		if loc >= r.dst && loc < r.dst + r.n {
			//fmt.Println("s:", seed)
			//fmt.Println(r.src)
			//fmt.Println(r.dst)
			return (loc - r.dst + r.src)
		}
	}
	return loc
}


func parse(scanner *bufio.Scanner) ([]int64, [][]Range) {
	scanner.Scan()
	_, seeds_str, _ := strings.Cut(scanner.Text(), ":")
	seeds := make([]int64, 0)
	for _, seed := range(strings.Fields(seeds_str)) {
		seed_value, _ := strconv.Atoi(seed)
		seeds = append(seeds, int64(seed_value))
	}
	fmt.Println(seeds)

	scanner.Scan()
	scanner.Text()

	mapper := make([][]Range, 0)
	for scanner.Scan() {
		area := scanner.Text()
		from, to, _ := strings.Cut(area[:len(area)-len(" map:")], "-to-")

		ranges := make([]Range, 0)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			range_str := strings.Fields(line)
			dst, _ := strconv.Atoi(range_str[0])
			src, _ := strconv.Atoi(range_str[1])
			n, _ := strconv.Atoi(range_str[2])
			ranges = append(ranges, Range{int64(dst), int64(src), int64(n)})
		}
		mapper = append(mapper, ranges)

		fmt.Println(from, "->", to)
		fmt.Println(ranges)
	}

	return seeds, mapper
}

func Part1(scanner *bufio.Scanner) {
	seeds, mappers := parse(scanner)
	mapped := make([]int64, 0)
	min := int64(math.MaxInt64)

	for _, seed := range(seeds) {
		from := seed
		to := seed

		for _, mapper := range(mappers) {
			to = map_seed(to, mapper)
		}
		mapped = append(mapped, to)

		if to < min {
			min = to
		}
		fmt.Println(from, " -> ", to)
	}

	fmt.Println(min)
}

func is_seed(maybe_seed int64, seeder []int64) bool {
	for i := 0; i < len(seeder); i+= 2 {
		if maybe_seed >= seeder[i] && maybe_seed < seeder[i] + seeder[i+1] {
			return true
		}
	}
	return false
}

func reseed(location int64, mappers[][]Range) int64 {
	locate := location

	for i := len(mappers) - 1; i >= 0; i-- {
		mapper := mappers[i]
		locate = map_location(locate, mapper)
	}

	return locate
}

func Part2(scanner *bufio.Scanner) {
	seeder, mappers := parse(scanner)

	for loc := 0; ; loc++ {
		seed := reseed(int64(loc), mappers)

		if is_seed(seed, seeder) {
			fmt.Println(seed, loc)
			break
		}
	}
}
