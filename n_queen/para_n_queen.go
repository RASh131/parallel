package main

import (
	"fmt"
	"runtime"
	"time"
)

func solver(f func([]int), n, right, left int, xs []int) {
	if n == 0 {
		f(xs)
	} else {
		for i := n; i > 0; i &= i - 1 {
			q := i & (-i)
			if q&(right|left) != 0 {
				continue
			}
			solver(f, n^q, (right|q)>>1, (left|q)<<1, append(xs, q))
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 12; i <= 16; i++ {
		fmt.Println("~~~~~", i, "x", i, "~~~~~")
		ch := make(chan int, i)
		s := time.Now()
		for j := (1 << uint(i)) - 1; j > 0; j &= j - 1 {
			go func(i uint, j int) {
				c := 0
				j &= -j
				xs := make([]int, 1, i)
				xs[0] = j
				solver(func(_ []int) { c++ }, ((1<<i)-1)^j, j>>1, j<<1, xs)
				ch <- c
			}(uint(i), j)
		}
		sum := 0
		for k := i; k > 0; k-- {
			sum += <-ch
		}
		sec := time.Now().Sub(s)
		fmt.Println(sum)
		fmt.Println(sec)
	}
}
