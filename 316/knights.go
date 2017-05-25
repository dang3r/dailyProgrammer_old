package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type move struct {
	x int
	y int
}

// https://www.reddit.com/r/dailyprogrammer/comments/6coqwk/20170522_challenge_316_easy_knights_metric/
// Use BFS and don't revisit past nodes

func main() {
	var x, y int
	var val string
	moves := []move{
		move{-1, -2},
		move{1, -2},
		move{-1, 2},
		move{1, 2},
		move{-2, -1},
		move{2, -1},
		move{-2, 1},
		move{2, 1},
	}

	reader := bufio.NewReader(os.Stdin)
	val, _ = reader.ReadString(' ')
	x, _ = strconv.Atoi(val[:len(val)-1])
	val, _ = reader.ReadString('\n')
	y, _ = strconv.Atoi(val[:len(val)-1])

	goal := move{x, y}
	start := move{0, 0}

	steps := map[move]int{
		start: 0,
	}
	previousPos := map[move]move{}
	queue := []move{start}

	for {
		p := queue[0]
		if p == goal {
			back := p
			path := []move{}
			for {
				// WTF
				path = append(path, move{})
				copy(path[1:], path[0:])
				path[0] = back
				if back == start {
					break
				}
				back = previousPos[back]
			}

			fmt.Println(steps[p])
			for _, step := range path {
				fmt.Println(step)
			}
			os.Exit(0)
		}

		for _, m := range moves {
			newPos := move{
				p.x + m.x,
				p.y + m.y,
			}
			if _, ok := steps[newPos]; !ok {
				previousPos[newPos] = p
				steps[newPos] = steps[p] + 1
			}
			queue = append(queue, newPos)
		}
		queue = queue[1:]
	}

}
