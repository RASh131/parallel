package main

import (
	"fmt"
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
	for i := uint(12); i <= 16; i++ {
		fmt.Println("~~~~~", i, "x", i, "~~~~~")
		s := time.Now()
		sum := 0
		xs := make([]int, 0, i)
		solver(func(_ []int) { sum++ }, (1<<i)-1, 0, 0, xs)
		sec := time.Now().Sub(s)
		fmt.Println("合計解数", sum)
		fmt.Println("所要時間", sec)
	}
}
