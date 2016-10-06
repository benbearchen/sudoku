package nine2nine

import (
	_ "fmt"
)

func inter_set(a, b []int) []int {
	v := make([]int, 0, N)
	for len(a) > 0 && len(b) > 0 {
		d := a[0] - b[0]
		if d < 0 {
			a = a[1:]
		} else if d > 0 {
			b = b[1:]
		} else {
			v = append(v, a[0])
			a = a[1:]
			b = b[1:]
		}
	}

	return v
}

func (s *State) Trim() bool {
	meetError := false

	//fmt.Println("trim")
	loop := func() bool {
		c, h, v, n := s.board.SplitLeft()
		for i, a := range c {
			if len(a.num) != 1 {
				continue
			}

			for j := 0; j < N; j++ {
				x, y := Cell2XY(i, j)
				if s.board.Empty(x, y) {
					//s.board.Set(x, y, a.num[0])
					//return true
				}
			}
		}

		for i, a := range h {
			if len(a.num) != 1 {
				continue
			}

			for j := 0; j < N; j++ {
				x, y := j, i
				if s.board.Empty(x, y) {
					//s.board.Set(x, y, a.num[0])
					//return true
				}
			}
		}

		for i, a := range v {
			if len(a.num) != 1 {
				continue
			}

			for j := 0; j < N; j++ {
				x, y := i, j
				if s.board.Empty(x, y) {
					//s.board.Set(x, y, a.num[0])
					//return true
				}
			}
		}

		for i, a := range n {
			for _, mc := range a.c {
				can := make([]int, 0, N)
				for _, mp := range c[mc].pos {
					mx, my := Cell2XY(mc, mp)
					fx := false
					fy := false
					for _, x := range a.h {
						if x == mx {
							fx = true
							break
						}
					}

					for _, y := range a.v {
						if y == my {
							fy = true
							break
						}
					}

					if fx && fy {
						can = append(can, mp)
					}
				}

				if len(can) == 1 {
					x, y := Cell2XY(mc, can[0])
					s.board.Set(x, y, i+1)
					return true
				} else if len(can) == 0 {
					meetError = true
					//fmt.Println("meet error:", i+1, a, mc)
					return false
				}
			}
		}

		return false
	}

	for loop() {
	}

	//fmt.Println("trim out:", meetError)
	return meetError
}

func (s *State) Trim2() bool {
	meetError := false

	loop := func() bool {
		x, y, n, exist := s.choice.best()
		if !exist {
			return false
		}

		if !s.board.Empty(x, y) {
			meetError = true
			return false
		}

		if !s.choice.trySet(x, y, n) {
			meetError = true
			return false
		}

		s.board.Set(x, y, n)
		if len(s.board.Validate()) > 0 {
			meetError = true
			return false
		}

		return true
	}

	for loop() {
	}

	return meetError
}
