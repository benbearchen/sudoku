package nine2nine

import (
	"fmt"
	"strings"
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
	numbers := s.choice.allByNumbers()

	return &node{s, numbers}
}

func (n *node) next() (*State, bool) {
	for len(n.numbers) > 0 {
		number := n.numbers[0][0]
		c := n.numbers[0][1]
		p := n.numbers[0][2]
		n.numbers = n.numbers[1:]

		s, ok := n.s.Next(number, c, p)
		if ok {
			return s, true
		}
	}

	return nil, false
}

func (s *Search) Search() []Board {
	nodes := []*node{newNode(s.s)}
	results := make([]Board, 0)
	c := 0
	d := 0
	for len(nodes) > 0 {
		s, ok := nodes[len(nodes)-1].next()
		if !ok {
			nodes = nodes[:len(nodes)-1]
			continue
		}

		if s.Board().Finished() {
			b := s.Board()
			found := false
			for _, board := range results {
				if board == b {
					found = true
				}
			}

			if !found {
				//fmt.Println("one result:")
				//s.Board().Print()
				results = append(results, s.Board())
			}

			break
		} else {
			nodes = append(nodes, newNode(s))

			c++
			show := func() bool {
				if c%10000 == 0 {
					return true
				}

				dd := d - len(nodes)
				if dd <= -3 || dd >= 3 {
					return true
				}

				return false
			}

			if show() {
				d = len(nodes)
				fmt.Printf("\rdepath %2d times %-10d numbers %2d %s%s", len(nodes), c, s.Board().Numbers(), strings.Repeat("=", len(nodes)), strings.Repeat(" ", 30-len(nodes)))
			}
		}
	}

	fmt.Println()
	return results
}
