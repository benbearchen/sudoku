package nine2nine

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Search struct {
	s *State
}

func NewSearch(s *State) *Search {
	return &Search{s}
}

type node struct {
	s       *State
	numbers [][3]int
}

func newNode(s *State) *node {
	numbers := s.choice.AllByNumbers()

	return &node{s, numbers}
}

func (n *node) next() (*node, bool) {
	for len(n.numbers) > 0 {
		tup := n.numbers[0]
		n.numbers = n.numbers[1:]

		number := tup[0]
		x, y := Cell2XY(tup[1], tup[2])
		s, ok := n.s.Next(number, x, y)

		//n.s.choice.kill(number, x, y)

		if ok {
			return newNode(s), true
		}
	}

	return nil, false
}

func (s *Search) Search(maxTimes int, single bool) []Board {
	nodes := []*node{newNode(s.s)}
	results := make([]Board, 0)
	c := 0
	d := 0
	maxDepth := len(nodes)
	start := time.Now()
	for len(nodes) > 0 {
		n, ok := nodes[len(nodes)-1].next()
		if !ok {
			nodes = nodes[:len(nodes)-1]
			continue
		}

		c++
		if n.s.Board().Finished() {
			b := n.s.Board()
			found := false
			for _, board := range results {
				if board == b {
					found = true
				}
			}

			if !found {
				//fmt.Println("one result:")
				//n.s.Board().Print()
				results = append(results, n.s.Board())
			}

			if single {
				break
			}
		} else {
			show := false
			nodes = append(nodes, n)
			if len(nodes) > maxDepth {
				maxDepth = len(nodes)
				show = true
			}

			if c > maxTimes {
				break
			}

			dd := d - len(nodes)
			if dd <= -3 || dd >= 3 {
				show = true
			}

			if show {
				d = len(nodes)
				fmt.Printf("\rdepth %2d times %-10d numbers %2d %s%s", len(nodes), c, n.s.Board().Numbers(), strings.Repeat("=", len(nodes)), strings.Repeat(" ", 30-len(nodes)))
			}
		}
	}

	result := "has result"
	if len(results) == 0 {
		if len(nodes) > 0 {
			result = "give up..."
		} else {
			result = "no result"
		}
	} else if len(results) > 1 {
		result = "has " + strconv.Itoa(len(results)) + " results"
	}

	fmt.Printf("\rmax depth %2d, times %d, used time %v, %s  \n", maxDepth, c, time.Now().Sub(start), result)
	return results
}
