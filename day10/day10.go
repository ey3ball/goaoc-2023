package day10

import (
	"bufio"
	"fmt"
)

type Pos struct {
	X int
	Y int
}

type World struct {
	Map [][]rune
	Xmax int
	Ymax int
	Start Pos
}

type State struct {
	Direction Pos
	Position Pos
}

const Verbose = false
var North = Pos{0, -1}
var South = Pos{0, 1}
var East = Pos{1, 0}
var West = Pos{-1, 0}

var Symbols = [6]rune{'|', '-', 'L', 'J', '7', 'F'}
var Connects = [6][2]Pos{
	{North, South},
	{East, West},
	{North, East},
	{North, West},
	{South, West},
	{South, East},
}

var Pipe map[rune][2]Pos

func init() {
	Pipe = make(map[rune][2]Pos)
	for i, s := range(Symbols) {
		Pipe[s] = Connects[i]
	}
}

func parse(scanner *bufio.Scanner) World {

	world := make([][]rune, 0)
	y := 0
	start := Pos{}
	for scanner.Scan() {
		line := scanner.Text()
		
		world = append(world, []rune(line))
		for x, c := range(line) {
			if c == 'S' {
				start = Pos{x,y}	
			}
		}
		y++
	}

	return World{world, len(world[0]), len(world), start}
}

func Show(w World) {
	for _, l := range(w.Map) {
		fmt.Println(string(l))
	}
	fmt.Println()
	fmt.Println("S:  ", w.Start)
	fmt.Println("Xs: ", w.Xmax)
	fmt.Println("Ys: ", w.Ymax)
}

func Add(p1 Pos, p2 Pos) Pos {
	return Pos{p1.X + p2.X, p1.Y + p2.Y}
}

func Inv(p Pos) Pos {
	return Pos{-p.X, -p.Y}
}


func Reachable(p Pos, w World) bool {
	if p.X > w.Xmax || p.Y > w.Ymax {
		return false
	}

	if p.X < 0 || p.Y < 0 {
		return false
	}

	return true
}

/* 
 * Arriving into Pos by moving into direction with_heading
 * Can the cell be reached by entering the pipe ? -> bool
 * What is the exit heading through the pipe ? -> Pos
 */
func InOutPos(w World, p Pos, with_heading Pos) (Pos, bool) {
	symbol := w.Map[p.Y][p.X]
	if symbol == '.' {
		return Pos{}, false
	}

	into := Inv(with_heading)
	//fmt.Println("Into:", into)
	//fmt.Println("Symbol:", string(symbol))
	//fmt.Println("Into:", Pipe[symbol])
	if Pipe[symbol][0] == into {
		return Pipe[symbol][1], true
	} else if Pipe[symbol][1] == into {
		return Pipe[symbol][0], true
	}

	return Pos{}, false
}

func maze(w World, start State) int {
	visited := []State{start}

	for {
		cur := visited[len(visited) - 1]

		fmt.Println(cur)
		next_pos := Add(cur.Position, cur.Direction)
		if next_pos == start.Position {
			fmt.Println("#### Mazel tov ! ", len(visited))
			return len(visited) / 2
		}

		if Verbose { fmt.Println("N:", next_pos) }

		if !Reachable(next_pos, w) {
			if Verbose { fmt.Println("NR:") }
			break
		}

		next_heading, ok := InOutPos(w, next_pos, cur.Direction)
		if !ok {
			break
		}
		if Verbose { fmt.Println("H:", next_heading) }

		visited = append(visited, State{next_heading, next_pos})
	}

	return -1
}

func Part1(scanner *bufio.Scanner) {
	world := parse(scanner)

	start_states := []State{
		{North, world.Start},
		{South, world.Start},
		{East, world.Start},
		{West, world.Start},
	}

	steps := -1
	for _, s := range(start_states) {
		fmt.Println("", s)
		fmt.Println(">> Enter maze from : ", s)
		steps = maze(world, s)

		if steps != -1 {
			break
		}
	}

	Show(world)
	fmt.Println("Steps: ", steps)
}

func Part2(scanner *bufio.Scanner) {
	parse(scanner)

	fmt.Println(0)
}
