package day18b

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Pos struct {
	X int
	Y int
}

type Edge struct {
	Direction Pos
	Count int
	Color int
}

const Verbose = false
var Up = Pos{0, -1}
var Down = Pos{0, 1}
var Right = Pos{1, 0}
var Left = Pos{-1, 0}

var Symbols = [4]int{3, 1, 0, 2}
var Directions = [4]Pos{Up, Down, Right, Left}

var Direction map[int]Pos

func init() {
	Direction = make(map[int]Pos)
	for i, s := range(Symbols) {
		Direction[s] = Directions[i]
	}
}

func parse(scanner *bufio.Scanner) ([]Edge, []Pos) { 
	edges := make([]Edge, 0)
	for scanner.Scan() {
		line := scanner.Text()
		
		_, rest, _ := strings.Cut(line, " ")
		_, rest, _ = strings.Cut(rest, " ")
		color_ := rest[2:7]

		dir, _ := strconv.Atoi(rest[7:8])
		count, _ := strconv.ParseInt(color_, 16, 32)
		color := 0

		edges = append(edges, Edge{Direction[dir], int(count), int(color)})
	}

	X := 0
	Y := 0

	pos := make([]Pos, 0)

	for _, e := range(edges) {
		X += e.Direction.X * e.Count
		Y += e.Direction.Y * e.Count

		pos = append(pos, Pos{X+999999,Y+999999})
	}

	return edges, pos
}

//func Show(w World) {
//	for _, l := range(w.Map) {
//		fmt.Println(string(l))
//	}
//	fmt.Println()
//	fmt.Println("S:  ", w.Start)
//	fmt.Println("Xs: ", w.Xmax)
//	fmt.Println("Ys: ", w.Ymax)
//}

func Add(p1 Pos, p2 Pos) Pos {
	return Pos{p1.X + p2.X, p1.Y + p2.Y}
}

func Inv(p Pos) Pos {
	return Pos{-p.X, -p.Y}
}


//func Reachable(p Pos, w World) bool {
//	if p.X >= w.Xmax || p.Y >= w.Ymax {
//		return false
//	}
//
//	if p.X < 0 || p.Y < 0 {
//		return false
//	}
//
//	return true
//}
//
///*
// * Arriving into Pos by moving into direction with_heading
// * Can the cell be reached by entering the pipe ? -> bool
// * What is the exit heading through the pipe ? -> Pos
// */
//func InOutPos(w World, p Pos, with_heading Pos) (Pos, bool) {
//	symbol := w.Map[p.Y][p.X]
//	if symbol == '.' {
//		return Pos{}, false
//	}
//
//	into := Inv(with_heading)
//	//fmt.Println("Into:", into)
//	//fmt.Println("Symbol:", string(symbol))
//	//fmt.Println("Into:", Pipe[symbol])
//	if Pipe[symbol][0] == into {
//		return Pipe[symbol][1], true
//	} else if Pipe[symbol][1] == into {
//		return Pipe[symbol][0], true
//	}
//
//	return Pos{}, false
//}

//func maze(w World, start State) []State {
//	visited := []State{start}
//
//	for {
//		cur := visited[len(visited) - 1]
//
//		//fmt.Println(cur)
//		next_pos := Add(cur.Position, cur.Direction)
//		if next_pos == start.Position {
//			fmt.Println("#### Mazel tov ! ", len(visited))
//			return visited
//		}
//
//		if Verbose { fmt.Println("N:", next_pos) }
//
//		if !Reachable(next_pos, w) {
//			if Verbose { fmt.Println("NR:") }
//			break
//		}
//
//		next_heading, ok := InOutPos(w, next_pos, cur.Direction)
//		if !ok {
//			break
//		}
//		if Verbose { fmt.Println("H:", next_heading) }
//
//		visited = append(visited, State{next_heading, next_pos})
//	}
//
//	return nil
//}

func Shoelace(pos []Pos) int64 {
	area := int64(0)
	fmt.Println(pos)
	pos = append(pos[len(pos)-1:], pos[:]...)
	pos = append(pos, pos[1])
	fmt.Println(pos)

	for i := 0; i < len(pos) - 2; i++ {
		area += int64(pos[i].X * pos[i+1].Y - pos[i+1].X * pos[i].Y)
	}
	area /= 2
	return area
}

func Part1(scanner *bufio.Scanner) {
	edges, pos:= parse(scanner)
	fmt.Println("Edges: ", edges)
	fmt.Println("Pos: ", pos)

	area := Shoelace(pos)
	ext := int64(0)
	for _, e := range(edges) {
		ext += int64(e.Count)
	}
	fmt.Println(area)
	fmt.Println(ext)
	fmt.Println(1 + area + ext / 2)

}

func Part2(scanner *bufio.Scanner) {
	//world := parse(scanner)

	//fmt.Println("Edges: ", world.Edges)
	//fmt.Println("Max: ", world.XMax, world.YMax)
}
