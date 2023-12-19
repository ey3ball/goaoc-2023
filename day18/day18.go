package day18

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

type World struct {
	Edges []Edge
	XMax int
	YMax int
	XMin int
	YMin int
}

const Verbose = false
var Up = Pos{0, -1}
var Down = Pos{0, 1}
var Right = Pos{1, 0}
var Left = Pos{-1, 0}

var Symbols = [4]string{"U", "D", "R", "L"}
var Directions = [4]Pos{Up, Down, Right, Left}

var Direction map[string]Pos

func init() {
	Direction = make(map[string]Pos)
	for i, s := range(Symbols) {
		Direction[s] = Directions[i]
	}
}

func parse(scanner *bufio.Scanner) World { 
	edges := make([]Edge, 0)
	for scanner.Scan() {
		line := scanner.Text()
		
		dir, rest, _ := strings.Cut(line, " ")
		count_, rest, _ := strings.Cut(rest, " ")
		color_ := rest[2:8]

		count, _ := strconv.Atoi(count_)
		color, _ := strconv.ParseInt(color_, 16, 32)

		edges = append(edges, Edge{Direction[dir], count, int(color)})
	}

	X := 0
	XMax := 0
	XMin := 0
	Y := 0
	YMax := 0
	YMin := 0

	pos := make([]Pos, 0)

	for _, e := range(edges) {
		X += e.Direction.X * e.Count
		Y += e.Direction.Y * e.Count

		if X > XMax {
			XMax = X
		}
		if Y > YMax {
			YMax = Y
		}
		if X < XMin {
			XMin = X
		}
		if Y < YMin {
			YMin = Y
		}

	}

	return World{edges, XMax, YMax, XMin, YMin}
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

func Draw(w World) [][]int {
	m := make([][]int, w.YMax-w.YMin+1)
	for i := range(m) {
		m[i] = make([]int, w.XMax-w.XMin+1)
	}

	p := Pos{-w.XMin,-w.YMin}
	for _, e := range(w.Edges) {
		for i := 0; i < e.Count; i++ {
			p = Add(p, e.Direction)
			m[p.Y][p.X] = e.Color
		}
	}
	return m
}

func Points(w World) [][]int {
	m := make([][]int, w.YMax-w.YMin+1)
	for i := range(m) {
		m[i] = make([]int, w.XMax-w.XMin+1)
	}

	p := Pos{-w.XMin,-w.YMin}
	for _, e := range(w.Edges) {
		for i := 0; i < e.Count; i++ {
			p = Add(p, e.Direction)
			m[p.Y][p.X] = e.Color
		}
	}
	return m
}


func Show(m [][]int) int {
	cnt := 0
	for _, l := range(m) {
		for _, c := range(l) {
			if c != 0 {
				fmt.Printf("#")
				cnt += 1
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println("")
	}
	return cnt
}

func Next(p Pos) []Pos {
	return []Pos{
		Add(p, Up),
		Add(p, Down),
		Add(p, Left),
		Add(p, Right),
	}
}

func Scan(w World, m[][]int) int {
	start := Pos{len(m[0])/2, len(m)/2}
	m[start.Y][start.X] = -1

	seen := make(map[Pos]bool)
	seen[start] = true

	next := Next(start)
	for len(next) != 0 {
		pop, next_ := next[len(next)-1], next[:len(next)-1]
		if seen[pop] {
			next = next_
			continue
		} else {
			if m[pop.Y][pop.X] == 0 {
				m[pop.Y][pop.X] = -1
				next_ = append(next_, Next(pop)...)
			}
			seen[pop] = true
		}
		next = next_

	}

	fmt.Println()
	cnt := Show(m)
	return cnt
}


func Part1(scanner *bufio.Scanner) {
	world := parse(scanner)
	fmt.Println("Edges: ", world.Edges)
	fmt.Println("Max: ", world.XMax, world.YMax)
	fmt.Println("Min: ", world.XMin, world.YMin)

	m := Draw(world)
	Show(m)
	cnt := Scan(world, m)
	fmt.Println("Cnt: ", cnt)

}

func Part2(scanner *bufio.Scanner) {
	//world := parse(scanner)

	//fmt.Println("Edges: ", world.Edges)
	//fmt.Println("Max: ", world.XMax, world.YMax)
}
