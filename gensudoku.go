package main

import (
	"github.com/benbearchen/sudoku/generator"

	"fmt"
	"time"
)

func main() {
	start := time.Now()
	GEN := 100
	succ := 0
	times := make(map[int]int)
	for i := 0; i < GEN; i++ {
		_, _, t, ok := generator.GenerateComplete()
		times[t]++

		if ok {
			succ++
		}
	}

	d := time.Now().Sub(start)
	fmt.Println("gen", GEN, "times used time:", d, "avg:", d/time.Duration(GEN), "succ:", succ)

	sum := 0
	for i := 1; len(times) > 0; i++ {
		t, ok := times[i]
		if ok {
			sum += t
			delete(times, i)
			fmt.Printf("try %3d in %4d times, %5.2f%% / %5.2f%%\n", i, t, float64(t)*100.0/float64(GEN), float64(sum)*100.0/float64(GEN))
		}
	}

	for i := 0; i < 10; i++ {
		start = time.Now()
		b, key := generator.SureGenerateComplete()
		fmt.Println("used time:", time.Now().Sub(start))

		b.Print()
		fmt.Println("-----------------")
		fmt.Println("key numbers:", key.Numbers())
		fmt.Println(key.OneLine())
		key.Print()
		fmt.Println("-----------------")

		fmt.Println("digg...")
		start = time.Now()
		low := generator.Digg(b)
		fmt.Println("used time:", time.Now().Sub(start))
		fmt.Println("numbers:", low.Numbers())
		fmt.Println(low.OneLine())
		fmt.Println("-----------------")
		low.Print()
	}
}
