package nine2nine

import (
	"fmt"
	"sort"
	"strconv"
)

const M = 3
const N = M * M

func XY2Cell(x, y int) (cell, pos int) {
	return (x/M + y/M*M), (x%M + y%M*M)
}

func Cell2XY(cell, pos int) (x, y int) {
	return (cell%M*M + pos%M), (cell/M*M + pos/M)
}

type Board [N * N]int

func (b Board) Numbers() int {
	c := 0
	for _, n := range b {
		if n != 0 {
			c++
		}
	}

	return c
}

func (b Board) Finished() bool {
	for _, v := range b {
		if v == 0 {
			return false
		}
	}

	return true
}

func (b Board) Validate() []error {
	errors := make([]error, 0)

	c := [N][N]int{}
	for x := 0; x < N; x++ {
		for y := 0; y < N; y++ {
			c[x][y] = -1
		}
	}

	h := c
	v := c

	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			n := b.Value(x, y)
			if n == 0 {
				continue
			} else if n < 0 || n > N {
				errors = append(errors, fmt.Errorf("invalid number %d, xy (%d, %d)", n, x, y))
				continue
			}

			mc, mp := XY2Cell(x, y)
			pv := c[mc][n-1]
			if pv == -1 {
				c[mc][n-1] = mp
			} else {
				ex, ey := Cell2XY(mc, pv)
				errors = append(errors, fmt.Errorf("multi %d in Cell[%d], xy (%d, %d) vs (%d, %d)", n, mc, ex, ey, x, y))
			}

			pv = h[y][n-1]
			if pv == -1 {
				h[y][n-1] = x
			} else {
				errors = append(errors, fmt.Errorf("multi %d, xy (%d, %d) vs (%d, %d)", n, pv, y, x, y))
			}

			pv = v[x][n-1]
			if pv == -1 {
				v[x][n-1] = y
			} else {
				errors = append(errors, fmt.Errorf("multi %d, xy (%d, %d) vs (%d, %d)", n, x, pv, x, y))
			}
		}
	}

	if len(errors) > 0 {
		return errors
	} else {
		return nil
	}
}

func (b Board) Value(x, y int) int {
	return b[y*N+x]
}

func (b Board) Empty(x, y int) bool {
	return b.Value(x, y) == 0
}

func (b *Board) Set(x, y, number int) {
	if b.Value(x, y) != 0 {
		fmt.Println("set ", x, y, number)
		fmt.Println("duplicate: ")
		b.Print()
		panic(fmt.Sprintf("duplicate set"))
	}

	b[y*N+x] = number
}

func (b Board) Print() {
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			n := b.Value(x, y)
			if n == 0 {
				fmt.Printf("- ")
			} else {
				fmt.Printf("%d ", n)
			}
		}

		fmt.Println()
	}
}

func (b Board) OneLine() string {
	r := ""
	for _, n := range b {
		r += strconv.Itoa(n)
	}

	return r
}

type group struct {
	num []int
	pos []int
}

type number struct {
	c []int
	h []int
	v []int
}

// 分裂各组、数字的剩余位置状态
func (b Board) SplitLeft() (cell, horizon, vertical [N]group, numbers [N]number) {
	cn := [N]map[int]bool{}
	cp := [N]map[int]bool{}
	hn := [N]map[int]bool{}
	hp := [N]map[int]bool{}
	vn := [N]map[int]bool{}
	vp := [N]map[int]bool{}

	for i := 0; i < N; i++ {
		cn[i] = make(map[int]bool)
		cp[i] = make(map[int]bool)
		hn[i] = make(map[int]bool)
		hp[i] = make(map[int]bool)
		vn[i] = make(map[int]bool)
		vp[i] = make(map[int]bool)

		for j := 0; j < N; j++ {
			cn[i][j+1] = true
			cp[i][j] = true
			hn[i][j+1] = true
			hp[i][j] = true
			vn[i][j+1] = true
			vp[i][j] = true
		}
	}

	ns := [N][3][N]bool{}

	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			n := b.Value(x, y)
			if n == 0 {
				continue
			}

			mc, mp := XY2Cell(x, y)
			delete(cn[mc], n)
			delete(cp[mc], mp)

			delete(hn[y], n)
			delete(hp[y], x)

			delete(vn[x], n)
			delete(vp[x], y)

			p := n - 1
			ns[p][0][mc] = true
			ns[p][1][x] = true
			ns[p][2][y] = true
		}
	}

	np := func(n, p [N]map[int]bool) (r [N]group) {
		for i := 0; i < N; i++ {
			num := make([]int, 0, N)
			for n := range n[i] {
				num = append(num, n)
			}

			pos := make([]int, 0, N)
			for p := range p[i] {
				pos = append(pos, p)
			}

			sort.Ints(num)
			sort.Ints(pos)

			r[i] = group{num, pos}
		}

		return
	}

	cell = np(cn, cp)
	horizon = np(hn, hp)
	vertical = np(vn, vp)

	dn := func() (r [N]number) {
		for i := 0; i < N; i++ {
			s := [3][]int{}
			for j := range ns[i] {
				pos := make([]int, 0, N)
				for p, ok := range ns[i][j] {
					if !ok {
						pos = append(pos, p)
					}
				}

				sort.Ints(pos)

				s[j] = pos
			}

			r[i] = number{s[0], s[1], s[2]}
		}

		return
	}

	numbers = dn()
	return
}

func (b Board) Choice() *Choice {
	c := NewChoice()
	c.Init(b)
	return c
}

type Choice [N][3][N][]int

func NewChoice() *Choice {
	a := func() []int {
		a := make([]int, N)
		for i := 0; i < N; i++ {
			a[i] = i
		}

		return a
	}

	c := new(Choice)
	for i := 0; i < N; i++ {
		for t := 0; t < 3; t++ {
			for j := 0; j < N; j++ {
				c[i][t][j] = a()
			}
		}
	}

	return c
}

func (c *Choice) Init(b Board) {
	for x := 0; x < N; x++ {
		for y := 0; y < N; y++ {
			n := b.Value(x, y)
			if n != 0 {
				c.trySet(x, y, n)
			}
		}
	}
}

func (c *Choice) can(x, y, number int) bool {
	mc, mp := XY2Cell(x, y)

	if !c.has(number, 0, mc, mp) || !c.has(number, 1, y, x) || !c.has(number, 2, x, y) {
		return false
	} else {
		return true
	}
}

func (c *Choice) trySet(x, y, number int) bool {
	mc, mp := XY2Cell(x, y)

	if !c.has(number, 0, mc, mp) || !c.has(number, 1, y, x) || !c.has(number, 2, x, y) {
		return false
	}

	//fmt.Printf("try set %d @ (%d, %d), <%d, %d>\n", number, x, y, mc, mp)
	c.freeGroup(number, 0, mc)
	c.freeGroup(number, 1, y)
	c.freeGroup(number, 2, x)

	for n := 1; n <= N; n++ {
		c.kill(n, x, y)
	}

	for i := 0; i < N; i++ {
		c.kill(number, x, i)
		c.kill(number, i, y)
	}

	mx0, my0 := x/M*M, y/M*M
	for i := 0; i < M; i++ {
		for j := 0; j < M; j++ {
			xc := mx0 + i
			yc := my0 + j
			c.kill(number, xc, yc)
		}
	}

	c.checkGroupPair(0, mc)
	for i := 0; i < M; i++ {
		c.checkGroupPair(1, my0+i)
		c.checkGroupPair(2, mx0+i)
	}

	return true
}

func (c *Choice) kill(n, x, y int) {
	nc, np := XY2Cell(x, y)
	c.free(n, 0, nc, np)
	c.free(n, 1, y, x)
	c.free(n, 2, x, y)
}

func (c *Choice) free(n, t, a, b int) {
	s := c[n-1][t][a]
	i := 0
	for ; i < len(s); i++ {
		if s[i] == b {
			break
		}
	}

	if i != len(s) {
		copy(s[i:], s[i+1:])
		s = s[:len(s)-1]
		c[n-1][t][a] = s

		if t == 0 {
			// 处理 Cell 内变成单行或者单列的情况
			h := -1
			v := -1
			switch len(s) {
			case 2:
				if s[0]/M == s[1]/M {
					h = s[0] / M
				} else if s[0]%M == s[1]%M {
					v = s[0] % M
				}
			case 3:
				if s[0]/M == s[1]/M && s[0]/M == s[2]/M {
					h = s[0] / M
				} else if s[0]%M == s[1]%M && s[0]%M == s[2]%M {
					v = s[0] % M
				}
			}

			if h != -1 {
				//fmt.Println("cell h:", s, "/", n, "<", a, ",", b, ">")
				for i := 0; i < M; i++ {
					nc := a/M*M + i
					if nc == a {
						continue
					}

					for j := 0; j < M; j++ {
						np := h*M + j
						nx, ny := Cell2XY(nc, np)
						c.kill(n, nx, ny)
						//fmt.Println("kill xy (", nx, ",", ny, ")")
					}
				}
			} else if v != -1 {
				//fmt.Println("cell v:", s, "/", n, "<", a, ",", b, ">")
				for i := 0; i < M; i++ {
					nc := i*M + a%M
					if nc == a {
						continue
					}

					for j := 0; j < M; j++ {
						np := j*M + v
						nx, ny := Cell2XY(nc, np)
						c.kill(n, nx, ny)
						//fmt.Println("kill xy (", nx, ",", ny, ")")
					}
				}
			}
		} else if t == 1 {
			h := -1
			switch len(s) {
			case 2:
				if s[0]/M == s[1]/M {
					//fmt.Printf("H2 n %d, h %d, pos %v\n", n, a, s)
					h = s[0] / M
				}
			case 3:
				if s[0]/M == s[1]/M && s[0]/M == s[2]/M {
					//fmt.Printf("H3 n %d, h %d, pos %v\n", n, a, s)
					h = s[0] / M
				}
			}

			if h != -1 {
				y0 := a / M * M
				x0 := h * M
				for j := 0; j < M; j++ {
					hy := y0 + j
					if hy == a {
						continue
					}

					for i := 0; i < M; i++ {
						hx := x0 + i
						c.kill(n, hx, hy)
					}
				}
			}
		} else if t == 2 {
			v := -1
			switch len(s) {
			case 2:
				if s[0]/M == s[1]/M {
					//fmt.Printf("V2 n %d, v %d, pos %v\n", n, a, s)
					v = s[0] / M
				}
			case 3:
				if s[0]/M == s[1]/M && s[0]/M == s[2]/M {
					//fmt.Printf("V3 n %d, v %d, pos %v\n", n, a, s)
					v = s[0] / M
				}
			}

			if v != -1 {
				x0 := a / M * M
				y0 := v * M
				for i := 0; i < M; i++ {
					hx := x0 + i
					if hx == a {
						continue
					}

					for j := 0; j < M; j++ {
						hy := y0 + j
						c.kill(n, hx, hy)
					}
				}
			}
		}

		// 检查是否有重叠的数字
		if len(c[n-1][t][a]) > 0 {
			ss := make(map[string][]int)
			sn := make(map[string][]int)
			for i := 0; i < N; i++ {
				s := c[i][t][a]
				if len(s) <= 1 {
					continue
				}

				v := fmt.Sprintf("%v", s)
				_, ok := ss[v]
				if ok {
					sn[v] = append(sn[v], i+1)
				} else {
					ss[v] = s
					sn[v] = []int{i + 1}
				}
			}

			if len(ss) > 1 {
				for v, s := range ss {
					ns := sn[v]
					if len(s) == len(ns) && len(s) > 1 {
						//fmt.Printf("pair t %d, a %d, pos %v, numbers %v\n", t, a, s, ns)
						set := make(map[int]bool)
						for _, n := range ns {
							set[n] = true
						}

						for i := 0; i < N; i++ {
							if _, ok := set[i+1]; !ok {
								for _, p := range s {
									c.free(i+1, t, a, p)
								}
							}
						}
					}
				}
			}
		}
	}
}

func (c *Choice) checkGroupPair(t, a int) {
	grids := [N][]int{}
	for i := range grids {
		grids[i] = make([]int, 0, N)
	}

	for i := 0; i < N; i++ {
		s := c[i][t][a]
		for _, p := range s {
			grids[p] = append(grids[p], i+1)
		}
	}

	ss := make(map[string][]int)
	sp := make(map[string][]int)
	for i := 0; i < N; i++ {
		ns := grids[i]
		if len(ns) <= 1 {
			continue
		}

		v := fmt.Sprintf("%v", ns)
		_, ok := ss[v]
		if ok {
			sp[v] = append(sp[v], i)
		} else {
			ss[v] = ns
			sp[v] = []int{i}
		}
	}

	if len(ss) > 1 {
		for v, ns := range ss {
			ps := sp[v]
			if len(ns) == len(ps) && len(ns) > 1 {
				//fmt.Printf("pair t %d, <%d>, numbers %v, pos %v\n", t, a, ns, ps)
				set := make(map[int]bool)
				for _, p := range ps {
					set[p] = true
				}

				for i := 0; i < N; i++ {
					if _, ok := set[i]; !ok {
						for _, n := range ns {
							if t == 0 {
								x, y := Cell2XY(a, i)
								c.kill(n, x, y)
							} else if t == 1 {
								c.kill(n, i, a)
							} else {
								c.kill(n, a, i)
							}
						}
					}
				}
			}
		}
	}
}

func (c *Choice) freeGroup(n, t, a int) {
	c[n-1][t][a] = []int{}
}

func (c *Choice) has(n, t, a, b int) bool {
	for _, v := range c[n-1][t][a] {
		if v == b {
			return true
		}
	}

	return false
}

func (c *Choice) Clone() *Choice {
	v := Choice{}
	for i := 0; i < N; i++ {
		for t := 0; t < 3; t++ {
			for j := 0; j < N; j++ {
				a := c[i][t][j]
				s := make([]int, len(a), N)
				copy(s, a)
				v[i][t][j] = s
			}
		}
	}

	return &v
}

func (c *Choice) best() (x, y, number int, exist bool) {
	for n := 1; n <= N; n++ {
		for i := 0; i < N; i++ {
			if len(c[n-1][0][i]) == 1 {
				x, y := Cell2XY(i, c[n-1][0][i][0])
				return x, y, n, true
			}

			if len(c[n-1][1][i]) == 1 {
				x, y := c[n-1][1][i][0], i
				return x, y, n, true
			}

			if len(c[n-1][2][i]) == 1 {
				x, y := i, c[n-1][2][i][0]
				return x, y, n, true
			}
		}
	}

	return 0, 0, 0, false
}

func (c *Choice) Print() {
	for n := 1; n <= N; n++ {
		fmt.Println(n, ":")
		for i := 0; i < N; i++ {
			if i == 0 {
				fmt.Printf("  > ")
			} else {
				fmt.Printf("  - ")
			}

			fmt.Println(c[n-1][0][i])
		}

		for i := 0; i < N; i++ {
			if i == 0 {
				fmt.Printf("  > ")
			} else {
				fmt.Printf("  - ")
			}

			fmt.Println(c[n-1][1][i])
		}

		for i := 0; i < N; i++ {
			if i == 0 {
				fmt.Printf("  > ")
			} else {
				fmt.Printf("  - ")
			}

			fmt.Println(c[n-1][2][i])
		}
	}
}

func (c *Choice) All() [][3]int {
	sn := [N + 1][]int{}
	sc := [N + 1][]int{}
	sp := [N + 1][][]int{}

	numbers := make([][3]int, 0)
	for i, a := range c {
		for c, pos := range a[0] {
			if len(pos) <= 1 {
				continue
			}

			n := i + 1
			d := len(pos)
			if len(sn[d]) == 0 {
				sn[d] = []int{n}
				sc[d] = []int{c}
				sp[d] = [][]int{pos}
			} else {
				sn[d] = append(sn[d], n)
				sc[d] = append(sc[d], c)
				sp[d] = append(sp[d], pos)
			}
		}
	}

	for i := 2; i <= N; i++ {
		for j := 0; j < len(sn[i]); j++ {
			n := sn[i][j]
			c := sc[i][j]
			for _, p := range sp[i][j] {
				numbers = append(numbers, [3]int{n, c, p})
			}
		}
	}

	return numbers
}

func (c *Choice) AllByNumbers() [][3]int {
	numbers := make([][3]int, 0)
	for i, a := range c {
		for c, pos := range a[0] {
			for _, p := range pos {
				numbers = append(numbers, [3]int{i + 1, c, p})
			}
		}
	}

	return numbers
}

type State struct {
	base   Board
	board  Board
	choice *Choice
}

func NewState(b Board) *State {
	return &State{b, b, b.Choice()}
}

func (s *State) Base() Board {
	return s.base
}

func (s *State) Board() Board {
	return s.board
}

func (s *State) Choice() *Choice {
	return s.choice
}

func (s *State) Next(n, x, y int) (*State, bool) {
	if !s.board.Empty(x, y) {
		return nil, false
	}

	ns := &State{s.board, s.board, s.choice.Clone()}
	if ns.choice.trySet(x, y, n) {
		ns.board.Set(x, y, n)
		if ns.Trim2() {
			return nil, false
		}

		if ns.board.Validate() == nil {
			return ns, true
		}
	}

	return nil, false
}

func (s *State) Debug() {
	s.choice.Print()
	if len(s.choice) > 0 {
		return
	}

	c, h, v, n := s.board.SplitLeft()
	fmt.Println("c: ", c)
	fmt.Println("h: ", h)
	fmt.Println("v: ", v)
	fmt.Println("n: ", n)
}
