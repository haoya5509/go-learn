package main

import (
	"fmt"
	"time"
)

func popCount(x int) int {
	cnt := 0
	for x != 0 {
		x = x & (x - 1)
		cnt++
	}
	return cnt
}

func main() {
	start := time.Now()
	fmt.Println(popCount(2))
	fmt.Println(popCount(3))
	fmt.Println(popCount(4))
	fmt.Println(popCount(5))
	fmt.Println(time.Since(start).Seconds())
}