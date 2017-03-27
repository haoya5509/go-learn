package main

import "fmt"

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())
	fmt.Println(x.Len())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())
	fmt.Println(y.Len())

	x.DifferenceWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.Len())

	x.Remove(9)
	fmt.Println(x.String())
	fmt.Println(x.Len())

	x.Clear()
	fmt.Println(&x)
}
