package main

import "fmt"

func main()  {
	a := [...]string{"a"}
	//b := [...]string{"b"}
	//equal(a[:], b[:])
	a[0] = "b"

	arr := [3]int{1, 2, 3}
	reverse(&arr)
	for _, v := range arr {
		fmt.Println(v)
	}

	strings := []string{}
	strings = delDup(strings)
	for _, v := range strings {
		fmt.Println(v)
	}

	test()
}

func equal(x, y []string) bool {
	return true
}

func reverse(arr *[3]int) {
	l := len(*arr)
	for i := 0; i < l / 2; i++ {
		arr[i], arr[l-1-i] = arr[l-1-i], arr[i]
	}
}

func delDup(strings []string) []string {
	if len(strings) == 0 {
		return nil
	}

	p, q := 0, 0
	for q < len(strings) {
		if (strings[p] != strings[q]) {
			p++
			strings[p] = strings[q]
		}
		q++
	}
	return strings[:p+1]
}

func test() {
	arr := [3]int{1, 2, 3}
	sa := arr[:]
	sa[0] = 4
	for _, v := range arr {
		fmt.Println(v)
	}
}