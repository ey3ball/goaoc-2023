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
	if p.X >= w.Xmax || p.Y >= w.Ymax {
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

func maze(w World, start State) []State {
	visited := []State{start}

	for {
		cur := visited[len(visited) - 1]

		//fmt.Println(cur)
		next_pos := Add(cur.Position, cur.Direction)
		if next_pos == start.Position {
			fmt.Println("#### Mazel tov ! ", len(visited))
			return visited
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

	return nil
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
		visited := maze(world, s)

		if visited != nil {
			steps = len(visited) / 2
			break
		}
	}

	//Show(world)
	fmt.Println("Steps: ", steps)
}

func Fill(w World) {
	visited := make(map[Pos]bool);
	edge := []Pos{{0,0},{0,w.Ymax-1},{w.Xmax-1,0},{w.Xmax-1,w.Ymax-1}}

	filled := 0
	for len(edge) != 0 {
		new_edge := make([]Pos, 0)
		for _, p := range(edge) {
			for _, d := range([]Pos{North, South, East, West}) {
				next_p := Add(p, d)
				v := false
				if next_p.X == 3 && next_p.Y == 5 {
					v = true	
				}
				if v {
					fmt.Println("Visiting: ", next_p)
				}
				if !Reachable(next_p, w) {
					continue
				}
				if visited[next_p] {
					continue
				}

				direction := Add(next_p, Inv(p))
				if v { fmt.Println("Dir: ", direction) }

				if w.Map[next_p.Y][next_p.X] == '.' {
					w.Map[next_p.Y][next_p.X] = 'O'
					filled += 1
				} else {
					inout, ok := InOutPos(w, next_p, direction)
					if v {
						fmt.Println("InOut: ", inout)
						fmt.Println("InOut: ", string(w.Map[next_p.Y][next_p.X]))
					}
					if !ok {
						// Cannot progress along pipe
						continue
					}
				}

				visited[next_p] = true
				new_edge = append(new_edge, next_p)
			}
		}
		edge = new_edge
	}
	Show(w)
	fmt.Println(filled)
}

func Part2(scanner *bufio.Scanner) {
	world := parse(scanner)

	visited := maze(world, State{North, world.Start})
	//visited = maze(world, State{South, world.Start})
	maze := make(map[Pos]bool)
	for _, v := range(visited) {
		maze[v.Position] = true

	}

	new_map := make([][]rune, 0)
	for y, l := range(world.Map) {
		new_line := make([]rune, 0)
		for x, c := range(l) {
			if !maze[Pos{x,y}] {
				c = '.'
			}
			new_line = append(new_line, c)
		}
		new_map = append(new_map, new_line)
	}

	new_world := World{new_map, world.Xmax, world.Ymax, world.Start}
	fmt.Println("Steps: ", len(visited)/2)

	Show(world)
	Fill(new_world)
	acc := 0
	for _, l := range(new_world.Map) {
		for _, c := range(l) {
			if c == '.' {
				acc += 1
			}
		}
	}
	fmt.Println("Filled: ", acc)
}
