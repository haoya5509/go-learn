package main

import "fmt"

type Test int
func (a Test) String() string {
	return fmt.Sprintf("%g", a)
}

func main() {
	fmt.Println(gcd(4, 2))
	fmt.Println(gcd(4, 3))
}

func gcd(x, y int) int {
	if x < y {
		x, y = y, x
	}
	for y != 0 {
		x, y = y, x % y
	}
	return x
}