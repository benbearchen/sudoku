package main

import (
	n2n "github.com/benbearchen/sudoku/nine2nine"

	"fmt"
)

func main() {
	zero := n2n.Board([81]int{
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
	})

	zero.Validate()

	easy := n2n.Board([81]int{
		0, 0, 0, 4, 0, 0, 8, 0, 0,
		0, 0, 0, 0, 0, 7, 6, 2, 0,
		0, 1, 0, 0, 0, 0, 0, 0, 3,
		0, 0, 0, 5, 0, 4, 0, 0, 7,
		9, 3, 4, 0, 1, 0, 0, 8, 0,
		7, 0, 5, 9, 2, 3, 0, 0, 6,
		1, 0, 0, 8, 6, 9, 4, 5, 2,
		5, 2, 9, 3, 0, 1, 7, 6, 8,
		8, 4, 0, 2, 7, 0, 9, 3, 1,
	})

	low := n2n.Board([81]int{
		0, 1, 3, 0, 5, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 1, 0, 2, 0,
		4, 9, 6, 0, 0, 3, 0, 1, 0,
		0, 3, 4, 0, 0, 0, 1, 0, 0,
		0, 0, 0, 3, 0, 0, 0, 0, 6,
		9, 8, 1, 4, 0, 0, 0, 0, 3,
		7, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 6, 0, 9, 0, 2,
		0, 0, 0, 0, 9, 7, 5, 0, 0,
	})

	middle := n2n.Board([81]int{
		0, 5, 0, 0, 4, 0, 7, 0, 3,
		8, 2, 0, 0, 6, 7, 0, 4, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 2, 0, 0, 5, 0, 1, 6,
		6, 0, 9, 1, 2, 0, 8, 5, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 7, 0, 5, 8, 3, 0, 0,
		3, 0, 1, 0, 9, 4, 0, 0, 8,
	})

	high := n2n.Board([81]int{
		0, 0, 3, 0, 0, 0, 7, 0, 0,
		2, 0, 5, 0, 0, 0, 0, 8, 0,
		0, 8, 0, 2, 0, 3, 4, 0, 0,
		0, 0, 0, 8, 6, 0, 0, 1, 0,
		0, 0, 9, 4, 2, 0, 0, 0, 0,
		8, 0, 0, 0, 0, 5, 0, 0, 6,
		0, 5, 0, 0, 9, 8, 1, 0, 0,
		0, 1, 0, 0, 0, 0, 0, 4, 0,
		0, 0, 6, 0, 0, 0, 0, 3, 8,
	})

	high2 := n2n.Board([81]int{
		0, 0, 4, 9, 0, 0, 3, 7, 5,
		0, 9, 7, 8, 5, 0, 4, 6, 2,
		0, 5, 0, 0, 4, 7, 8, 1, 9,
		0, 4, 1, 5, 0, 0, 7, 9, 3,
		0, 3, 0, 1, 0, 0, 6, 0, 8,
		9, 0, 0, 0, 0, 0, 0, 2, 0,
		0, 0, 9, 0, 0, 0, 0, 0, 0,
		4, 1, 0, 0, 6, 5, 9, 0, 7,
		5, 0, 3, 7, 0, 0, 0, 0, 6,
	})

	hard := n2n.Board([81]int{
		0, 5, 0, 0, 8, 0, 0, 1, 7,
		0, 0, 2, 0, 5, 0, 0, 0, 0,
		9, 0, 0, 0, 0, 0, 0, 4, 0,
		0, 6, 0, 8, 0, 7, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		5, 0, 8, 3, 0, 0, 7, 0, 1,
		2, 0, 4, 9, 3, 6, 0, 0, 0,
		6, 0, 0, 0, 0, 0, 0, 0, 0,
		1, 0, 0, 0, 0, 0, 9, 3, 0,
	})

	hard2 := n2n.Board([81]int{
		0, 7, 0, 0, 0, 4, 0, 0, 0,
		0, 0, 0, 6, 0, 0, 7, 3, 0,
		4, 0, 8, 0, 5, 0, 0, 0, 1,
		0, 0, 1, 0, 0, 5, 0, 9, 0,
		8, 0, 5, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 8, 2,
		3, 0, 0, 9, 0, 0, 0, 1, 8,
		0, 0, 0, 0, 0, 7, 2, 0, 0,
		0, 6, 4, 0, 1, 0, 0, 0, 0,
	})

	all := []n2n.Board{easy, low, middle, high, high2, hard2, hard}

	for i, sudoku := range all {
		if i == 6 {
			break
		}

		fmt.Println("  ---")
		fmt.Println("sudoku:")
		sudoku.Print()
		fmt.Println("")

		errors := sudoku.Validate()
		if len(errors) > 0 {
			fmt.Println("sudoku has errors:")
			for _, e := range errors {
				fmt.Println(e)
			}

			continue
		}

		state := n2n.NewState(&sudoku)
		if state == nil {
			fmt.Println("can't create state")
			continue
		}

		state.Trim2()

		fmt.Println("after trim:")
		state.Board().Print()
		if !state.Board().Finished() {
			//state.Debug()

			//continue
			fmt.Println("start searching")
			search := n2n.NewSearch(state)
			results := search.Search()
			if len(results) == 0 {
				fmt.Println("Sh*t, can't found result")
			} else {
				fmt.Println(len(results), "result(s)")
				for _, r := range results {
					r.Print()
					fmt.Println("-----------------")
				}
			}
		}
	}
}
