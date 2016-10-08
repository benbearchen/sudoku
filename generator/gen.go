package generator

import (
	n2n "github.com/benbearchen/sudoku/nine2nine"

	"fmt"
	"math/rand"
	"time"
)

type genNode struct {
	key  n2n.Board
	s    *n2n.State
	left [][3]int
}

var (
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func newGenNode(key n2n.Board, s *n2n.State) *genNode {
	all := s.Choice().All()

	left := make([][3]int, len(all))
	for i, r := range rnd.Perm(len(left)) {
		left[i] = all[r]
	}

	return &genNode{key, s, left}
}

func (n *genNode) next() *genNode {
	for len(n.left) > 0 {
		tup := n.left[0]
		n.left = n.left[1:]

		number := tup[0]
		x, y := n2n.Cell2XY(tup[1], tup[2])
		s, ok := n.s.Next(number, x, y)
		if ok {
			key := n.key
			key.Set(x, y, number)
			return newGenNode(key, s)
		}
	}

	return nil
}

func SureGenerateComplete() (board n2n.Board, key n2n.Board) {
	for {
		b, key, _, ok := GenerateComplete()
		if ok {
			return b, key
		}
	}
}

func GenerateComplete() (board n2n.Board, key n2n.Board, times int, found bool) {
	c := 0
	maxDepth := 1
	defer func() {
		//fmt.Println("try times:", c, "maxDepth:", maxDepth)
	}()

	nodes := []*genNode{newGenNode(n2n.Board{}, n2n.NewState(n2n.Board{}))}
	for len(nodes) > 0 {
		n := nodes[len(nodes)-1].next()
		c++

		if n == nil {
			nodes = nodes[:len(nodes)-1]
			continue
		}

		if n.s.Board().Finished() {
			board = n.s.Board()
			key = n.key
			times = c
			found = true
			return
		}

		if c > 50 {
			break
		}

		nodes = append(nodes, n)
		if len(nodes) > maxDepth {
			maxDepth = len(nodes)
		}
	}

	times = c
	return
}

func Digg(b n2n.Board) n2n.Board {
	const min = 30
	for {
		seq := rnd.Perm(n2n.N * n2n.N)
		for i := min; i < 35; i++ {
			p := seq[i:]

			b1 := b
			for _, p := range p {
				b1[p] = 0
			}

			s := n2n.NewState(b1)
			if s.Trim2() {
				continue
			}

			if s.Board().Finished() {
				return b1
			} else {
				continue
			}

			fmt.Println(b1.OneLine())
			k := len(n2n.NewSearch(s).Search(5000, false))
			fmt.Println("find", k, "result(s)")
			if k == 1 {
				return b1
			}
		}
	}
}
