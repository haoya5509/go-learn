package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println(expand("footestfoofoo", change))
}

func add1(r rune) rune {
	return r + 1
}

func change(s string) string{
	return strings.Map(func (r rune) rune {
		return r + 1
	}, s)
}

func expand(s string, f func(string) string) string {
	return strings.Join(strings.Split(s, "foo"), f("foo"))
}
